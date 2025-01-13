// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"
	"fmt"

	"github.com/blang/semver"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
)

type Component struct {
	codeinfra.ResourceState
	First  codeinfra.StringOutput `codeinfra:"first"`
	Second codeinfra.StringOutput `codeinfra:"second"`
}

type ComponentArgs struct {
	First  codeinfra.StringInput `codeinfra:"first"`
	Second codeinfra.StringInput `codeinfra:"second"`
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

	component.First = args.First.ToStringOutput()
	component.Second = args.Second.ToStringOutput()

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"first":  args.First,
		"second": args.Second,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

type ComponentGetMessageArgs struct {
	Name codeinfra.StringInput `codeinfra:"name"`
}

type ComponentGetMessageResult struct {
	Message codeinfra.StringOutput `codeinfra:"message"`
}

func (c *Component) GetMessage(args *ComponentGetMessageArgs) (*ComponentGetMessageResult, error) {
	return &ComponentGetMessageResult{
		Message: codeinfra.Sprintf("%s %s, %s!", c.First, c.Second, args.Name),
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

			args := &ComponentArgs{}
			if err := inputs.CopyTo(args); err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}

			component, err := NewComponent(ctx, name, args, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}

			return codeinfraprovider.NewConstructResult(component)
		},
		Call: func(ctx *codeinfra.Context, tok string, args codeinfraprovider.CallArgs) (*codeinfraprovider.CallResult, error) {
			if tok != "testcomponent:index:Component/getMessage" {
				return nil, fmt.Errorf("unknown method %s", tok)
			}

			methodArgs := &ComponentGetMessageArgs{}
			res, err := args.CopyTo(methodArgs)
			if err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}
			component := res.(*Component)

			result, err := component.GetMessage(methodArgs)
			if err != nil {
				return nil, fmt.Errorf("calling method: %w", err)
			}

			return codeinfraprovider.NewCallResult(result)
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
