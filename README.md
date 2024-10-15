<!--
 Copyright 2024 Google, LLC
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     https://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# Video Warehouse & Search

This is a project for processing video, extracting intelligence, persisting to a dataset,
and enabling AI interactions with the data. The modular side of the project is built on the
Chain of Responsibility (COR) design pattern. Each unit of work is atomic and state is only
conveyed via Context to each chain and/or command.

## Tooling

Bazel is a unique build tool in that it supports most modern languages and propagates the
mono-repo style of development. In addition it's ideal for Go as it builds a hermetic
environment to run your CI/CD pipelines in.

```shell
# Add Node JS to your development environment

# Install the version manager
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash

# Reload your profile
source ~/.zshrc 

# Install Node
nvm install 22

# Install Bazelisk (used for managing multiple Bazel environments)
npm install -g bazelisk
```

## Building
```shell

# Build all targets
bazel build //...

# Build a specific target (The pipeline target in the pkg directory)
bazel build //pkg/model

# Build all targets in a specific package
bazel build //pkg/...

# Testing
bazel test //...

# Running Commands

bazel run //cmd:pipeline

# Cleaning
bazel clean

# Clean all cache
bazel clean --expunge
```

## Dependencies

The process for updating dependencies requires two steps:

1) Add a dependency using go `go get ...`
2) Make bazel aware of the change: `bazel run //:gazelle-update-repos`

> NOTE: Since bazel builds all go dependencies it's important to keep these in sync.

## Using Gazelle

Gazelle is a go program created to make developing Go for the enterprise a little more seamless.

```shell
# Update all build scripts: USE with caution as it may break existing builds
bazel run //:gazelle
```

