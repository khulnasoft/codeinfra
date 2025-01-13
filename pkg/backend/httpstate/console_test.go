// Copyright 2016-2018, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpstate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // sets env var, must be run in isolation
func TestConsoleURL(t *testing.T) {
	//nolint:paralleltest // sets env var, must be run in isolation
	t.Run("HonorEnvVar", func(t *testing.T) {
		// Honor the CODEINFRA_CONSOLE_DOMAIN environment variable.
		t.Setenv("CODEINFRA_CONSOLE_DOMAIN", "codeinfra-console.contoso.com")
		assert.Equal(t,
			"https://codeinfra-console.contoso.com/1/2",
			cloudConsoleURL("https://api.codeinfra.contoso.com", "1", "2"))

		// Unset the variable, confirm the "standard behavior" where we
		// replace "api." with "app.".
		os.Unsetenv("CODEINFRA_CONSOLE_DOMAIN")
		assert.Equal(t,
			"https://app.codeinfra.contoso.com/1/2",
			cloudConsoleURL("https://api.codeinfra.contoso.com", "1", "2"))
	})

	t.Run("CloudURLUsingStandardPattern", func(t *testing.T) {
		assert.Equal(t,
			"https://app.codeinfra.com/codeinfra-bot/my-stack",
			cloudConsoleURL("https://api.codeinfra.com", "codeinfra-bot", "my-stack"))

		assert.Equal(t,
			"http://app.codeinfra.example.com/codeinfra-bot/my-stack",
			cloudConsoleURL("http://api.codeinfra.example.com", "codeinfra-bot", "my-stack"))
	})

	t.Run("LocalDevelopment", func(t *testing.T) {
		assert.Equal(t,
			"http://localhost:3000/codeinfra-bot/my-stack",
			cloudConsoleURL("http://localhost:8080", "codeinfra-bot", "my-stack"))
	})

	t.Run("ConsoleDomainUnknown", func(t *testing.T) {
		assert.Equal(t, "", cloudConsoleURL("https://codeinfra.example.com", "codeinfra-bot", "my-stack"))
		assert.Equal(t, "", cloudConsoleURL("not-even-a-real-url", "codeinfra-bot", "my-stack"))
	})
}
