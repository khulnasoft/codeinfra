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

type Component struct {
	codeinfra.ResourceState
}

func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	component := &Component{}
	if err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...); err != nil {
		return nil, err
	}
	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{}); err != nil {
		return nil, err
	}
	return component, nil
}

type ComponentCreateRandomArgs struct {
	Length codeinfra.IntInput `codeinfra:"length"`
}

type ComponentCreateRandomeResult struct {
	Result codeinfra.StringOutput `codeinfra:"result"`
}

func (c *Component) CreateRandom(ctx *codeinfra.Context, args *ComponentCreateRandomArgs) (*ComponentCreateRandomeResult,
	error,
) {
	random, err := NewRandom(ctx, "myrandom", &RandomArgs{Length: args.Length}, codeinfra.Parent(c))
	if err != nil {
		return nil, err
	}

	return &ComponentCreateRandomeResult{
		Result: random.Result,
	}, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *codeinfra.Context, name, typ, urn string) (r codeinfra.Resource, err error) {
	switch typ {
	case "testcomponent:index:Component":
		r = &Component{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, codeinfra.URN_(urn))
	return
}

func main() {
	// Register any resources that can come back as resource references that need to be rehydrated.
	codeinfra.RegisterResourceModule("testcomponent", "index", &module{semver.MustParse(version)})

	if err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Construct: func(ctx *codeinfra.Context, typ, name string, inputs codeinfraprovider.ConstructInputs,
			options codeinfra.ResourceOption,
		) (*codeinfraprovider.ConstructResult, error) {
			if typ != "testcomponent:index:Component" {
				return nil, fmt.Errorf("unknown resource type %s", typ)
			}
			component, err := NewComponent(ctx, name, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}
			return codeinfraprovider.NewConstructResult(component)
		},
		Call: func(ctx *codeinfra.Context, tok string,
			args codeinfraprovider.CallArgs,
		) (*codeinfraprovider.CallResult, error) {
			if tok != "testcomponent:index:Component/createRandom" {
				return nil, fmt.Errorf("unknown method %s", tok)
			}

			methodArgs := &ComponentCreateRandomArgs{}
			res, err := args.CopyTo(methodArgs)
			if err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}
			component := res.(*Component)

			result, err := component.CreateRandom(ctx, methodArgs)
			if err != nil {
				return nil, fmt.Errorf("calling method: %w", err)
			}

			return codeinfraprovider.NewCallResult(result)
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
