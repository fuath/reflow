load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["jsonrpc.go"],
    importmap = "grail/go/src/vendor/github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    importpath = "github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    visibility = ["//visibility:public"],
    deps = [
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws/awserr:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws/request:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/private/protocol/json/jsonutil:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/private/protocol/rest:go_default_library",
    ],
)
