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

func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (c *Component) CreateRandom(ctx *codeinfra.Context, args *ComponentCreateRandomArgs) (ComponentCreateRandomResultOutput, error) {
	out, err := ctx.Call("testcomponent:index:Component/createRandom", args, ComponentCreateRandomResultOutput{}, c)
	if err != nil {
		return ComponentCreateRandomResultOutput{}, err
	}
	return out.(ComponentCreateRandomResultOutput), nil
}

type componentCreateRandomArgs struct {
	Length int `codeinfra:"length"`
}

type ComponentCreateRandomArgs struct {
	Length codeinfra.IntInput
}

func (ComponentCreateRandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentCreateRandomArgs)(nil)).Elem()
}

type ComponentCreateRandomResult struct {
	Result string `codeinfra:"result"`
}

type ComponentCreateRandomResultOutput struct{ *codeinfra.OutputState }

func (ComponentCreateRandomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentCreateRandomResult)(nil)).Elem()
}

func (o ComponentCreateRandomResultOutput) Result() codeinfra.StringOutput {
	return o.ApplyT(func(v ComponentCreateRandomResult) string { return v.Result }).(codeinfra.StringOutput)
}

func init() {
	codeinfra.RegisterOutputType(ComponentCreateRandomResultOutput{})
}
