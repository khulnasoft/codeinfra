// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Component struct {
	codeinfra.ResourceState
}

func NewComponent(
	ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
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
	Echo string `codeinfra:"echo"`
}

type ComponentGetMessageArgs struct {
	Echo codeinfra.StringInput
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
