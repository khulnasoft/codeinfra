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

package env

import (
	"github.com/spf13/cobra"

	"github.com/pulumi/esc/cmd/esc/cli"
	escWorkspace "github.com/pulumi/esc/cmd/esc/cli/workspace"
	"github.com/khulnasoft/codeinfra/pkg/v3/backend/httpstate"
	"github.com/khulnasoft/codeinfra/pkg/v3/backend/httpstate/client"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
)

func NewEnvCmd() *cobra.Command {
	escCLI := cli.New(&cli.Options{
		ParentPath:      "codeinfra",
		Colors:          cmdutil.GetGlobalColorization(),
		Login:           httpstate.NewLoginManager(),
		CodeinfraWorkspace: escWorkspace.DefaultCodeinfraWorkspace(),
		UserAgent:       client.UserAgent(),
	})

	// Add the `env` command to the root.
	envCommand := escCLI.Commands()[0]
	return envCommand
}
