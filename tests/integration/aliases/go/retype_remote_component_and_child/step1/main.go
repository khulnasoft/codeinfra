// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type BucketComponent struct {
	codeinfra.ResourceState
}

func NewBucketComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*BucketComponent, error) {
	component := &BucketComponent{}
	err := ctx.RegisterRemoteComponentResource("wibble:index:BucketComponent", name, nil, component, opts...)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewBucketComponent(ctx, "main-bucket")
		return err
	})
}
