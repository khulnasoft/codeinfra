// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Resource struct {
	codeinfra.CustomResourceState
}

type resourceArgs struct {
	Echo interface{} `codeinfra:"echo"`
}

type ResourceArgs struct {
	Echo codeinfra.Input
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

func NewResource(ctx *codeinfra.Context, name string, echo codeinfra.Input,
	opts ...codeinfra.ResourceOption,
) (*Resource, error) {
	args := &ResourceArgs{Echo: echo}
	var resource Resource
	err := ctx.RegisterResource(providerName+":index:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type Component struct {
	codeinfra.ResourceState

	Echo    codeinfra.Input        `codeinfra:"echo"`
	ChildID codeinfra.IDOutput     `codeinfra:"childId"`
	Secret  codeinfra.StringOutput `codeinfra:"secret"`
}

type ComponentArgs struct {
	Echo codeinfra.Input `codeinfra:"echo"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	secretKey := "secret"
	fullSecretKey := fmt.Sprintf("%s:%s", ctx.Project(), secretKey)
	if !ctx.IsConfigSecret(fullSecretKey) {
		return nil, fmt.Errorf("expected configuration key to be secret: %s", fullSecretKey)
	}

	conf := config.New(ctx, "")
	secret := conf.RequireSecret(secretKey)

	component := &Component{}
	err := ctx.RegisterComponentResource(providerName+":index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	res, err := NewResource(ctx, fmt.Sprintf("child-%s", name), args.Echo, codeinfra.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Echo = args.Echo
	component.ChildID = res.ID()
	component.Secret = secret

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"secret":  component.Secret,
		"echo":    component.Echo,
		"childId": component.ChildID,
	}); err != nil {
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

type Provider struct {
	codeinfrarpc.UnimplementedResourceProviderServer

	host    *provider.HostClient
	name    string
	version string

	expectResourceArg bool
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
	typ := urn.Type()
	if typ != providerName+":index:Resource" {
		return nil, fmt.Errorf("Unknown resource type '%s'", typ)
	}

	if s, ok := req.GetProperties().Fields["echo"].AsInterface().(string); ok &&
		s == "checkExpected" && !p.expectResourceArg {
		return nil, fmt.Errorf("did not receive configured provider")
	}

	id := currentID
	currentID++

	return &codeinfrarpc.CreateResponse{
		Id: fmt.Sprintf("%v", id),
	}, nil
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
	if _, ok := req.GetArgs().Fields["expectResourceArg"]; ok {
		p.expectResourceArg = true
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
