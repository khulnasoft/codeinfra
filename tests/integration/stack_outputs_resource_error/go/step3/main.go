// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type FailsOnCreate struct {
	codeinfra.CustomResourceState

	Value codeinfra.Float64Output `codeinfra:"value"`
}

func NewFailsOnCreate(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FailsOnCreate, error) {
	var resource FailsOnCreate
	err := ctx.RegisterResource("testprovider:index:FailsOnCreate", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("xyz", codeinfra.String("DEF"))
		NewFailsOnCreate(ctx, "test")
		ctx.Export("foo", codeinfra.Float64(1))
		return nil
	})
}
