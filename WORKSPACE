workspace(name = "gloodemo")

# -----------------------------------------------------------------------------
# Global variables
# -----------------------------------------------------------------------------

# Requisite minimal Golang toolchain version
MINIMAL_GOLANG_VERSION = "1.12.4"

# Requisite minimal Bazel version requested to build this project
MINIMAL_BAZEL_VERSION = "0.23.0"

# Requisite minimal Gazelle version compatible with Golang Bazel rules
MINIMAL_GAZELLE_VERSION = "0.18.1"

# Requisite minimal Golang Bazel rules (must be set in accordance with minimal Gazelle version)
#
# @see https://github.com/bazelbuild/bazel-gazelle#compatibility)
MINIMAL_GOLANG_BAZEL_RULES_VERSION = "0.19.0"

MINIMAL_GOLANG_BAZEL_RULES_SHASUM = "9fb16af4d4836c8222142e54c9efa0bb5fc562ffc893ce2abeac3e25daead144"

# -----------------------------------------------------------------------------
# Basic Bazel settings
# -----------------------------------------------------------------------------

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Import Bazel Skylib repository into the workspace
http_archive(
    name = "bazel_skylib",
    sha256 = "eb5c57e4c12e68c0c20bc774bfbc60a568e800d025557bc4ea022c6479acc867",
    strip_prefix = "bazel-skylib-0.6.0",
    url = "https://github.com/bazelbuild/bazel-skylib/archive/0.6.0.tar.gz",
)

load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(
    minimum_bazel_version = MINIMAL_BAZEL_VERSION,
)

# -----------------------------------------------------------------------------
# Golang and gRPC tools and external dependencies
# -----------------------------------------------------------------------------

# Fetch Protobuf dependencies
http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = ["https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz"],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# Fetch gRPC and Protobuf dependencies (should be fetched before Go rules)
http_archive(
    name = "build_stack_rules_proto",
    sha256 = "c62f0b442e82a6152fcd5b1c0b7c4028233a9e314078952b6b04253421d56d61",
    strip_prefix = "rules_proto-b93b544f851fdcd3fc5c3d47aee3b7ca158a8841",
    urls = ["https://github.com/stackb/rules_proto/archive/b93b544f851fdcd3fc5c3d47aee3b7ca158a8841.tar.gz"],
)

load("@build_stack_rules_proto//go:deps.bzl", "go_grpc_library")

go_grpc_library()

# Import Golang Bazel repository into the workspace
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "513c12397db1bc9aa46dd62f02dd94b49a9b5d17444d49b5a04c5a89f3053c1c",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
    ],
)

# Fetch Golang dependencies.
#
# The 'go_rules_dependencies()' is a macro that registers external dependencies needed by
# the Go and proto rules in rules_go.
# You can override a dependency declared in go_rules_dependencies by declaring a repository
# rule in WORKSPACE with the same name BEFORE the call to 'go_rules_dependencies()' macro.
#
# You can find the full implementation in repositories.bzl availble at https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl.
#
# @see: https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id5
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# NOTE:
# To override dependencies declared in 'go_rules_dependencies' macro, you should
# declare your dependencies here, before invoking 'go_rules_dependencies' macro.

# Fetch external dependencies needed by the Go and proto rules in rules_go, as
# well as some basic Golang packages, such as, for instance, 'golang.org/x/text'
# i18n tool.
go_rules_dependencies()

go_register_toolchains()
#  go_version = MINIMAL_GOLANG_VERSION,
#)

# -----------------------------------------------------------------------------
# Bazel Gazelle build files generator settings
# -----------------------------------------------------------------------------

# Import Gazelle repository into the workspace
http_archive(
    name = "bazel_gazelle",
    sha256 = "41bff2a0b32b02f20c227d234aa25ef3783998e5453f7eade929704dcff7cd4b",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.0/bazel-gazelle-v0.19.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.0/bazel-gazelle-v0.19.0.tar.gz",
    ],
)

# Fetch Gazelle dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# -----------------------------------------------------------------------------
# Docker Bazel rules dependencies
# -----------------------------------------------------------------------------

# Import the rules_docker repository at release v0.7.0 - be sure the Go rules
# are previously loaded - see above Go rules section in this file
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

# Load the macro that allows you to customize the docker toolchain configuration.
load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

docker_toolchain_configure(
    name = "docker_config",
    # Replace this with a path to a directory which has a custom docker client
    # config.json. Docker allows you to specify custom authentication credentials
    # in the client configuration JSON file.
    # See https://docs.docker.com/engine/reference/commandline/cli/#configuration-files
    # for more details.
    #
    # IMPORTANT: This path is relative to the sandbox workspace directory
    client_config = "/workspace/tooling/docker",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# -----------------------------------------------------------------------------
# External dependencies
# -----------------------------------------------------------------------------

load("//tooling/bazel:dependencies.bzl", "ext_dependencies")

ext_dependencies()
