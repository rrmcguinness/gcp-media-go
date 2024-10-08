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

# Setup

## CI/CD with Bazel

Bazel is a unique build tool in that it supports most modern languages and propagates the
mono-repo style of development. In addition it's ideal for Go as it builds a hermetic
environment to run your CI/CD pipelines in.

```shell
# Add Node JS to your system

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
bazel build //cmd:pipeline 

# Build all targets in a specific package
bazel build //pkg/...

# Testing
bazel test //...

# Running Commands

bazel run //pkg:pipeline 
```

## Dependencies

Since Bazel is a hermetic build system, dependencies are done in two steps:

1. Install your dependency as normal: `go get github.com/stretchr/testify`
2. Update the mod spec for bazel: `bazel run //:gazelle-update-repos`

