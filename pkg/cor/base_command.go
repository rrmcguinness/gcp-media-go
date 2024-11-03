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

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// BaseCommand is the default implementation of Command
type BaseCommand struct {
	Name            string
	InputParamName  string
	OutputParamName string
	Tracer          trace.Tracer
}

func NewBaseCommand(name string) *BaseCommand {
	return &BaseCommand{Name: name, Tracer: otel.Tracer(name)}
}

func (c *BaseCommand) GetName() string {
	return c.Name
}

// IsExecutable a default implementation of IsExecutable.
func (c *BaseCommand) IsExecutable(context Context) bool {
	return context != nil && context.Get(c.GetInputParam()) != nil && context.GetContext() != nil
}

// GetInputParam the name of the parameter expected as the primary input,
// if empty it will default to CtxIn, during a chain execution event CtxIn
// will be mapped to the previous executions CtxOut to ensure PIPE / chain behaviors.
func (c *BaseCommand) GetInputParam() string {
	if len(c.InputParamName) == 0 {
		return CtxIn
	}
	return c.InputParamName
}

// GetOutputParam the name of the output parameter, the default is CtxOut
// See the chain execute method for more detail.
func (c *BaseCommand) GetOutputParam() string {
	if len(c.OutputParamName) == 0 {
		return CtxOut
	}
	return c.OutputParamName
}
