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

package main

import (
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/gin-gonic/gin"
)

func MediaRouter(r *gin.Engine) {

	r.GET("/media", func(c *gin.Context) {
		query := c.Query("s")
		if len(query) == 0 {
			c.Status(404)
			return
		}
		results, err := state.searchService.FindScenes(query)

		if err != nil {
			c.Status(404)
			log.Println(err)
			return
		}

		out := make([]*model.Scene, 0)

		for _, r := range results {
			s, err := state.mediaService.GetScene(r.MediaId, r.SequenceNumber)
			if err != nil {
				c.Status(400)
				return
			}
			out = append(out, s)
		}
		c.JSON(200, out)
	})

	r.GET("/media/:id", func(c *gin.Context) {
		id := c.Param("id")
		out, err := state.mediaService.Get(id)
		if err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, out)
	})

	r.GET("/media/:id/scenes/:scene_id", func(c *gin.Context) {
		id := c.Param("id")
		scene_id, err := strconv.Atoi(c.Param("scene_id"))
		if err != nil {
			c.Status(400)
			return
		}
		out, err := state.mediaService.GetScene(id, scene_id)
		if err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, out)
	})
}
