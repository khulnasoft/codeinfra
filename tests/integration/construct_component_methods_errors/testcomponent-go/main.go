// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

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
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
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

func main() {
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
		Call: func(ctx *codeinfra.Context, tok string, args codeinfraprovider.CallArgs) (*codeinfraprovider.CallResult, error) {
			if tok != "testcomponent:index:Component/getMessage" {
				return nil, fmt.Errorf("unknown method %s", tok)
			}

			return &codeinfraprovider.CallResult{
				Failures: []codeinfraprovider.CallFailure{
					{
						Property: "the failure property",
						Reason:   "the failure reason",
					},
				},
			}, nil
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
