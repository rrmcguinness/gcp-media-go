# Workstation Setup

This will help you set up your developer environment outside the IDE.

## Node JS

Node JS is used as a general purpose program manager and JavaScript / Typescript transpiler.

```shell
# Add Node JS to your development environment

# Install the version manager
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash

# Reload your profile
source ~/.zshrc 

# Install Node
nvm install 22
````

## PNPM

This project uses PNPM a more efficient package manager. Be sure to install the stable 8.x version.
Version 9 introduced breaking changes not compatible yet with other tooling.

```shell
npm install -g pnpm@8.15.8
```

## Bazel

Bazel is a unique build tool in that it supports most modern languages and propagates the
mono-repo style of development. In addition, it's ideal for Go as it builds a hermetic
environment to run your CI/CD pipelines in.

```shell

# The CLI for building
npm install -g @bazel/bazelisk
# An efficient runtime wrapper for bazel targets
npm install -g @bazel/ibazel
```

## Go

```shell
# Install Buildifier, Buildozer, and Unused Deps
go install github.com/bazelbuild/buildtools/buildifier@latest
go install github.com/bazelbuild/buildtools/buildozer@latest
go install github.com/bazelbuild/buildtools/unused_deps@latest

# Install dlv, the golang debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Lastly, add go tooling to your system path
# vim ~/.zshrc or ~/.bashrc or ~/.bash_profile
# export PATH=$PATH:$HOME/go/bin
```

## Terraform


