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

package services

const (
	QRY_SEQUENCE_KNN     = "SELECT base.media_id, base.sequence_number FROM VECTOR_SEARCH(TABLE `%s`, 'embeddings', (SELECT [ %s ] as embed), top_k => %d, distance_type => 'EUCLIDEAN') ORDER BY distance asc"
	QRY_FIND_MEDIA_BY_ID = "SELECT * from `%s` WHERE id = '%s'"
	QRY_GET_SCENE        = "SELECT sequence, start, `end`, script FROM `%s`, UNNEST(scenes) as s WHERE id = '%s' and s.sequence = %d"
)
