// Copyright 2024 Google, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cor

import "fmt"

type BaseChain struct {
	BaseCommand
	continueOnFailure bool
	commands          []Command
}

func (c *BaseChain) ContinueOnFailure(continueOnFailure bool) Chain {
	c.continueOnFailure = continueOnFailure
	return c
}

func (c *BaseChain) AddCommand(command Command) Chain {
	c.commands = append(c.commands, command)
	return c
}

func (c *BaseChain) IsExecutable(context Context) bool {
	return true
}

func (c *BaseChain) Execute(chCtx Context) {
	defer chCtx.Close()
	for i, command := range c.commands {
		// Ensure that the next parameter is callable in a pipe stack
		fmt.Printf("%d\n%v\n", i, chCtx)

		if chCtx.HasErrors() && !c.continueOnFailure {
			break
		} else if command.IsExecutable(chCtx) {
			command.Execute(chCtx)
		}
		fmt.Printf("%d\n%v\n", i, chCtx)
		// Chain the output
		chCtx.Remove(CTX_IN)
		chCtx.Add(CTX_IN, chCtx.Get(CTX_OUT))
		chCtx.Remove(CTX_OUT)
	}
}
