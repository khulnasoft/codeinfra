// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
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

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		if _, err := NewComponent(ctx, "a"); err != nil {
			return err
		}
		return nil
	})
}
