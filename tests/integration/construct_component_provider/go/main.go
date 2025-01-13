// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Provider struct {
	codeinfra.ProviderResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

func NewProvider(ctx *codeinfra.Context,
	name string, args *ProviderArgs, opts ...codeinfra.ResourceOption,
) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	var resource Provider
	err := ctx.RegisterResource("codeinfra:providers:testcomponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	Message string `codeinfra:"message"`
}

type ProviderArgs struct {
	Message codeinfra.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		provider, err := NewProvider(ctx, "myprovider", &ProviderArgs{
			Message: codeinfra.String("hello world"),
		})
		if err != nil {
			return err
		}

		component, err := NewComponent(ctx, "mycomponent", codeinfra.Providers(provider))
		if err != nil {
			return err
		}

		ctx.Export("message", component.Message)

		return nil
	})
}
