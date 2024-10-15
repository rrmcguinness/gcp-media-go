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

import "os"

type ChainContext interface {
	Add(key string, value interface{}) ChainContext
	AddError(err error)
	GetErrors() []error
	Get(key string) interface{}
	Remove(key string)
	HasErrors() bool
	AddTempFile(file string)
	GetTempFiles() []string
	Close()
}

type BaseChainContext struct {
	data      map[string]interface{}
	errors    []error
	tempFiles []string
}

func NewChainContext() ChainContext {
	return &BaseChainContext{
		data:      make(map[string]interface{}),
		errors:    make([]error, 0),
		tempFiles: make([]string, 0),
	}
}

func (c *BaseChainContext) Close() {
	// Clean up any temp files created along the way
	for _, file := range c.GetTempFiles() {
		os.Remove(file)
	}
}

func (c *BaseChainContext) Add(key string, value interface{}) ChainContext {
	c.data[key] = value
	return c
}

func (c *BaseChainContext) AddTempFile(file string) {
	c.tempFiles = append(c.tempFiles, file)
}

func (c *BaseChainContext) GetTempFiles() []string {
	return c.tempFiles
}

func (c *BaseChainContext) AddError(err error) {
	c.errors = append(c.errors, err)
}

func (c *BaseChainContext) GetErrors() []error {
	return c.errors
}

func (c *BaseChainContext) Get(key string) interface{} {
	return c.data[key]
}

func (c *BaseChainContext) Remove(key string) {
	delete(c.data, key)
}

func (c *BaseChainContext) HasErrors() bool {
	return len(c.errors) > 0
}
