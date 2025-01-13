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

package ints

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	ptesting "github.com/khulnasoft/codeinfra/sdk/v3/go/common/testing"
)

// TestPolicyWithConfig runs integration tests against the policy pack in the policy_pack_w_config
// directory using version 0.4.1-dev of the codeinfra/policy sdk.
//
//nolint:paralleltest // mutates environment variables
func TestPolicyWithConfig(t *testing.T) {
	t.Skip("Skip test that is causing unrelated tests to fail - codeinfra/codeinfra#4149")

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	// Confirm we have credentials.
	if os.Getenv("CODEINFRA_ACCESS_TOKEN") == "" {
		t.Fatal("CODEINFRA_ACCESS_TOKEN not found, aborting tests.")
	}

	name, _ := e.RunCommand("codeinfra", "whoami")
	orgName := strings.TrimSpace(name)
	// Pack and push a Policy Pack for the organization.
	policyPackName := fmt.Sprintf("%s-%x", "test-policy-pack", time.Now().UnixNano())
	e.ImportDirectory("policy_pack_w_config")
	e.RunCommand("yarn", "install")
	t.Setenv("TEST_POLICY_PACK", policyPackName)

	// Publish the Policy Pack twice.
	publishPolicyPackWithVersion(e, orgName, `"0.0.1"`)
	publishPolicyPackWithVersion(e, orgName, `"0.0.2"`)

	// Check the policy ls commands.
	packsOutput, _ := e.RunCommand("codeinfra", "policy", "ls", "--json")
	var packs []policyPacksJSON
	assertJSON(e, packsOutput, &packs)

	groupsOutput, _ := e.RunCommand("codeinfra", "policy", "group", "ls", "--json")
	var groups []policyGroupsJSON
	assertJSON(e, groupsOutput, &groups)

	// Enable, Disable and then Delete the Policy Pack.
	e.RunCommand("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "0.0.1")

	// Validate Policy Pack Configuration.
	e.RunCommand("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/valid-config.json", "0.0.1")
	// Valid config, but no version specified.
	e.RunCommandExpectError("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/config.json")
	// Invalid configs
	e.RunCommandExpectError("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/invalid-config.json", "0.0.1")
	// Invalid - missing required property.
	e.RunCommandExpectError("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/invalid-required-prop.json", "0.0.1")
	// Required config flag not present.
	e.RunCommandExpectError("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName))
	e.RunCommandExpectError("codeinfra", "policy", "validate-config", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config", "0.0.1")

	// Enable Policy Pack with Configuration.
	e.RunCommand("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/valid-config.json", "0.0.1")
	e.RunCommandExpectError("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName),
		"--config=configs/invalid-config.json", "0.0.1")

	// Disable Policy Pack specifying version.
	e.RunCommand("codeinfra", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName), "--version=0.0.1")

	// Enable and Disable without specifying the version number.
	e.RunCommand("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "latest")
	e.RunCommand("codeinfra", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName))

	e.RunCommand("codeinfra", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "0.0.1")
	e.RunCommand("codeinfra", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "all")
}

// TestPolicyWithoutConfig runs integration tests against the policy pack in the policy_pack_w_config
// directory. This tests against version 0.4.0 of the codeinfra/policy sdk, prior to policy config being supported.
//
//nolint:paralleltest // mutates environment variables
func TestPolicyWithoutConfig(t *testing.T) {
	t.Skip("Skip test that is causing unrelated tests to fail - codeinfra/codeinfra#4149")

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	// Confirm we have credentials.
	if os.Getenv("CODEINFRA_ACCESS_TOKEN") == "" {
		t.Fatal("CODEINFRA_ACCESS_TOKEN not found, aborting tests.")
	}

	name, _ := e.RunCommand("codeinfra", "whoami")
	orgName := strings.TrimSpace(name)

	// Pack and push a Policy Pack for the organization.
	policyPackName := fmt.Sprintf("%s-%x", "test-policy-pack", time.Now().UnixNano())
	e.ImportDirectory("policy_pack_wo_config")
	e.RunCommand("yarn", "install")
	t.Setenv("TEST_POLICY_PACK", policyPackName)

	// Publish the Policy Pack twice.
	e.RunCommand("codeinfra", "policy", "publish", orgName)
	e.RunCommand("codeinfra", "policy", "publish", orgName)

	// Check the policy ls commands.
	packsOutput, _ := e.RunCommand("codeinfra", "policy", "ls", "--json")
	var packs []policyPacksJSON
	assertJSON(e, packsOutput, &packs)

	groupsOutput, _ := e.RunCommand("codeinfra", "policy", "group", "ls", "--json")
	var groups []policyGroupsJSON
	assertJSON(e, groupsOutput, &groups)

	// Enable, Disable and then Delete the Policy Pack.
	e.RunCommand("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "1")
	e.RunCommand("codeinfra", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName), "--version=1")

	// Enable and Disable without specifying the version number.
	e.RunCommand("codeinfra", "policy", "enable", fmt.Sprintf("%s/%s", orgName, policyPackName), "latest")
	e.RunCommand("codeinfra", "policy", "disable", fmt.Sprintf("%s/%s", orgName, policyPackName))

	e.RunCommand("codeinfra", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "1")
	e.RunCommand("codeinfra", "policy", "rm", fmt.Sprintf("%s/%s", orgName, policyPackName), "all")
}

type policyPacksJSON struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type policyGroupsJSON struct {
	Name           string `json:"name"`
	Default        bool   `json:"default"`
	NumPolicyPacks int    `json:"numPolicyPacks"`
	NumStacks      int    `json:"numStacks"`
}

//nolint:unused // Used by skipped test
func assertJSON(e *ptesting.Environment, out string, respObj interface{}) {
	err := json.Unmarshal([]byte(out), &respObj)
	if err != nil {
		e.Errorf("unable to unmarshal %v", out)
	}
}

// publishPolicyPackWithVersion updates the version in package.json so we can
// dynamically publish different versions for testing.
//
//nolint:unused // Used by skipped test
func publishPolicyPackWithVersion(e *ptesting.Environment, orgName, version string) {
	cmd := fmt.Sprintf(`sed 's/{ policyVersion }/%s/g' package.json.tmpl | tee package.json`, version)
	e.RunCommand("bash", "-c", cmd)
	e.RunCommand("codeinfra", "policy", "publish", orgName)
}
