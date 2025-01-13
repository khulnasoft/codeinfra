// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
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

// A remote component resource.
type Component struct {
	codeinfra.ResourceState

	Message codeinfra.StringOutput `codeinfra:"message"`
}

// Creates a remote component resource.
func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// A local component resource.
type LocalComponent struct {
	codeinfra.ResourceState

	Message codeinfra.StringOutput
}

// Creates a regular local component resource, which creates a child remote component resource.
func NewLocalComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*LocalComponent, error) {
	var resource LocalComponent
	err := ctx.RegisterComponentResource("my:index:LocalComponent", name, &resource, opts...)
	if err != nil {
		return nil, err
	}

	component, err := NewComponent(ctx, name+"-mycomponent", codeinfra.Parent(&resource))
	if err != nil {
		return nil, err
	}
	resource.Message = component.Message

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

		component, err := NewComponent(ctx, "mycomponent", codeinfra.Provider(provider))
		if err != nil {
			return err
		}

		localComponent, err := NewLocalComponent(ctx, "mylocalcomponent", codeinfra.Providers(provider))
		if err != nil {
			return err
		}

		ctx.Export("message", component.Message)
		ctx.Export("nestedMessage", localComponent.Message)

		return nil
	})
}
