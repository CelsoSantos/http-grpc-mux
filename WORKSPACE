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

# Fetch gRPC and Protobuf dependencies (should be fetched before Go rules)
http_archive(
    name = "build_stack_rules_proto",
    sha256 = "5474d1b83e24ec1a6db371033a27ff7aff412f2b23abba86fedd902330b61ee6",
    strip_prefix = "rules_proto-91cbae9bd71a9c51406014b8b3c931652fb6e660",
    urls = ["https://github.com/stackb/rules_proto/archive/91cbae9bd71a9c51406014b8b3c931652fb6e660.tar.gz"],
)

load("@build_stack_rules_proto//go:deps.bzl", "go_grpc_library")

go_grpc_library()

# Import Golang Bazel repository into the workspace
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8df59f11fb697743cbb3f26cfb8750395f30471e9eabde0d174c3aebc7a1cd39",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
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
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
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
