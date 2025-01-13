// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"fmt"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Bucket struct {
	codeinfra.CustomResourceState
}

func NewBucket(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Bucket, error) {
	resource := &Bucket{}
	err := ctx.RegisterResource(typeToken("Bucket"), name, nil, resource, opts...)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

type BucketComponent struct {
	codeinfra.ResourceState
}

func NewBucketComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*BucketComponent, error) {
	component := &BucketComponent{}
	err := ctx.RegisterComponentResource(typeToken("BucketComponent"), name, component, opts...)
	if err != nil {
		return nil, err
	}

	_, err = NewBucket(ctx, name+"child", codeinfra.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}

type BucketV2 struct {
	codeinfra.CustomResourceState
}

func NewBucketV2(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*BucketV2, error) {
	resource := &BucketV2{}
	aliases := codeinfra.Aliases([]codeinfra.Alias{
		{
			Type: codeinfra.String(typeToken("Bucket")),
		},
	})
	opts = append(opts, aliases)
	err := ctx.RegisterResource(typeToken("BucketV2"), name, nil, resource, opts...)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

type BucketComponentV2 struct {
	codeinfra.ResourceState
}

func NewBucketComponentV2(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*BucketComponentV2, error) {
	component := &BucketComponentV2{}
	aliases := codeinfra.Aliases([]codeinfra.Alias{
		{
			Type: codeinfra.String(typeToken("BucketComponent")),
		},
	})
	opts = append(opts, aliases)
	err := ctx.RegisterComponentResource(typeToken("BucketComponentV2"), name, component, opts...)
	if err != nil {
		return nil, err
	}

	_, err = NewBucketV2(ctx, name+"child", codeinfra.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "wibble"
	version      = "0.0.1"
)

func typeToken(t string) string {
	return fmt.Sprintf("%s:index:%s", providerName, t)
}

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
}

func makeProvider(host *provider.HostClient, name, version string) (codeinfrarpc.ResourceProviderServer, error) {
	return &Provider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

func (p *Provider) Create(ctx context.Context, req *codeinfrarpc.CreateRequest) (*codeinfrarpc.CreateResponse, error) {
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
		var component codeinfra.ComponentResource
		var err error
		switch typ {
		case typeToken("BucketComponent"):
			component, err = NewBucketComponent(ctx, name, options)
		case typeToken("BucketComponentV2"):
			component, err = NewBucketComponentV2(ctx, name, options)
		default:
			err = fmt.Errorf("unknown resource type %s", req.GetType())
		}
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
