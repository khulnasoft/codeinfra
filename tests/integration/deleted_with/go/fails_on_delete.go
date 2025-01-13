// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

// Exposes the FailsOnDelete resource from the testprovider.

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type FailsOnDelete struct {
	codeinfra.CustomResourceState
}

func NewFailsOnDelete(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FailsOnDelete, error) {
	var resource FailsOnDelete
	err := ctx.RegisterResource("testprovider:index:FailsOnDelete", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
