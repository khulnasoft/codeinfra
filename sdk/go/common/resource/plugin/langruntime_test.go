// Copyright 2024, KhulnaSoft Ltd.
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

package plugin

import (
	"context"
	"errors"
	"testing"

	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"

	"google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"
)

func TestMakeExecutablePromptChoices(t *testing.T) {
	t.Parallel()

	// Not found executables come after the found ones, and have a [not found] suffix.
	choices := MakeExecutablePromptChoices("executable_that_does_not_exist_in_path", "ls")
	require.Equal(t, 2, len(choices))
	require.Equal(t, choices[0].StringValue, "ls")
	require.Equal(t, choices[0].DisplayName, "ls")
	require.Equal(t, choices[1].StringValue, "executable_that_does_not_exist_in_path")
	require.Equal(t, choices[1].DisplayName, "executable_that_does_not_exist_in_path [not found]")

	// Executables are not reordered within their group.
	choices = MakeExecutablePromptChoices("ls", "cat", "zzz_does_not_exist", "aaa_does_not_exist")
	require.Equal(t, choices[0].StringValue, "ls")
	require.Equal(t, choices[0].DisplayName, "ls")
	require.Equal(t, choices[1].StringValue, "cat")
	require.Equal(t, choices[1].DisplayName, "cat")
	require.Equal(t, choices[2].StringValue, "zzz_does_not_exist")
	require.Equal(t, choices[2].DisplayName, "zzz_does_not_exist [not found]")
	require.Equal(t, choices[3].StringValue, "aaa_does_not_exist")
	require.Equal(t, choices[3].DisplayName, "aaa_does_not_exist [not found]")
}

type MockLanguageRuntimeClient struct {
	RunPluginF (func(ctx context.Context, info *codeinfrarpc.RunPluginRequest,
	) (codeinfrarpc.LanguageRuntime_RunPluginClient, error))
}

func (m *MockLanguageRuntimeClient) RunPlugin(
	ctx context.Context, info *codeinfrarpc.RunPluginRequest, _ ...grpc.CallOption,
) (codeinfrarpc.LanguageRuntime_RunPluginClient, error) {
	return m.RunPluginF(ctx, info)
}

func (m *MockLanguageRuntimeClient) GetRequiredPackages(
	ctx context.Context, in *codeinfrarpc.GetRequiredPackagesRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GetRequiredPackagesResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GetRequiredPlugins(
	ctx context.Context, in *codeinfrarpc.GetRequiredPluginsRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GetRequiredPluginsResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) Run(
	ctx context.Context, in *codeinfrarpc.RunRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.RunResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GetPluginInfo(
	ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption,
) (*codeinfrarpc.PluginInfo, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) InstallDependencies(
	ctx context.Context, in *codeinfrarpc.InstallDependenciesRequest, opts ...grpc.CallOption,
) (codeinfrarpc.LanguageRuntime_InstallDependenciesClient, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) RuntimeOptionsPrompts(
	ctx context.Context, in *codeinfrarpc.RuntimeOptionsRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.RuntimeOptionsResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) About(
	ctx context.Context, in *codeinfrarpc.AboutRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.AboutResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GetProgramDependencies(
	ctx context.Context, in *codeinfrarpc.GetProgramDependenciesRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GetProgramDependenciesResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GenerateProgram(
	ctx context.Context, in *codeinfrarpc.GenerateProgramRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GenerateProgramResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GenerateProject(
	ctx context.Context, in *codeinfrarpc.GenerateProjectRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GenerateProjectResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GeneratePackage(
	ctx context.Context, in *codeinfrarpc.GeneratePackageRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.GeneratePackageResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) Pack(
	ctx context.Context, in *codeinfrarpc.PackRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.PackResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) Handshake(
	ctx context.Context, req *codeinfrarpc.LanguageHandshakeRequest, opts ...grpc.CallOption,
) (*codeinfrarpc.LanguageHandshakeResponse, error) {
	panic("not implemented")
}

func TestRunPluginPassesCorrectPwd(t *testing.T) {
	t.Parallel()

	returnErr := errors.New("erroring so we don't need to implement the whole thing")
	mockLanguageRuntime := &MockLanguageRuntimeClient{
		RunPluginF: func(
			ctx context.Context, info *codeinfrarpc.RunPluginRequest,
		) (codeinfrarpc.LanguageRuntime_RunPluginClient, error) {
			require.Equal(t, "/tmp", info.Pwd)
			return nil, returnErr
		},
	}

	pCtx, err := NewContext(nil, nil, nil, nil, "", nil, false, nil)
	require.NoError(t, err)
	host := &langhost{
		ctx:     pCtx,
		runtime: "go",
		plug:    nil,
		client:  mockLanguageRuntime,
	}

	// Test that the plugin is run with the correct working directory.
	_, _, _, err = host.RunPlugin(RunPluginInfo{
		WorkingDirectory: "/tmp",
	})
	require.Equal(t, returnErr, err)
}
