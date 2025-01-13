// Copyright 2016-2022, KhulnaSoft Ltd.
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

//go:build all

package ints

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/codeinfra/pkg/v3/testing/integration"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource/config"
	ptesting "github.com/khulnasoft/codeinfra/sdk/v3/go/common/testing"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/workspace"
)

// TestConfigSave ensures that config commands in the Codeinfra CLI work as expected.
func TestConfigSave(t *testing.T) {
	t.Parallel()
	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	// Initialize an empty stack.
	path := filepath.Join(e.RootPath, "Codeinfra.yaml")
	project := workspace.Project{
		Name:    "testing-config",
		Runtime: workspace.NewProjectRuntimeInfo("nodejs", nil),
	}

	err := project.Save(path)
	assert.NoError(t, err)
	e.RunCommand("codeinfra", "login", "--cloud-url", e.LocalURL())
	e.RunCommand("codeinfra", "stack", "init", "testing-2")
	e.RunCommand("codeinfra", "stack", "init", "testing-1")

	// Now configure and save a few different things:
	e.RunCommand("codeinfra", "config", "set", "configA", "value1")
	e.RunCommand("codeinfra", "config", "set", "configB", "value2", "--stack", "testing-2")

	e.RunCommand("codeinfra", "stack", "select", "testing-2")

	e.RunCommand("codeinfra", "config", "set", "configD", "value4")
	e.RunCommand("codeinfra", "config", "set", "configC", "value3", "--stack", "testing-1")

	// Now read back the config using the CLI:
	{
		stdout, _ := e.RunCommand("codeinfra", "config", "get", "configB")
		assert.Equal(t, "value2\n", stdout)
	}
	{
		// the config in a different stack, so this should error.
		stdout, stderr := e.RunCommandExpectError("codeinfra", "config", "get", "configA")
		assert.Equal(t, "", stdout)
		assert.NotEqual(t, "", stderr)
	}
	{
		// but selecting the stack should let you see it
		stdout, _ := e.RunCommand("codeinfra", "config", "get", "configA", "--stack", "testing-1")
		assert.Equal(t, "value1\n", stdout)
	}

	// Finally, check that the stack file contains what we expected.
	validate := func(k string, v string, cfg config.Map) {
		key, err := config.ParseKey("testing-config:config:" + k)
		assert.NoError(t, err)
		d, ok := cfg[key]
		assert.True(t, ok, "config key %v should be set", k)
		dv, err := d.Value(nil)
		assert.NoError(t, err)
		assert.Equal(t, v, dv)
	}

	testStack1, err := workspace.LoadProjectStack(&project, filepath.Join(e.CWD, "Codeinfra.testing-1.yaml"))
	assert.NoError(t, err)
	testStack2, err := workspace.LoadProjectStack(&project, filepath.Join(e.CWD, "Codeinfra.testing-2.yaml"))
	assert.NoError(t, err)

	assert.Equal(t, 2, len(testStack1.Config))
	assert.Equal(t, 2, len(testStack2.Config))

	validate("configA", "value1", testStack1.Config)
	validate("configC", "value3", testStack1.Config)

	validate("configB", "value2", testStack2.Config)
	validate("configD", "value4", testStack2.Config)

	e.RunCommand("codeinfra", "stack", "rm", "--yes")
}

func TestRotatePassphrase(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	e.ImportDirectory("rotate_passphrase")
	e.RunCommand("codeinfra", "login", "--cloud-url", e.LocalURL())

	e.RunCommand("codeinfra", "stack", "init", "dev")
	e.RunCommand("codeinfra", "up", "--skip-preview", "--yes")

	e.RunCommand("codeinfra", "config", "set", "--secret", "foo", "bar")

	e.SetEnvVars("CODEINFRA_TEST_PASSPHRASE=true")
	e.Stdin = strings.NewReader("qwerty\nqwerty\n")
	e.RunCommand("codeinfra", "stack", "change-secrets-provider", "passphrase")

	e.Stdin, e.Passphrase = nil, "qwerty"
	e.RunCommand("codeinfra", "config", "get", "foo")
}

//nolint:paralleltest // uses parallel programtest
func TestJSONOutputWithStreamingPreview(t *testing.T) {
	stdout := &bytes.Buffer{}

	// Test with env var for streaming preview (should *not* print previewSummary).
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          filepath.Join("stack_outputs", "nodejs"),
		Dependencies: []string{"@codeinfra/codeinfra"},
		Stdout:       stdout,
		Verbose:      true,
		JSONOutput:   true,
		Env:          []string{"CODEINFRA_ENABLE_STREAMING_JSON_PREVIEW=1"},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			output := stdout.String()

			// Check that the previewSummary is *not* present.
			assert.NotRegexp(t, previewSummaryRegex, output)

			// Check that each event present in the event stream is also in stdout.
			for _, evt := range stack.Events {
				assertOutputContainsEvent(t, evt, output)
			}
		},
	})
}

func TestPassphrasePrompting(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	e.NoPassphrase = true
	// Setting CODEINFRA_TEST_PASSPHRASE allows prompting (reading from stdin)
	// even though the test won't be interactive.
	e.SetEnvVars("CODEINFRA_TEST_PASSPHRASE=true")

	e.RunCommand("codeinfra", "login", "--cloud-url", e.LocalURL())

	e.Stdin = strings.NewReader("qwerty\nqwerty\n")
	e.RunCommand("codeinfra", "new", "go",
		"--name", "pphraseprompt",
		"--description", "A project that tests passphrase prompts",
		"--stack", "dev",
		"--secrets-provider", "passphrase",
		"--yes",
		"--force")

	e.Stdin = strings.NewReader("qwerty\n")
	e.RunCommand("codeinfra", "up", "--stack", "dev", "--skip-preview", "--yes")

	e.Stdin = strings.NewReader("qwerty\n")
	e.RunCommand("codeinfra", "stack", "export", "--stack", "dev", "--file", "stack.json")

	e.Stdin = strings.NewReader("qwerty\n")
	e.RunCommand("codeinfra", "stack", "import", "--stack", "dev", "--file", "stack.json")

	e.Stdin = strings.NewReader("qwerty\n")
	e.RunCommand("codeinfra", "destroy", "--stack", "dev", "--skip-preview", "--yes")
}
