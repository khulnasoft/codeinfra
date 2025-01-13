// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

// Exposes the Random resource from the testprovider.

package main

import (
	"errors"
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Random struct {
	codeinfra.CustomResourceState

	Length codeinfra.IntOutput    `codeinfra:"length"`
	Result codeinfra.StringOutput `codeinfra:"result"`
}

func NewRandom(ctx *codeinfra.Context,
	name string, args *RandomArgs, opts ...codeinfra.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (r *Random) RandomInvoke(ctx *codeinfra.Context, args map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := ctx.Invoke("testprovider:index:returnArgs", args, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type randomArgs struct {
	Length int    `codeinfra:"length"`
	Prefix string `codeinfra:"prefix"`
}

type RandomArgs struct {
	Length codeinfra.IntInput
	Prefix codeinfra.StringInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState

	Length  codeinfra.IntOutput    `codeinfra:"length"`
	ChildID codeinfra.StringOutput `codeinfra:"childId"`
}

func NewComponent(ctx *codeinfra.Context,
	name string, args *ComponentArgs, opts ...codeinfra.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterRemoteComponentResource("testprovider:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	Length int `codeinfra:"length"`
}

type ComponentArgs struct {
	Length codeinfra.IntInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Provider struct {
	codeinfra.ProviderResourceState
}

func NewProvider(ctx *codeinfra.Context,
	name string, opts ...codeinfra.ResourceOption,
) (*Provider, error) {
	var resource Provider
	err := ctx.RegisterResource("codeinfra:providers:testprovider", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
