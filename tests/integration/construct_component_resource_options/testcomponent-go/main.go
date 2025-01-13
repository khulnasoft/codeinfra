// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ComponentArgs struct {
	Echo codeinfra.StringInput `codeinfra:"echo"`
}

type Component struct {
	codeinfra.ResourceState

	Echo codeinfra.StringOutput `codeinfra:"echo"`
	Foo  codeinfra.StringOutput `codeinfra:"foo"`
	Bar  codeinfra.StringOutput `codeinfra:"bar"`
}

func NewComponent(
	ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var comp Component
	if err := ctx.RegisterComponentResource("testcomponent:index:Component", name, &comp, opts...); err != nil {
		return nil, err
	}

	_, err := NewResource(ctx, fmt.Sprintf("%s-child", name), codeinfra.Parent(&comp))
	if err != nil {
		return nil, err
	}

	comp.Echo = args.Echo.ToStringOutput()
	comp.Foo = codeinfra.String("foo").ToStringOutput()
	comp.Bar = codeinfra.String("bar").ToStringOutput()

	return &comp, ctx.RegisterResourceOutputs(&comp, codeinfra.Map{
		"echo": comp.Echo,
		"foo":  comp.Foo,
		"bar":  comp.Bar,
	})
}

type Resource struct {
	codeinfra.CustomResourceState
}

func NewResource(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Resource, error) {
	var res Resource
	err := ctx.RegisterResource("testcomponent:index:Resource", name, nil, &res, opts...)
	return &res, err
}

func main() {
	err := provider.Main("testcomponent", func(host *provider.HostClient) (codeinfrarpc.ResourceProviderServer, error) {
		return NewProvider(host, "testcomponent", "0.0.1"), nil
	})
	if err != nil {
		cmdutil.Exit(err)
	}
}

type Provider struct {
	codeinfrarpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string

	// The next ID to use for a resource.
	currentID atomic.Int64
}

func NewProvider(host *provider.HostClient, name, version string) codeinfrarpc.ResourceProviderServer {
	return &Provider{
		host:    host,
		name:    name,
		version: version,
	}
}

func (p *Provider) Create(ctx context.Context,
	req *codeinfrarpc.CreateRequest,
) (*codeinfrarpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	typ := urn.Type()
	if typ != "testcomponent:index:Resource" {
		return nil, fmt.Errorf("unknown resource type %q", typ)
	}

	id := p.currentID.Add(1)
	return &codeinfrarpc.CreateResponse{
		Id: strconv.FormatInt(id, 10),
	}, nil
}

func (p *Provider) Construct(ctx context.Context,
	req *codeinfrarpc.ConstructRequest,
) (*codeinfrarpc.ConstructResponse, error) {
	return codeinfraprovider.Construct(ctx, req, p.host.EngineConn(), func(
		ctx *codeinfra.Context,
		typ, name string,
		inputs codeinfraprovider.ConstructInputs,
		options codeinfra.ResourceOption,
	) (*codeinfraprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %q", typ)
		}

		var args ComponentArgs
		if err := inputs.CopyTo(&args); err != nil {
			return nil, fmt.Errorf("unmarshal inputs: %w", err)
		}

		comp, err := NewComponent(ctx, name, &args, options)
		if err != nil {
			return nil, err
		}

		return codeinfraprovider.NewConstructResult(comp)
	})
}

func (p *Provider) CheckConfig(ctx context.Context,
	req *codeinfrarpc.CheckRequest,
) (*codeinfrarpc.CheckResponse, error) {
	return &codeinfrarpc.CheckResponse{Inputs: req.GetNews()}, nil
}

func (p *Provider) DiffConfig(ctx context.Context,
	req *codeinfrarpc.DiffRequest,
) (*codeinfrarpc.DiffResponse, error) {
	return &codeinfrarpc.DiffResponse{}, nil
}

func (p *Provider) Configure(ctx context.Context,
	req *codeinfrarpc.ConfigureRequest,
) (*codeinfrarpc.ConfigureResponse, error) {
	return &codeinfrarpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: true,
		AcceptResources: true,
		AcceptOutputs:   true,
	}, nil
}

func (p *Provider) Check(ctx context.Context,
	req *codeinfrarpc.CheckRequest,
) (*codeinfrarpc.CheckResponse, error) {
	return &codeinfrarpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (p *Provider) Diff(ctx context.Context, req *codeinfrarpc.DiffRequest) (*codeinfrarpc.DiffResponse, error) {
	return &codeinfrarpc.DiffResponse{}, nil
}

func (p *Provider) Read(ctx context.Context, req *codeinfrarpc.ReadRequest) (*codeinfrarpc.ReadResponse, error) {
	return &codeinfrarpc.ReadResponse{
		Id:         req.GetId(),
		Properties: req.GetProperties(),
	}, nil
}

func (p *Provider) Update(ctx context.Context,
	req *codeinfrarpc.UpdateRequest,
) (*codeinfrarpc.UpdateResponse, error) {
	return &codeinfrarpc.UpdateResponse{
		Properties: req.GetNews(),
	}, nil
}

func (p *Provider) Delete(ctx context.Context, req *codeinfrarpc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetPluginInfo(context.Context, *emptypb.Empty) (*codeinfrarpc.PluginInfo, error) {
	return &codeinfrarpc.PluginInfo{
		Version: p.version,
	}, nil
}

func (p *Provider) Attach(ctx context.Context, req *codeinfrarpc.PluginAttach) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetSchema(ctx context.Context,
	req *codeinfrarpc.GetSchemaRequest,
) (*codeinfrarpc.GetSchemaResponse, error) {
	return &codeinfrarpc.GetSchemaResponse{}, nil
}

func (p *Provider) Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (p *Provider) GetMapping(context.Context, *codeinfrarpc.GetMappingRequest) (*codeinfrarpc.GetMappingResponse, error) {
	return &codeinfrarpc.GetMappingResponse{}, nil
}

func (p *Provider) Invoke(ctx context.Context,
	req *codeinfrarpc.InvokeRequest,
) (*codeinfrarpc.InvokeResponse, error) {
	return nil, fmt.Errorf("unknown Invoke %q", req.GetTok())
}

func (p *Provider) StreamInvoke(req *codeinfrarpc.InvokeRequest,
	server codeinfrarpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("unknown StreamInvoke %q", req.GetTok())
}

func (p *Provider) Call(ctx context.Context,
	req *codeinfrarpc.CallRequest,
) (*codeinfrarpc.CallResponse, error) {
	return nil, fmt.Errorf("unknown Call %q", req.GetTok())
}
