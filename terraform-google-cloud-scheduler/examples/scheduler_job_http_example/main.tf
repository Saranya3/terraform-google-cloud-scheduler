/**
 * Copyright 2021 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

module "cloud_scheduler" {
  source = "../.."

  name             = var.name
  description      = "Sample Cloud scheduler"
  project_id       = var.project_id
  location         = "us-central1"
  schedule         = "*/4 * * * *"
  time_zone        = "UTC"
  paused           = false
  attempt_deadline = "15s"

  retry_config = {
    retry_count = 1
  }

  pubsub_target = null

  app_engine_http_target = null

  http_target = {
    http_method = "POST"
    uri         = "https://example.com/"
    body        = base64encode("{\"foo\":\"bar\"}")
    headers = {
      "Content-Type" = "application/json"
    }
  }
}
