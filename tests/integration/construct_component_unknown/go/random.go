// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

// Exposes the Random resource from the testprovider.
// Requires running `make test_build` and having the built provider on PATH.

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
	if args == nil {
		args = &RandomArgs{}
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomArgs struct {
	Length int `codeinfra:"length"`
}

type RandomArgs struct {
	Length codeinfra.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}
