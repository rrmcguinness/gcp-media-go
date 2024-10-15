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

import "os"

type BaseContext struct {
	data      map[string]interface{}
	errors    []error
	tempFiles []string
}

func NewBaseContext() Context {
	return &BaseContext{
		data:      make(map[string]interface{}),
		errors:    make([]error, 0),
		tempFiles: make([]string, 0),
	}
}

func (c *BaseContext) Close() {
	// Clean up any temp files created along the way
	for _, file := range c.GetTempFiles() {
		os.Remove(file)
	}
}

func (c *BaseContext) Add(key string, value interface{}) Context {
	c.data[key] = value
	return c
}

func (c *BaseContext) AddTempFile(file string) {
	c.tempFiles = append(c.tempFiles, file)
}

func (c *BaseContext) GetTempFiles() []string {
	return c.tempFiles
}

func (c *BaseContext) AddError(err error) {
	c.errors = append(c.errors, err)
}

func (c *BaseContext) GetErrors() []error {
	return c.errors
}

func (c *BaseContext) Get(key string) interface{} {
	return c.data[key]
}

func (c *BaseContext) Remove(key string) {
	delete(c.data, key)
}

func (c *BaseContext) HasErrors() bool {
	return len(c.errors) > 0
}
