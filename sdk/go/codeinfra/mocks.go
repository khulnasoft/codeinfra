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

package codeinfra

import (
	"fmt"
	"log"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource/plugin"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/tokens"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"
)

type MockResourceMonitor interface {
	Call(args MockCallArgs) (resource.PropertyMap, error)
	NewResource(args MockResourceArgs) (string, resource.PropertyMap, error)
}

func WithMocks(project, stack string, mocks MockResourceMonitor) RunOption {
	return func(info *RunInfo) {
		info.Project, info.Stack, info.Mocks = project, stack, mocks
	}
}

func WithMocksWithOrganization(organization, project, stack string, mocks MockResourceMonitor) RunOption {
	return func(info *RunInfo) {
		info.Project, info.Stack, info.Mocks, info.Organization = project, stack, mocks, organization
	}
}

// MockCallArgs is used to construct a call Mock
type MockCallArgs struct {
	// Token indicates which function is being called. This token is of the form "package:module:function".
	Token string
	// Args are the arguments provided to the function call.
	Args resource.PropertyMap
	// Provider is the identifier of the provider instance being used to make the call.
	Provider string
}

// MockResourceArgs is a used to construct a newResource Mock
type MockResourceArgs struct {
	// TypeToken is the token that indicates which resource type is being constructed. This token
	// is of the form "package:module:type".
	TypeToken string
	// Name is the logical name of the resource instance.
	Name string
	// Inputs are the inputs for the resource.
	Inputs resource.PropertyMap
	// Provider is the identifier of the provider instance being used to manage this resource.
	Provider string
	// ID is the physical identifier of an existing resource to read or import.
	ID string
	// Custom specifies whether or not the resource is Custom (i.e. managed by a resource provider).
	Custom bool
	// Full register RPC call, if available.
	RegisterRPC *codeinfrarpc.RegisterResourceRequest
	// Full read RPC call, if available
	ReadRPC *codeinfrarpc.ReadResourceRequest
}

type mockMonitor struct {
	project   string
	stack     string
	mocks     MockResourceMonitor
	resources sync.Map // map[string]resource.PropertyMap
}

func (m *mockMonitor) newURN(parent, typ, name string) string {
	parentType := tokens.Type("")
	if parentURN := resource.URN(parent); parentURN != "" && parentURN.QualifiedType() != resource.RootStackType {
		parentType = parentURN.QualifiedType()
	}

	return string(resource.NewURN(tokens.QName(m.stack), tokens.PackageName(m.project), parentType, tokens.Type(typ),
		name))
}

func (m *mockMonitor) SupportsFeature(ctx context.Context, in *codeinfrarpc.SupportsFeatureRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.SupportsFeatureResponse, error) {
	id := in.GetId()

	// Support for "outputValues" is deliberately disabled for the mock monitor so
	// instances of `Output` don't show up in `MockResourceArgs` Inputs.
	hasSupport := id != "outputValues"

	return &codeinfrarpc.SupportsFeatureResponse{
		HasSupport: hasSupport,
	}, nil
}

func (m *mockMonitor) Invoke(ctx context.Context, in *codeinfrarpc.ResourceInvokeRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.InvokeResponse, error) {
	args, err := plugin.UnmarshalProperties(in.GetArgs(), plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: true,
	})
	if err != nil {
		return nil, err
	}

	if in.GetTok() == "codeinfra:codeinfra:getResource" {
		urn := args["urn"].StringValue()
		registeredResourceV, ok := m.resources.Load(urn)
		if !ok {
			return nil, fmt.Errorf("unknown resource %s", urn)
		}
		registeredResource := registeredResourceV.(resource.PropertyMap)
		result, err := plugin.MarshalProperties(registeredResource, plugin.MarshalOptions{
			KeepSecrets:   true,
			KeepResources: true,
		})
		if err != nil {
			return nil, err
		}
		return &codeinfrarpc.InvokeResponse{
			Return: result,
		}, nil
	}
	resultV, err := m.mocks.Call(MockCallArgs{
		Token:    in.GetTok(),
		Args:     args,
		Provider: in.GetProvider(),
	})
	if err != nil {
		return nil, err
	}

	result, err := plugin.MarshalProperties(resultV, plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: in.GetAcceptResources(),
	})
	if err != nil {
		return nil, err
	}

	return &codeinfrarpc.InvokeResponse{
		Return: result,
	}, nil
}

func (m *mockMonitor) StreamInvoke(ctx context.Context, in *codeinfrarpc.ResourceInvokeRequest,
	opts ...grpc.CallOption,
) (codeinfrarpc.ResourceMonitor_StreamInvokeClient, error) {
	panic("not implemented")
}

func (m *mockMonitor) Call(ctx context.Context, in *codeinfrarpc.ResourceCallRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.CallResponse, error) {
	panic("not implemented")
}

func (m *mockMonitor) ReadResource(ctx context.Context, in *codeinfrarpc.ReadResourceRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.ReadResourceResponse, error) {
	stateIn, err := plugin.UnmarshalProperties(in.GetProperties(), plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: true,
	})
	if err != nil {
		return nil, err
	}

	id, state, err := m.mocks.NewResource(MockResourceArgs{
		TypeToken: in.GetType(),
		Name:      in.GetName(),
		Inputs:    stateIn,
		Provider:  in.GetProvider(),
		ID:        in.GetId(),
		Custom:    false,
		ReadRPC:   in,
	})
	if err != nil {
		return nil, err
	}

	urn := m.newURN(in.GetParent(), in.GetType(), in.GetName())

	m.resources.Store(urn, resource.PropertyMap{
		resource.PropertyKey("urn"):   resource.NewStringProperty(urn),
		resource.PropertyKey("id"):    resource.NewStringProperty(id),
		resource.PropertyKey("state"): resource.NewObjectProperty(state),
	})

	stateOut, err := plugin.MarshalProperties(state, plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: true,
	})
	if err != nil {
		return nil, err
	}

	return &codeinfrarpc.ReadResourceResponse{
		Urn:        urn,
		Properties: stateOut,
	}, nil
}

func (m *mockMonitor) RegisterResource(ctx context.Context, in *codeinfrarpc.RegisterResourceRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.RegisterResourceResponse, error) {
	if in.GetType() == string(resource.RootStackType) && in.GetParent() == "" {
		return &codeinfrarpc.RegisterResourceResponse{
			Urn: m.newURN(in.GetParent(), in.GetType(), in.GetName()),
		}, nil
	}

	inputs, err := plugin.UnmarshalProperties(in.GetObject(), plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: true,
	})
	if err != nil {
		return nil, err
	}

	id, state, err := m.mocks.NewResource(MockResourceArgs{
		TypeToken:   in.GetType(),
		Name:        in.GetName(),
		Inputs:      inputs,
		Provider:    in.GetProvider(),
		ID:          in.GetImportId(),
		Custom:      in.GetCustom(),
		RegisterRPC: in,
	})
	if err != nil {
		return nil, err
	}

	urn := m.newURN(in.GetParent(), in.GetType(), in.GetName())

	m.resources.Store(urn, resource.PropertyMap{
		resource.PropertyKey("urn"):   resource.NewStringProperty(urn),
		resource.PropertyKey("id"):    resource.NewStringProperty(id),
		resource.PropertyKey("state"): resource.NewObjectProperty(state),
	})

	stateOut, err := plugin.MarshalProperties(state, plugin.MarshalOptions{
		KeepSecrets:   true,
		KeepResources: true,
	})
	if err != nil {
		return nil, err
	}

	return &codeinfrarpc.RegisterResourceResponse{
		Urn:    urn,
		Id:     id,
		Object: stateOut,
	}, nil
}

func (m *mockMonitor) RegisterResourceOutputs(ctx context.Context, in *codeinfrarpc.RegisterResourceOutputsRequest,
	opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (m *mockMonitor) RegisterStackTransform(ctx context.Context, in *codeinfrarpc.Callback,
	opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (m *mockMonitor) RegisterStackInvokeTransform(ctx context.Context, in *codeinfrarpc.Callback,
	opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (m *mockMonitor) RegisterPackage(ctx context.Context, in *codeinfrarpc.RegisterPackageRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.RegisterPackageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "RegisterPackage is not implemented")
}

type mockEngine struct {
	logger       *log.Logger
	rootResource string
}

// Log logs a global message in the engine, including errors and warnings.
func (m *mockEngine) Log(ctx context.Context, in *codeinfrarpc.LogRequest,
	opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	if m.logger != nil {
		m.logger.Printf("%s: %s", in.GetSeverity(), in.GetMessage())
	}
	return &emptypb.Empty{}, nil
}

// GetRootResource gets the URN of the root resource, the resource that should be the root of all
// otherwise-unparented resources.
func (m *mockEngine) GetRootResource(ctx context.Context, in *codeinfrarpc.GetRootResourceRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.GetRootResourceResponse, error) {
	return &codeinfrarpc.GetRootResourceResponse{
		Urn: m.rootResource,
	}, nil
}

// SetRootResource sets the URN of the root resource.
func (m *mockEngine) SetRootResource(ctx context.Context, in *codeinfrarpc.SetRootResourceRequest,
	opts ...grpc.CallOption,
) (*codeinfrarpc.SetRootResourceResponse, error) {
	m.rootResource = in.GetUrn()
	return &codeinfrarpc.SetRootResourceResponse{}, nil
}

func (m *mockEngine) StartDebugging(ctx context.Context, in *codeinfrarpc.StartDebuggingRequest,
	opts ...grpc.CallOption,
) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
