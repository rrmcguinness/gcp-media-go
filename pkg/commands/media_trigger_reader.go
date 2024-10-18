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

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaTriggerToGCSObject struct {
	cor.BaseCommand
}

func (c *MediaTriggerToGCSObject) Execute(context cor.Context) {
	in := context.Get(c.GetInputParam()).(string)
	var out model.GCSPubSubNotification
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		context.AddError(err)
		return
	}

	msg := &model.GCSObject{Bucket: out.Bucket, Name: out.Name, MIMEType: out.ContentType}
	context.Add(model.GetGCSObjectName(), msg)
	context.Add(c.GetOutputParam(), msg)
}
