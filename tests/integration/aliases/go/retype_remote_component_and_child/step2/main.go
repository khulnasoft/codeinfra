// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type BucketComponentV2 struct {
	codeinfra.ResourceState
}

func NewBucketComponentV2(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*BucketComponentV2, error) {
	component := &BucketComponentV2{}
	err := ctx.RegisterRemoteComponentResource("wibble:index:BucketComponentV2", name, nil, component, opts...)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewBucketComponentV2(ctx, "main-bucket")
		return err
	})
}
