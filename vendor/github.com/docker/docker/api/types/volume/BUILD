load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "volumes_create.go",
        "volumes_list.go",
    ],
    importmap = "grail/go/src/vendor/github.com/docker/docker/api/types/volume",
    importpath = "github.com/docker/docker/api/types/volume",
    visibility = ["//visibility:public"],
    deps = ["//go/src/vendor/github.com/docker/docker/api/types:go_default_library"],
)
