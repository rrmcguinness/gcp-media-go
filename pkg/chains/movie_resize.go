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

package chains

import (
	"strings"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/commands"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

// The default command requires ffmpeg on the path of the running computer.
const DEFAULT_FFMPEG_COMMAND = "ffmpeg"

// The default width is the recommended size.
const DEFAULT_WIDTH = "240"

func MovieResizeChain(
	ffmpegCommand string,
	videoFormat *model.VideoFormatFilter,
	storageClient *storage.Client,
	outputBucketName string) cor.Chain {

	if storageClient == nil {
		panic("FFMPegChain requires a valid storage client")
	}

	// Ensure the FFMPegCommand is set, otherwise use the default
	if len(strings.Trim(ffmpegCommand, " ")) == 0 {
		ffmpegCommand = DEFAULT_FFMPEG_COMMAND
	}

	// Set the default width
	videoWidth := DEFAULT_WIDTH
	if videoFormat != nil {
		videoWidth = videoFormat.Width
	}

	out := &cor.BaseChain{}

	// Convert the Message to an Object
	out.AddCommand(&commands.MediaTriggerToGCSObject{})

	// Write a temp file
	out.AddCommand(&commands.GCSToTempFileCommand{Client: storageClient, TempFilePrefix: "ffmpeg-tmp-"})

	// Run FFMpeg
	out.AddCommand(&commands.FFMpegCommand{ExecutableCommand: ffmpegCommand, TargetWidth: videoWidth})

	// Write to a GCS Bucket
	out.AddCommand(&commands.GCSFileUpload{Client: *storageClient, Bucket: outputBucketName})

	return out
}
