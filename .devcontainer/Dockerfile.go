###############################################################################
# .devcontainer/Dockerfile.go                                                 #
#                                                                             #
# Description: Docker image used by VSCode for building a Golang development  #
#              sandbox.                                                       #
# See:         https://code.visualstudio.com/docs/remote/containers           #
###############################################################################


# -----------------------------------------------------------------------------
# BUILD-TIME VARIABLES
#
# Unfortunately, build-time arguments cannot be given in the devcontainer.json 
# file. This feature will certainly be available soon... so let's be prepared.
# -----------------------------------------------------------------------------

ARG ARG_BAZEL_VERSION=0.29.1


# -----------------------------------------------------------------------------
# SANDBOX STAGE
# -----------------------------------------------------------------------------

FROM golang:1.12.7-stretch

ARG ARG_BAZEL_VERSION

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure APT, install packages and tools
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    apt-utils \
    bash-completion \
    curl \
    g++ \
    git \
    lsb-release \
    patch \
    procps \
    unzip \
    zlib1g-dev \
    #
    # Install Bazel
 && curl -LO "https://github.com/bazelbuild/bazel/releases/download/$ARG_BAZEL_VERSION/bazel_$ARG_BAZEL_VERSION-linux-x86_64.deb" \
 && dpkg -i bazel_*.deb \
 && rm -rf bazel_*.deb \
    #
    # Install Gocode autocomplemention deamon
 && go get -x -d github.com/stamblerre/gocode 2>&1 \
 && go build -o gocode-gomod github.com/stamblerre/gocode \
 && mv gocode-gomod $GOPATH/bin/ \
    #
    # Install various Go tools
 && go get -u -v \
    github.com/bazelbuild/buildtools/buildifier \
    github.com/bazelbuild/bazel-gazelle/cmd/gazelle \
    github.com/mdempsky/gocode \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    github.com/godoctor/godoctor \
    golang.org/x/tools/cmd/guru \
    golang.org/x/tools/cmd/gorename \
    github.com/rogpeppe/godef \
    github.com/zmb3/gogetdoc \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/sqs/goreturns \
    github.com/josharian/impl \
    github.com/davidrjenni/reftools/cmd/fillstruct \
    github.com/fatih/gomodifytags \
    github.com/cweill/gotests/... \
    golang.org/x/tools/cmd/goimports \
    golang.org/x/lint/golint \
    golang.org/x/tools/cmd/gopls \
    github.com/alecthomas/gometalinter \
    honnef.co/go/tools/... \
    github.com/golangci/golangci-lint/cmd/golangci-lint \
    github.com/mgechev/revive \
    github.com/99designs/gqlgen \
    github.com/bazelbuild/bazelisk \
    github.com/derekparker/delve/cmd/dlv 2>&1 \
    #
    # Clean up
 && apt-get autoremove -y \
 && apt-get clean -y \
 && rm -rf /var/lib/apt/lists/*