// Copyright 2016-2024, KhulnaSoft Ltd.
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

package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/codeinfra/pkg/v3/testing/integration"
	ptesting "github.com/khulnasoft/codeinfra/sdk/v3/go/common/testing"
)

func TestPreviewOnlyFlag(t *testing.T) {
	t.Run("PreviewOnlyRefresh", func(t *testing.T) {
		t.Parallel()

		e := ptesting.NewEnvironment(t)
		defer e.DeleteIfNotFailed()

		integration.CreateBasicCodeinfraRepo(e)
		e.ImportDirectory("integration/single_resource")
		e.RunCommand("yarn", "link", "@codeinfra/codeinfra")
		e.RunCommand("yarn", "install")
		e.SetBackend(e.LocalURL())
		e.RunCommand("codeinfra", "stack", "init", "foo")
		e.RunCommand("codeinfra", "up", "--skip-preview", "--yes")

		// Try some invalid flag combinations.
		_, stderr := e.RunCommandExpectError("codeinfra", "refresh", "--preview-only", "--yes")
		assert.Equal(t,
			"error: --yes and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "refresh", "--skip-preview", "--preview-only")
		assert.Equal(t,
			"error: --skip-preview and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "refresh", "--non-interactive")
		assert.Equal(t,
			"error: --yes or --skip-preview or --preview-only must be passed in to proceed when "+
				"running in non-interactive mode",
			strings.Trim(stderr, "\r\n"))

		// Now try just the flag.
		stdout, _ := e.RunCommand("codeinfra", "refresh", "--preview-only")
		assert.NotContains(t, stdout, "Do you want to perform this refresh?")
		// Make sure it works with --non-interactive too.
		e.RunCommand("codeinfra", "refresh", "--preview-only", "--non-interactive")

		e.RunCommand("codeinfra", "destroy", "--skip-preview", "--yes")
		// Remove the stack.
		e.RunCommand("codeinfra", "stack", "rm", "foo", "--yes")
	})

	t.Run("PreviewOnlyDestroy", func(t *testing.T) {
		t.Parallel()

		e := ptesting.NewEnvironment(t)
		defer e.DeleteIfNotFailed()

		integration.CreateBasicCodeinfraRepo(e)
		e.ImportDirectory("integration/single_resource")
		e.RunCommand("yarn", "link", "@codeinfra/codeinfra")
		e.RunCommand("yarn", "install")
		e.SetBackend(e.LocalURL())
		e.RunCommand("codeinfra", "stack", "init", "foo")
		e.RunCommand("codeinfra", "up", "--skip-preview", "--yes")

		// Try some invalid flag combinations.
		_, stderr := e.RunCommandExpectError("codeinfra", "destroy", "--preview-only", "--yes")
		assert.Equal(t,
			"error: --yes and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "destroy", "--skip-preview", "--preview-only")
		assert.Equal(t,
			"error: --skip-preview and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "destroy", "--non-interactive")
		assert.Equal(t,
			"error: --yes or --skip-preview or --preview-only must be passed in to proceed when running in non-interactive mode",
			strings.Trim(stderr, "\r\n"))

		// Now try just the flag.
		stdout, _ := e.RunCommand("codeinfra", "destroy", "--preview-only")
		assert.NotContains(t, stdout, "Do you want to perform this destroy?")
		assert.NotContains(t, stdout, "The resources in the stack have been deleted")
		// Make sure it works with --non-interactive too.
		e.RunCommand("codeinfra", "destroy", "--preview-only", "--non-interactive")

		e.RunCommand("codeinfra", "destroy", "--skip-preview", "--yes")
		// Remove the stack.
		e.RunCommand("codeinfra", "stack", "rm", "foo", "--yes")
	})

	t.Run("PreviewOnlyImport", func(t *testing.T) {
		t.Parallel()

		e := ptesting.NewEnvironment(t)
		defer e.DeleteIfNotFailed()

		integration.CreateBasicCodeinfraRepo(e)
		e.SetBackend(e.LocalURL())
		e.RunCommand("codeinfra", "stack", "init", "foo")

		// Make sure random is installed
		e.RunCommand("codeinfra", "plugin", "install", "resource", "random", "4.13.0")

		// Try some invalid flag combinations.
		_, stderr := e.RunCommandExpectError("codeinfra", "import", "random:index/randomId:RandomId",
			"identifier", "p-9hUg", "--preview-only", "--yes")
		assert.Equal(t,
			"error: --yes and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "import", "random:index/randomId:RandomId",
			"identifier", "p-9hUg", "--skip-preview", "--preview-only")
		assert.Equal(t,
			"error: --skip-preview and --preview-only cannot be used together",
			strings.Trim(stderr, "\r\n"))
		_, stderr = e.RunCommandExpectError("codeinfra", "import", "random:index/randomId:RandomId",
			"identifier", "p-9hUg", "--non-interactive")
		assert.Equal(t,
			"error: --yes or --skip-preview or --preview-only must be passed in to proceed when running in non-interactive mode",
			strings.Trim(stderr, "\r\n"))

		// Now try just the flag.
		stdout, _ := e.RunCommand("codeinfra", "import", "random:index/randomId:RandomId",
			"identifier", "p-9hUg", "--preview-only")
		assert.NotContains(t, stdout, "Do you want to perform this import?")
		// Make sure it works with --non-interactive too.
		e.RunCommand("codeinfra", "import", "random:index/randomId:RandomId",
			"identifier", "p-9hUg", "--preview-only", "--non-interactive")

		// Remove the stack.
		e.RunCommand("codeinfra", "stack", "rm", "foo", "--yes")
	})
}
