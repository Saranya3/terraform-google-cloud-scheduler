// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multiple_buckets

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSimpleExample(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)
	projectID := example.GetTFSetupStringOutput("project_id")
	jobName := example.GetTFSetupStringOutput("name")

	example.DefineVerify(func(assert *assert.Assertions) {
		example.DefaultVerify(assert)

		services := gcloud.Run(t, "services list", gcloud.WithCommonArgs([]string{"--project", projectID, "--format", "json"})).Array()
		match := utils.GetFirstMatchResult(t, services, "config.name", "cloudscheduler.googleapis.com")
		assert.Equal("ENABLED", match.Get("state").String(), "Cloud scheduler service should be enabled")

		id := example.GetStringOutput("id")
		assert.NotEmpty(id, "Terraform output 'id' should not be empty")

		expectedIDRegex := fmt.Sprintf("^projects/%s/locations/us-central1/jobs/%s$", projectID, jobName)
		assert.Regexp(expectedIDRegex, id, "Terraform output 'id' (%s) should match regex: %s", id, expectedIDRegex)

		state := example.GetStringOutput("state")
		assert.Equal("ENABLED", state, "Terraform output 'state' should be ENABLED")
	})
	example.Test()

}
