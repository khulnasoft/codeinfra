// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Resource struct {
	codeinfra.CustomResourceState
}

func NewResource(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Resource, error) {
	var resource Resource
	if err := ctx.RegisterResource("testcomponent:index:Resource", name, nil, &resource, opts...); err != nil {
		return nil, err
	}
	return &resource, nil
}

type Component struct {
	codeinfra.ResourceState
}

type ComponentArgs struct {
	Children int `codeinfra:"children"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if args.Children > 0 {
		for i := 0; i < args.Children; i++ {
			_, err := NewResource(ctx, fmt.Sprintf("child-%s-%v", name, i+1), codeinfra.Parent(component))
			if err != nil {
				return nil, err
			}
		}
	}

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

var currentID int

func main() {
	err := provider.Main(providerName, func(host *provider.HostClient) (codeinfrarpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

type testcomponentProvider struct {
	codeinfrarpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string
}

func makeProvider(host *provider.HostClient, name, version string) (codeinfrarpc.ResourceProviderServer, error) {
	return &testcomponentProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

func (p *testcomponentProvider) Create(ctx context.Context,
	req *codeinfrarpc.CreateRequest,
) (*codeinfrarpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	typ := urn.Type()
	if typ != "testcomponent:index:Resource" {
		return nil, fmt.Errorf("Unknown resource type '%s'", typ)
	}

	id := currentID
	currentID++

	return &codeinfrarpc.CreateResponse{
		Id: fmt.Sprintf("%v", id),
	}, nil
}

func (p *testcomponentProvider) Construct(ctx context.Context,
	req *codeinfrarpc.ConstructRequest,
) (*codeinfrarpc.ConstructResponse, error) {
	return codeinfraprovider.Construct(ctx, req, p.host.EngineConn(), func(ctx *codeinfra.Context, typ, name string,
		inputs codeinfraprovider.ConstructInputs, options codeinfra.ResourceOption,
	) (*codeinfraprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}

		args := &ComponentArgs{}
		if err := inputs.CopyTo(args); err != nil {
			return nil, fmt.Errorf("setting args: %w", err)
		}

		component, err := NewComponent(ctx, name, args, options)
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return codeinfraprovider.NewConstructResult(component)
	})
}

func (p *testcomponentProvider) CheckConfig(ctx context.Context,
	req *codeinfrarpc.CheckRequest,
) (*codeinfrarpc.CheckResponse, error) {
	return &codeinfrarpc.CheckResponse{Inputs: req.GetNews()}, nil
}

func (p *testcomponentProvider) DiffConfig(ctx context.Context,
	req *codeinfrarpc.DiffRequest,
) (*codeinfrarpc.DiffResponse, error) {
	return &codeinfrarpc.DiffResponse{}, nil
}

func (p *testcomponentProvider) Configure(ctx context.Context,
	req *codeinfrarpc.ConfigureRequest,
) (*codeinfrarpc.ConfigureResponse, error) {
	return &codeinfrarpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: true,
		AcceptResources: true,
	}, nil
}

func (p *testcomponentProvider) Invoke(ctx context.Context,
	req *codeinfrarpc.InvokeRequest,
) (*codeinfrarpc.InvokeResponse, error) {
	return nil, fmt.Errorf("Unknown Invoke token '%s'", req.GetTok())
}

func (p *testcomponentProvider) StreamInvoke(req *codeinfrarpc.InvokeRequest,
	server codeinfrarpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("Unknown StreamInvoke token '%s'", req.GetTok())
}

func (p *testcomponentProvider) Call(ctx context.Context,
	req *codeinfrarpc.CallRequest,
) (*codeinfrarpc.CallResponse, error) {
	return nil, fmt.Errorf("Unknown Call token '%s'", req.GetTok())
}

func (p *testcomponentProvider) Check(ctx context.Context,
	req *codeinfrarpc.CheckRequest,
) (*codeinfrarpc.CheckResponse, error) {
	return &codeinfrarpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (p *testcomponentProvider) Diff(ctx context.Context, req *codeinfrarpc.DiffRequest) (*codeinfrarpc.DiffResponse, error) {
	return &codeinfrarpc.DiffResponse{}, nil
}

func (p *testcomponentProvider) Read(ctx context.Context, req *codeinfrarpc.ReadRequest) (*codeinfrarpc.ReadResponse, error) {
	return &codeinfrarpc.ReadResponse{
		Id:         req.GetId(),
		Properties: req.GetProperties(),
	}, nil
}

func (p *testcomponentProvider) Update(ctx context.Context,
	req *codeinfrarpc.UpdateRequest,
) (*codeinfrarpc.UpdateResponse, error) {
	return &codeinfrarpc.UpdateResponse{
		Properties: req.GetNews(),
	}, nil
}

func (p *testcomponentProvider) Delete(ctx context.Context, req *codeinfrarpc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetPluginInfo(context.Context, *emptypb.Empty) (*codeinfrarpc.PluginInfo, error) {
	return &codeinfrarpc.PluginInfo{
		Version: p.version,
	}, nil
}

func (p *testcomponentProvider) Attach(ctx context.Context, req *codeinfrarpc.PluginAttach) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetSchema(ctx context.Context,
	req *codeinfrarpc.GetSchemaRequest,
) (*codeinfrarpc.GetSchemaResponse, error) {
	return &codeinfrarpc.GetSchemaResponse{}, nil
}

func (p *testcomponentProvider) Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *testcomponentProvider) GetMapping(
	context.Context, *codeinfrarpc.GetMappingRequest,
) (*codeinfrarpc.GetMappingResponse, error) {
	return &codeinfrarpc.GetMappingResponse{}, nil
}
