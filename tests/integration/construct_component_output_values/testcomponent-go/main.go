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
}

type ComponentArgs struct {
	Bar BarPtrInput `codeinfra:"bar"`
	Foo *FooArgs    `codeinfra:"foo"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	if args.Foo == nil {
		return nil, errors.New(`expected args.Foo to be non-nil`)
	}
	if args.Foo.Something == nil {
		return nil, errors.New(`expected args.Foo.Something to be non-nil`)
	}
	something, somethingIsString := args.Foo.Something.(codeinfra.String)
	if !somethingIsString {
		return nil, errors.New(`expected args.Foo.Something to be codeinfra.String`)
	}
	if something != "hello" {
		return nil, fmt.Errorf(`expected args.Foo.Something to equal "hello" but got %q`, something)
	}

	barArgs, isBarArgs := args.Bar.(BarArgs)
	if !isBarArgs {
		return nil, errors.New("expected args.Bar to be BarArgs")
	}
	tags, isStringMap := barArgs.Tags.(codeinfra.StringMap)
	if !isStringMap {
		return nil, errors.New("expected args.Bar.Tags to be codeinfra.StringMap")
	}

	a, aIsString := tags["a"].(codeinfra.String)
	if !aIsString {
		return nil, errors.New(`expected args.Bar.Tags["a"] to be codeinfra.String`)
	}
	if a != "world" {
		return nil, fmt.Errorf(`expected args.Bar.Tags["a"] to equal "world" but got %q`, a)
	}

	b, bIsStringOutput := tags["b"].(codeinfra.StringOutput)
	if !bIsStringOutput {
		return nil, errors.New(`expected args.Bar.Tags["b"] to be codeinfra.StringOutput`)
	}
	b.ApplyT(func(v string) (string, error) {
		if v != "shh" {
			return v, fmt.Errorf(`expected args.Bar.Tags["b"] to equal "shh" but got %q`, v)
		}
		return v, nil
	})

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
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
