// Copyright 2019-2024, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build (nodejs || all) && !xplatform_acceptance

package ints

import (
	"testing"

	"github.com/khulnasoft/codeinfra/pkg/v3/testing/integration"
)

// TestQuery creates a stack and runs a query over the stack's resource outputs.
//
//nolint:paralleltest // ProgramTest calls t.Parallel()
func TestQuery(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		// Create Codeinfra resources.
		Dir:          "step1",
		StackName:    "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2",
		Dependencies: []string{"@codeinfra/codeinfra"},
		CloudURL:     integration.MakeTempBackend(t), // Required; we hard-code the stack name
		EditDirs: []integration.EditDir{
			// Try to create resources during `codeinfra query`. This should fail.
			{
				Dir:           "step2",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: true,
			},
			// Run a query during `codeinfra query`. Should succeed.
			{
				Dir:           "step3",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: false,
			},
		},
	})
}
