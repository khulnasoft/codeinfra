// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type TestProvider struct {
	codeinfra.ProviderResourceState
}

func NewTestProvider(ctx *codeinfra.Context, name string) (*TestProvider, error) {
	var resource TestProvider
	err := ctx.RegisterResource("codeinfra:providers:testprovider", name, nil, &resource)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	First  string `codeinfra:"first"`
	Second string `codeinfra:"second"`
}

type ComponentArgs struct {
	First  codeinfra.StringInput
	Second codeinfra.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState
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

func (c *Component) GetMessage(ctx *codeinfra.Context, args *ComponentGetMessageArgs) (ComponentGetMessageResultOutput, error) {
	out, err := ctx.Call("testcomponent:index:Component/getMessage", args, ComponentGetMessageResultOutput{}, c)
	if err != nil {
		return ComponentGetMessageResultOutput{}, err
	}
	return out.(ComponentGetMessageResultOutput), nil
}

type componentGetMessageArgs struct {
	Name string `codeinfra:"name"`
}

type ComponentGetMessageArgs struct {
	Name codeinfra.StringInput
}

func (ComponentGetMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentGetMessageArgs)(nil)).Elem()
}

type ComponentGetMessageResult struct {
	Message string `codeinfra:"message"`
}

type ComponentGetMessageResultOutput struct{ *codeinfra.OutputState }

func (ComponentGetMessageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentGetMessageResult)(nil)).Elem()
}

func (o ComponentGetMessageResultOutput) Message() codeinfra.StringOutput {
	return o.ApplyT(func(v ComponentGetMessageResult) string { return v.Message }).(codeinfra.StringOutput)
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((*Component)(nil))
}

func init() {
	codeinfra.RegisterOutputType(ComponentGetMessageResultOutput{})
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		testProvider, err := NewTestProvider(ctx, "testProvider")
		if err != nil {
			return err
		}

		component1, err := NewComponent(ctx, "component1", &ComponentArgs{
			First:  codeinfra.String("Hello"),
			Second: codeinfra.String("World"),
		}, codeinfra.Provider(testProvider))
		if err != nil {
			return err
		}
		result1, err := component1.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: codeinfra.String("Alice"),
		})
		if err != nil {
			return err
		}

		component2, err := NewComponent(ctx, "component2", &ComponentArgs{
			First:  codeinfra.String("Hi"),
			Second: codeinfra.String("There"),
		}, codeinfra.Providers(testProvider))
		if err != nil {
			return err
		}
		result2, err := component2.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: codeinfra.String("Bob"),
		})
		if err != nil {
			return err
		}

		ctx.Export("message1", result1.Message())
		ctx.Export("message2", result2.Message())

		return nil
	})
}
