# Setting Up IntelliJ Ultimate as your IDE

<!---
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
--->

IntelliJ and Bazel work very well together, especially when using
multiple languages. That being said, there are some steps to follow
to make it operate well.

## Go Lang tool setup

Since this if for IntelliJ Ultimate, you should install a couple of the 
go tools first and ensure they are on your path. View the [Workstation Setup](WorkstationSetup.md)
file at the root for more information.

## Plugins

Install the following plugins:

* Go
* Bazel
* Terraform
* Google Cloud Code

> If you set up your environment correct, you're ready to go (pun intended)
> Otherwise, you'll need to configure your buildifier path on the Bazel tab.

## Why IntelliJ Idea

IntelliJ offers a higher quality of tooling over VS Code when it comes to executing and debugging
Go and Bazel and Node JS.

## Importing a bazel project

Once you've restarted you're IDE, you'll see the ability to import a bazel project.
Click on that button and navigate to the project directory.

You'll be presented with a configuration screen, below is a working configuration
you should copy and paste.

```yaml
directories:
  # Add the directories you want added as source here
  # By default, we've added your entire workspace ('.')
  -bazel-bin
  -bazel-out
  -bazel-testlogs
  -bazel-vide-warehouse-go
  -.bazelrc
  -video-warehouse.code-workspace
  -.trunk
  .

# Automatically includes all relevant targets under the 'directories' above
derive_targets_from_directories: true

targets:
  # If source code isn't resolving, add additional targets that compile it here

additional_languages:
  go
  javascript
  typescript
```

## Finalize
Now you're ready to go. If you want a clean file tree, click on the vertical "..."
on the project explorer, and choose "appearance > (uncheck) Exclude files.". This
will exclude the normally hidden files.

You may build and synchronize using the "Bazel" menu on the system tray.
