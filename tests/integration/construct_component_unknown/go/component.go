// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type componentArgs struct {
	Message string              `codeinfra:"message"`
	Nested  componentNestedArgs `codeinfra:"nested"`
}

type ComponentArgs struct {
	Message codeinfra.StringInput
	Nested  ComponentNestedInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type componentNestedArgs struct {
	Value string `codeinfra:"Value"`
}

type ComponentNestedArgs struct {
	Value codeinfra.StringInput
}

type ComponentNestedInput interface {
	codeinfra.Input
}

func (ComponentNestedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentNestedArgs)(nil)).Elem()
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
