load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["assume_role_provider.go"],
    importmap = "grail/go/src/vendor/github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    importpath = "github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    visibility = ["//visibility:public"],
    deps = [
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws/awserr:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws/client:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/aws/credentials:go_default_library",
        "//go/src/vendor/github.com/aws/aws-sdk-go/service/sts:go_default_library",
    ],
)
