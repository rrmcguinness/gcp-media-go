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

## Developer Tools

Use the following instructions to set up a development environment:
* [Workstation Setup](WorkstationSetup.md]
* [Setting Up IntelliJ](SettingUpIntelliJ.md]
* [Setting Up Visual Studio Code](SettingUpVisualStudioCode.md]

## Running the Demo

```shell
# The following command combines two commands to simplify how the demo can be run
# bazel run //web/apps/api_server and bazel run //web/apps/media-search:start  
bazel run //:demo
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

# Update Build Files and Dependencies
# Used when getting "missing strict dependency errors"
bazel run //:gazelle
```

