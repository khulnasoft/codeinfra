// Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type componentArgs struct {
	Echo interface{} `codeinfra:"echo"`
}

type ComponentArgs struct {
	Echo codeinfra.Input
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState

	Echo    codeinfra.AnyOutput    `codeinfra:"echo"`
	ChildID codeinfra.StringOutput `codeinfra:"childId"`
	Secret  codeinfra.StringOutput `codeinfra:"secret"`
}

func NewComponent(
	ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func NewSecondComponent(
	ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("secondtestcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func NewComponentComponent(
	ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("secondtestcomponent:index:ComponentComponent", name, codeinfra.Map{}, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

type provider struct {
	codeinfra.ProviderResourceState
	expectResourceArg codeinfra.Bool
}

type LocalComponent struct{ codeinfra.ResourceState }

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		componentA, err := NewComponent(ctx, "a", &ComponentArgs{Echo: codeinfra.Int(42)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "b", &ComponentArgs{Echo: componentA.Echo})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "C", &ComponentArgs{Echo: componentA.ChildID})
		if err != nil {
			return err
		}

		provider := &provider{}
		err = ctx.RegisterResource("codeinfra:providers:testcomponent", "provider", codeinfra.Map{
			"expectResourceArg": codeinfra.Bool(true),
		}, provider)
		if err != nil {
			return err
		}
		localComponent := &LocalComponent{}
		err = ctx.RegisterComponentResource("pkg:index:LocalComponent", "localComponent", localComponent, codeinfra.Providers(provider))
		if err != nil {
			return err
		}
		parentProvider := codeinfra.Parent(localComponent)
		_, err = NewComponent(ctx, "checkProvider1",
			&ComponentArgs{Echo: codeinfra.String("checkExpected")}, parentProvider)
		if err != nil {
			return err
		}
		_, err = NewSecondComponent(ctx, "checkProvider2",
			&ComponentArgs{Echo: codeinfra.String("checkExpected")}, parentProvider)
		if err != nil {
			return err
		}
		_, err = NewComponentComponent(ctx, "checkProvider12", parentProvider)
		if err != nil {
			return err
		}
		return nil
	})
}
