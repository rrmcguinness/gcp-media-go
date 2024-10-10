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

package model

type Command interface {
	IsExecutable(chCtx ChainContext) bool
	Execute(chCtx ChainContext)
}

type Chain interface {
	Command
	ContinueOnFailure(bool) Chain
	AddCommand(command Command) Chain
}

type BaseChain struct {
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

func (c *BaseChain) IsExecutable(context ChainContext) bool {
	return true
}

func (c *BaseChain) Execute(chCtx ChainContext) {
	for _, command := range c.commands {
		if chCtx.HasErrors() && !c.continueOnFailure {
			break
		} else if command.IsExecutable(chCtx) {
			command.Execute(chCtx)
		}
	}
}
