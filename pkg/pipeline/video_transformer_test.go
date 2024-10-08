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

package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoTransformer(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current path:", err)
		return
	}

	fmt.Println("Current path:", path)

	fmt.Print("Yo")
	cmd := exec.Command("bin/ffmpeg", "-help")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.NotNil(t, out)
	fmt.Println("Output:", string(out))
}
