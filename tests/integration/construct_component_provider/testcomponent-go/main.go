// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/blang/semver"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
)

type Provider struct {
	codeinfra.ProviderResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

type Component struct {
	codeinfra.ResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Test that we're indeed getting back an instance of `Provider` with its state.
	provider := component.GetProvider("testcomponent::").(*Provider)
	component.Message = provider.Message

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"message": component.Message,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *codeinfra.Context, name, typ, urn string) (codeinfra.ProviderResource, error) {
	if typ != "codeinfra:providers:testcomponent" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, codeinfra.URN_(urn))
	return r, err
}

func main() {
	codeinfra.RegisterResourcePackage(providerName, &pkg{semver.MustParse(version)})

	if err := provider.ComponentMain(providerName, version, nil, func(ctx *codeinfra.Context, typ, name string,
		inputs codeinfraprovider.ConstructInputs, options codeinfra.ResourceOption,
	) (*codeinfraprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}

		component, err := NewComponent(ctx, name, options)
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return codeinfraprovider.NewConstructResult(component)
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
