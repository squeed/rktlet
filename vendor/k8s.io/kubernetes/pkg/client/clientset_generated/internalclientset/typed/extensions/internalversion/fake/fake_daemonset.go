/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	api "k8s.io/kubernetes/pkg/api"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeDaemonSets implements DaemonSetInterface
type FakeDaemonSets struct {
	Fake *FakeExtensions
	ns   string
}

var daemonsetsResource = schema.GroupVersionResource{Group: "extensions", Version: "", Resource: "daemonsets"}

func (c *FakeDaemonSets) Create(daemonSet *extensions.DaemonSet) (result *extensions.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &extensions.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.DaemonSet), err
}

func (c *FakeDaemonSets) Update(daemonSet *extensions.DaemonSet) (result *extensions.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &extensions.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.DaemonSet), err
}

func (c *FakeDaemonSets) UpdateStatus(daemonSet *extensions.DaemonSet) (*extensions.DaemonSet, error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &extensions.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.DaemonSet), err
}

func (c *FakeDaemonSets) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction(daemonsetsResource, c.ns, name), &extensions.DaemonSet{})

	return err
}

func (c *FakeDaemonSets) DeleteCollection(options *api.DeleteOptions, listOptions v1.ListOptions) error {
	action := core.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &extensions.DaemonSetList{})
	return err
}

func (c *FakeDaemonSets) Get(name string, options v1.GetOptions) (result *extensions.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction(daemonsetsResource, c.ns, name), &extensions.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.DaemonSet), err
}

func (c *FakeDaemonSets) List(opts v1.ListOptions) (result *extensions.DaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction(daemonsetsResource, c.ns, opts), &extensions.DaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensions.DaemonSetList{}
	for _, item := range obj.(*extensions.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daemonSets.
func (c *FakeDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction(daemonsetsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched daemonSet.
func (c *FakeDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(core.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, data, subresources...), &extensions.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.DaemonSet), err
}
