load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "debug.go",
        "decode.go",
        "encode.go",
        "error.go",
        "header.go",
        "header_value.go",
        "message.go",
    ],
    importmap = "grail/go/src/vendor/github.com/aws/aws-sdk-go/private/protocol/eventstream",
    importpath = "github.com/aws/aws-sdk-go/private/protocol/eventstream",
    visibility = ["//visibility:public"],
    deps = ["//go/src/vendor/github.com/aws/aws-sdk-go/aws:go_default_library"],
)
