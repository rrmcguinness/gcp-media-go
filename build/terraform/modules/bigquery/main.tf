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

terraform {
  required_version = ">= 0.12"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.5.0"
    }
  }
}

resource "google_bigquery_dataset" "media_ds" {
  dataset_id                  = "media_ds"
  description                 = "Media data source for movie object"
  location                    = "US"
  delete_contents_on_destroy = false

  default_partition_expiration_ms = 3600000
  default_table_expiration_ms     = 3600000
  max_time_travel_hours = 96

  labels = {
    env = "default"
  }
}

resource "google_bigquery_table" "media_ds_movies" {
  dataset_id = google_bigquery_dataset.media_ds.dataset_id
  table_id   = "movies"

  schema = <<EOF
[
    {
        "name": "title",
        "type": "STRING",
        "mode": "REQUIRED"
    },
    {
        "name": "summary",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "director",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "release_year",
        "type": "INTEGER",
        "mode": "NULLABLE"
    },
    {
        "name": "genre",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "rating",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "cast",
        "type": "RECORD",
        "mode": "REPEATED",
        "fields": [
            {
                "name": "character_name",
                "type": "STRING",
                "mode": "NULLABLE"
            },
            {
                "name": "actor_name",
                "type": "STRING",
                "mode": "NULLABLE"
            }
        ]
    },
    {
        "name": "scenes",
        "type": "RECORD",
        "mode": "REPEATED",
        "fields": [
            {
                "name": "sequence",
                "type": "INTEGER",
                "mode": "REQUIRED"
            },
            {
                "name": "time_span",
                "type": "RECORD",
                "mode": "NULLABLE",
                "fields": [
                    {
                        "name": "start",
                        "type": "STRING",
                        "mode": "NULLABLE"
                    },
                    {
                        "name": "end",
                        "type": "STRING",
                        "mode": "NULLABLE"
                    }
                ]
            },
            {
                "name": "summary",
                "type": "STRING",
                "mode": "NULLABLE"
            },
            {
                "name": "dialog",
                "type": "RECORD",
                "mode": "REPEATED",
                "fields": [
                    {
                        "name": "character_name",
                        "type": "STRING",
                        "mode": "NULLABLE"
                    },
                    {
                        "name": "dialog",
                        "type": "STRING",
                        "mode": "NULLABLE"
                    }
                ]
            }
        ]
    }
]
EOF
}