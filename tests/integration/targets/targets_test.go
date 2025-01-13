// Copyright 2020-2024, KhulnaSoft Ltd.
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

//go:build !xplatform_acceptance

package ints

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	ptesting "github.com/khulnasoft/codeinfra/sdk/v3/go/common/testing"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/tokens"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/contract"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestUntargetedCreateDuringTargetedUpdate(t *testing.T) {
	t.Skip() // TODO[codeinfra/codeinfra#4149]
	t.Parallel()

	if os.Getenv("CODEINFRA_ACCESS_TOKEN") == "" {
		t.Skipf("Skipping: CODEINFRA_ACCESS_TOKEN is not set")
	}

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	stackName, err := resource.NewUniqueHex("test-", 8, -1)
	contract.AssertNoErrorf(err, "resource.NewUniqueHex should not fail with no maximum length is set")

	e.ImportDirectory("untargeted_create")
	e.RunCommand("yarn", "link", "@codeinfra/codeinfra")
	e.RunCommand("codeinfra", "stack", "init", stackName)
	e.RunCommand("codeinfra", "up", "--non-interactive", "--skip-preview", "--yes")
	urn, _ := e.RunCommand("codeinfra", "stack", "output", "urn")

	if err := fsutil.CopyFile(
		path.Join(e.RootPath, "untargeted_create", "index.ts"),
		path.Join("untargeted_create", "step1", "index.ts"), nil); err != nil {
		t.Fatalf("error copying index.ts file: %v", err)
	}

	e.RunCommand("codeinfra", "up", "--target", strings.TrimSpace(urn), "--non-interactive", "--skip-preview", "--yes")
	e.RunCommand("codeinfra", "refresh", "--non-interactive", "--yes")

	e.RunCommand("codeinfra", "destroy", "--skip-preview", "--non-interactive", "--yes")
	e.RunCommand("codeinfra", "stack", "rm", "--yes")
}

func TestDeleteManyTargets(t *testing.T) {
	t.Parallel()

	if os.Getenv("CODEINFRA_ACCESS_TOKEN") == "" {
		t.Skipf("Skipping: CODEINFRA_ACCESS_TOKEN is not set")
	}

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	// First just spin up the project.
	projName := "delete_targets_many_deps"
	stackName, err := resource.NewUniqueHex("test-", 8, -1)
	contract.AssertNoErrorf(err, "resource.NewUniqueHex should not fail with no maximum length is set")
	e.ImportDirectory(projName)
	e.RunCommand("codeinfra", "stack", "init", stackName)
	e.RunCommand("yarn", "link", "@codeinfra/codeinfra")
	e.RunCommand("yarn", "install")
	e.RunCommand("codeinfra", "up", "--non-interactive", "--skip-preview", "--yes")

	// Create a handy mkURN func to create URNs for dynamic resources in this project/stack.
	resourceType := tokens.Type("codeinfra-nodejs:dynamic:Resource")
	mkURNStr := func(resourceName string, parentType tokens.Type) string {
		return string(resource.NewURN(
			tokens.QName(stackName), tokens.PackageName(projName), parentType, resourceType, resourceName))
	}

	// Attempt to destroy the root-most node. It should fail and the error text should
	// mention every one of the nodes in the entire graph (since they all transitively depend on a).
	stdout, _ := e.RunCommandExpectError("codeinfra", "destroy", "--skip-preview", "--yes", "--non-interactive",
		"--target", mkURNStr("a", ""))
	assert.Contains(t, stdout, mkURNStr("b", ""))
	assert.Contains(t, stdout, mkURNStr("c", ""))
	assert.Contains(t, stdout, mkURNStr("d", ""))
	assert.Contains(t, stdout, mkURNStr("e", ""))
	assert.Contains(t, stdout, mkURNStr("f", resourceType))
	assert.Contains(t, stdout, mkURNStr("g", resourceType))

	// Destroy the leaf-most node. This should work just fine.
	e.RunCommand("codeinfra", "destroy", "--skip-preview", "--yes", "--non-interactive",
		"--target", mkURNStr("h", tokens.Type(fmt.Sprintf("%[1]s$%[1]s", resourceType))))

	// Finally, go back and try to delete the root-most node, but clean up the transitive closure.
	e.RunCommand("codeinfra", "destroy", "--skip-preview", "--yes", "--non-interactive",
		"--target", mkURNStr("a", ""), "--target-dependents")

	// Finally clean up the entire stack.
	e.RunCommand("codeinfra", "destroy", "--skip-preview", "--yes", "--non-interactive")
	e.RunCommand("codeinfra", "stack", "rm", "--yes")
}
