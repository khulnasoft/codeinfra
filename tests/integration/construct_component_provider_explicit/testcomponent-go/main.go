// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"fmt"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

type Component struct {
	codeinfra.ResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

func NewComponent(ctx *codeinfra.Context, name, message string, opts ...codeinfra.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource(providerName+":index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}
	component.Message = codeinfra.String(message).ToStringOutput()

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"message": component.Message,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

func main() {
	err := provider.Main(providerName, func(host *provider.HostClient) (codeinfrarpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

type Provider struct {
	codeinfrarpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string

	message string
}

func makeProvider(host *provider.HostClient, name, version string) (codeinfrarpc.ResourceProviderServer, error) {
	return &Provider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

func (p *Provider) Create(ctx context.Context,
	req *codeinfrarpc.CreateRequest,
) (*codeinfrarpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	return nil, fmt.Errorf("Unknown resource type '%s'", urn.Type())
}

func (p *Provider) Construct(ctx context.Context,
	req *codeinfrarpc.ConstructRequest,
) (*codeinfrarpc.ConstructResponse, error) {
	return codeinfraprovider.Construct(ctx, req, p.host.EngineConn(), func(ctx *codeinfra.Context, typ, name string,
		inputs codeinfraprovider.ConstructInputs, options codeinfra.ResourceOption,
	) (*codeinfraprovider.ConstructResult, error) {
		if typ != providerName+":index:Component" {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}

		component, err := NewComponent(ctx, name, p.message, options)
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return codeinfraprovider.NewConstructResult(component)
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
	if val, ok := req.GetArgs().Fields["message"]; ok {
		p.message = val.GetStringValue()
	}
	return &codeinfrarpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: true,
		AcceptResources: true,
	}, nil
}

func (p *Provider) Invoke(ctx context.Context,
	req *codeinfrarpc.InvokeRequest,
) (*codeinfrarpc.InvokeResponse, error) {
	return nil, fmt.Errorf("Unknown Invoke token '%s'", req.GetTok())
}

func (p *Provider) StreamInvoke(req *codeinfrarpc.InvokeRequest,
	server codeinfrarpc.ResourceProvider_StreamInvokeServer,
) error {
	return fmt.Errorf("Unknown StreamInvoke token '%s'", req.GetTok())
}

func (p *Provider) Call(ctx context.Context,
	req *codeinfrarpc.CallRequest,
) (*codeinfrarpc.CallResponse, error) {
	return nil, fmt.Errorf("Unknown Call token '%s'", req.GetTok())
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
