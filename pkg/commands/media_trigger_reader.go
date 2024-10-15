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

package commands

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaTriggerToGCSObject struct {
	cor.BaseCommand
}

func (c *MediaTriggerToGCSObject) IsExecutable(context cor.Context) bool {
	return context != nil && context.Get(c.GetInputParam()) != nil
}

func (c *MediaTriggerToGCSObject) Execute(context cor.Context) {
	in := context.Get(c.GetInputParam()).(string)
	var out model.TriggerMediaWrite
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		context.AddError(err)
		return
	}

	fmt.Printf("-------------- %s", c.GetOutputParam())

	msg := &model.GCSObject{Bucket: out.Bucket, Name: out.Name, MIMEType: out.ContentType}
	context.Add("__GCS_OBJ__", msg)
	context.Add(c.GetOutputParam(), msg)
}
