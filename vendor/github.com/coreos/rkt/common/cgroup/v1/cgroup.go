// Copyright 2016 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build linux

package v1

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/hashicorp/errwrap"
)

var (
	cgroupControllerRWFiles = map[string][]string{
		"memory":  {"memory.limit_in_bytes"},
		"cpu":     {"cpu.cfs_quota_us", "cpu.shares"},
		"devices": {"devices.allow", "devices.deny"},
	}
)

// mountFsRO remounts the given mountPoint using the given flags read-only.
func mountFsRO(mountPoint string, flags uintptr) error {
	flags = flags |
		syscall.MS_BIND |
		syscall.MS_REMOUNT |
		syscall.MS_RDONLY

	if err := syscall.Mount(mountPoint, mountPoint, "", flags, ""); err != nil {
		return errwrap.Wrap(fmt.Errorf("error remounting read-only %q", mountPoint), err)
	}

	return nil
}

func parseCgroups(f io.Reader) (map[int][]string, error) {
	sc := bufio.NewScanner(f)

	// skip first line since it is a comment
	sc.Scan()

	cgroups := make(map[int][]string)
	for sc.Scan() {
		var controller string
		var hierarchy int
		var num int
		var enabled int
		fmt.Sscanf(sc.Text(), "%s %d %d %d", &controller, &hierarchy, &num, &enabled)

		if enabled == 1 {
			if _, ok := cgroups[hierarchy]; !ok {
				cgroups[hierarchy] = []string{controller}
			} else {
				cgroups[hierarchy] = append(cgroups[hierarchy], controller)
			}
		}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return cgroups, nil
}

// GetEnabledCgroups returns a map with the enabled cgroup controllers grouped by
// hierarchy
func GetEnabledCgroups() (map[int][]string, error) {
	cgroupsFile, err := os.Open("/proc/cgroups")
	if err != nil {
		return nil, err
	}
	defer cgroupsFile.Close()

	cgroups, err := parseCgroups(cgroupsFile)
	if err != nil {
		return nil, errwrap.Wrap(errors.New("error parsing /proc/cgroups"), err)
	}

	return cgroups, nil
}

// GetControllerDirs takes a map with the enabled cgroup controllers grouped by
// hierarchy and returns the directory names as they should be in
// /sys/fs/cgroup
func GetControllerDirs(cgroups map[int][]string) []string {
	var controllers []string
	for _, cs := range cgroups {
		controllers = append(controllers, strings.Join(cs, ","))
	}

	return controllers
}

func getControllerSymlinks(cgroups map[int][]string) map[string]string {
	symlinks := make(map[string]string)

	for _, cs := range cgroups {
		if len(cs) > 1 {
			tgt := strings.Join(cs, ",")
			for _, ln := range cs {
				symlinks[ln] = tgt
			}
		}
	}

	return symlinks
}

func CgroupControllerRWFiles(isolator string) []string {
	files, _ := cgroupControllerRWFiles[isolator]
	return files
}

func getControllerRWFiles(controller string) []string {
	parts := strings.Split(controller, ",")
	for _, p := range parts {
		if files, ok := cgroupControllerRWFiles[p]; ok {
			// cgroup.procs always needs to be RW for allowing systemd to add
			// processes to the controller
			files = append(files, "cgroup.procs")
			return files
		}
	}

	return nil
}

func parseCgroupController(cgroupPath, controller string) ([]string, error) {
	cg, err := os.Open(cgroupPath)
	if err != nil {
		return nil, errwrap.Wrap(errors.New("error opening /proc/self/cgroup"), err)
	}
	defer cg.Close()

	s := bufio.NewScanner(cg)
	for s.Scan() {
		parts := strings.SplitN(s.Text(), ":", 3)
		if len(parts) < 3 {
			return nil, fmt.Errorf("error parsing /proc/self/cgroup")
		}
		controllerParts := strings.Split(parts[1], ",")
		for _, c := range controllerParts {
			if c == controller {
				return parts, nil
			}
		}
	}

	return nil, fmt.Errorf("controller %q not found", controller)
}

// GetOwnCgroupPath returns the cgroup path of this process in controller
// hierarchy
func GetOwnCgroupPath(controller string) (string, error) {
	parts, err := parseCgroupController("/proc/self/cgroup", controller)
	if err != nil {
		return "", err
	}
	return parts[2], nil
}

// GetCgroupPathByPid returns the cgroup path of the process with the given pid
// and given controller.
func GetCgroupPathByPid(pid int, controller string) (string, error) {
	parts, err := parseCgroupController(fmt.Sprintf("/proc/%d/cgroup", pid), controller)
	if err != nil {
		return "", err
	}
	return parts[2], nil
}

// JoinSubcgroup makes the calling process join the subcgroup hierarchy on a
// particular controller
func JoinSubcgroup(controller string, subcgroup string) error {
	subcgroupPath := filepath.Join("/sys/fs/cgroup", controller, subcgroup)
	if err := os.MkdirAll(subcgroupPath, 0600); err != nil {
		return errwrap.Wrap(fmt.Errorf("error creating %q subcgroup", subcgroup), err)
	}
	pidBytes := []byte(strconv.Itoa(os.Getpid()))
	if err := ioutil.WriteFile(filepath.Join(subcgroupPath, "cgroup.procs"), pidBytes, 0600); err != nil {
		return errwrap.Wrap(fmt.Errorf("error adding ourselves to the %q subcgroup", subcgroup), err)
	}

	return nil
}

// If /system.slice does not exist in the cpuset controller, create it and
// configure it.
// Since this is a workaround, we ignore errors
func fixCpusetKnobs(cpusetPath string) {
	cgroupPathFix := filepath.Join(cpusetPath, "system.slice")
	_ = os.MkdirAll(cgroupPathFix, 0755)
	knobs := []string{"cpuset.mems", "cpuset.cpus"}
	for _, knob := range knobs {
		parentFile := filepath.Join(filepath.Dir(cgroupPathFix), knob)
		childFile := filepath.Join(cgroupPathFix, knob)

		data, err := ioutil.ReadFile(childFile)
		if err != nil {
			continue
		}
		// If the file is already configured, don't change it
		if strings.TrimSpace(string(data)) != "" {
			continue
		}

		data, err = ioutil.ReadFile(parentFile)
		if err == nil {
			// Workaround: just write twice to workaround the kernel bug fixed by this commit:
			// https://github.com/torvalds/linux/commit/24ee3cf89bef04e8bc23788aca4e029a3f0f06d9
			ioutil.WriteFile(childFile, data, 0644)
			ioutil.WriteFile(childFile, data, 0644)
		}
	}
}

// IsControllerMounted returns whether a controller is mounted by checking that
// cgroup.procs is accessible
func IsControllerMounted(c string) bool {
	cgroupProcsPath := filepath.Join("/sys/fs/cgroup", c, "cgroup.procs")
	if _, err := os.Stat(cgroupProcsPath); err != nil {
		return false
	}

	return true
}

// CreateCgroups mounts the v1 cgroup controllers hierarchy in /sys/fs/cgroup
// under root
func CreateCgroups(root string, enabledCgroups map[int][]string, mountContext string) error {
	controllers := GetControllerDirs(enabledCgroups)

	sys := filepath.Join(root, "/sys")
	if err := os.MkdirAll(sys, 0700); err != nil {
		return err
	}

	var sysfsFlags uintptr = syscall.MS_NOSUID |
		syscall.MS_NOEXEC |
		syscall.MS_NODEV

	// If we're mounting the host cgroups, /sys is probably mounted so we
	// ignore EBUSY
	if err := syscall.Mount("sysfs", sys, "sysfs", sysfsFlags, ""); err != nil && err != syscall.EBUSY {
		return errwrap.Wrap(fmt.Errorf("error mounting %q", sys), err)
	}

	cgroupTmpfs := filepath.Join(root, "/sys/fs/cgroup")
	if err := os.MkdirAll(cgroupTmpfs, 0700); err != nil {
		return err
	}

	var cgroupTmpfsFlags uintptr = syscall.MS_NOSUID |
		syscall.MS_NOEXEC |
		syscall.MS_NODEV |
		syscall.MS_STRICTATIME

	options := "mode=755"
	if mountContext != "" {
		options = fmt.Sprintf("mode=755,context=\"%s\"", mountContext)
	}

	if err := syscall.Mount("tmpfs", cgroupTmpfs, "tmpfs", cgroupTmpfsFlags, options); err != nil {
		return errwrap.Wrap(fmt.Errorf("error mounting %q", cgroupTmpfs), err)
	}

	// Mount controllers
	for _, c := range controllers {
		cPath := filepath.Join(root, "/sys/fs/cgroup", c)
		if err := os.MkdirAll(cPath, 0700); err != nil {
			return err
		}

		var flags uintptr = syscall.MS_NOSUID |
			syscall.MS_NOEXEC |
			syscall.MS_NODEV

		if err := syscall.Mount("cgroup", cPath, "cgroup", flags, c); err != nil {
			return errwrap.Wrap(fmt.Errorf("error mounting %q", cPath), err)
		}
	}

	// Create symlinks for combined controllers
	symlinks := getControllerSymlinks(enabledCgroups)
	for ln, tgt := range symlinks {
		lnPath := filepath.Join(cgroupTmpfs, ln)
		if err := os.Symlink(tgt, lnPath); err != nil {
			return errwrap.Wrap(errors.New("error creating symlink"), err)
		}
	}

	systemdControllerPath := filepath.Join(root, "/sys/fs/cgroup/systemd")
	if err := os.MkdirAll(systemdControllerPath, 0700); err != nil {
		return err
	}

	// Bind-mount cgroup tmpfs filesystem read-only
	return mountFsRO(cgroupTmpfs, cgroupTmpfsFlags)
}

// RemountCgroupsRO remounts the v1 cgroup hierarchy under root read-only,
// leaving the needed knobs in the subcgroup for each app read-write so the
// systemd inside stage1 can apply isolators to them
func RemountCgroupsRO(root string, enabledCgroups map[int][]string, subcgroup string, serviceNames []string) error {
	controllers := GetControllerDirs(enabledCgroups)
	cgroupTmpfs := filepath.Join(root, "/sys/fs/cgroup")
	sysPath := filepath.Join(root, "/sys")

	var flags uintptr = syscall.MS_NOSUID |
		syscall.MS_NOEXEC |
		syscall.MS_NODEV

	// Mount RW knobs we need to make the enabled isolators work
	for _, c := range controllers {
		cPath := filepath.Join(cgroupTmpfs, c)
		subcgroupPath := filepath.Join(cPath, subcgroup, "system.slice")

		// Workaround for https://github.com/coreos/rkt/issues/1210
		if c == "cpuset" {
			fixCpusetKnobs(cPath)
		}

		// Create cgroup directories and mount the files we need over
		// themselves so they stay read-write
		for _, serviceName := range serviceNames {
			appCgroup := filepath.Join(subcgroupPath, serviceName)
			if err := os.MkdirAll(appCgroup, 0755); err != nil {
				return err
			}
			for _, f := range getControllerRWFiles(c) {
				cgroupFilePath := filepath.Join(appCgroup, f)
				// the file may not be there if kernel doesn't support the
				// feature, skip it in that case
				if _, err := os.Stat(cgroupFilePath); os.IsNotExist(err) {
					continue
				}
				if err := syscall.Mount(cgroupFilePath, cgroupFilePath, "", syscall.MS_BIND, ""); err != nil {
					return errwrap.Wrap(fmt.Errorf("error bind mounting %q", cgroupFilePath), err)
				}
			}
		}

		// Re-mount controller read-only to prevent the container modifying host controllers
		if err := mountFsRO(cPath, flags); err != nil {
			return err
		}
	}

	// Bind-mount sys filesystem read-only
	return mountFsRO(sysPath, flags)
}

// RemountCgroupKnobsRW remounts the needed knobs in the subcgroup for one
// specified app read-write so the systemd inside stage1 can apply isolators
// to them.
func RemountCgroupKnobsRW(enabledCgroups map[int][]string, subcgroup string, serviceName string, enterCmd []string) error {
	controllers := GetControllerDirs(enabledCgroups)

	// Mount RW knobs we need to make the enabled isolators work
	for _, c := range controllers {
		cPath := filepath.Join("/sys/fs/cgroup", c)
		subcgroupPath := filepath.Join(cPath, subcgroup, "system.slice")

		// Create cgroup directories and mount the files we need over
		// themselves so they stay read-write
		appCgroup := filepath.Join(subcgroupPath, serviceName)
		if err := os.MkdirAll(appCgroup, 0755); err != nil {
			return err
		}
		for _, f := range getControllerRWFiles(c) {
			cgroupFilePath := filepath.Join(appCgroup, f)
			// the file may not be there if kernel doesn't support the
			// feature, skip it in that case
			if _, err := os.Stat(cgroupFilePath); os.IsNotExist(err) {
				continue
			}

			// Go applications cannot be  reassociated  with a new mount
			// namespace because they are multithreaded. Instead of
			// syscall.Mount, uses the enter entrypoint.
			argsMountBind := append(enterCmd, "/bin/mount", "--bind", cgroupFilePath, cgroupFilePath)
			cmdMountBind := exec.Cmd{
				Path: argsMountBind[0],
				Args: argsMountBind,
			}

			if err := cmdMountBind.Run(); err != nil {
				return err
			}

			argsRemountRW := append(enterCmd, "/bin/mount", "-o", "remount,bind,rw", cgroupFilePath)
			cmdRemountRW := exec.Cmd{
				Path: argsRemountRW[0],
				Args: argsRemountRW,
			}

			if err := cmdRemountRW.Run(); err != nil {
				return err
			}
		}
	}

	return nil
}
