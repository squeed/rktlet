package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    tags = ["automanaged"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//pkg/genericapiserver/authenticator:all-srcs",
        "//pkg/genericapiserver/authorizer:all-srcs",
        "//pkg/genericapiserver/endpoints:all-srcs",
        "//pkg/genericapiserver/registry:all-srcs",
        "//pkg/genericapiserver/server:all-srcs",
    ],
    tags = ["automanaged"],
)
