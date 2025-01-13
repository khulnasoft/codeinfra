// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type MyComponent struct {
	codeinfra.ResourceState
}

func NewMyComponent(ctx *codeinfra.Context, name string) (*MyComponent, error) {
	component := &MyComponent{}

	err := ctx.RegisterComponentResource("test:index:MyComponent", name, component)
	if err != nil {
		return nil, err
	}

	return component, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		err := ctx.Log.Debug("A debug message", nil)
		if err != nil {
			return err
		}

		_, err = NewMyComponent(ctx, "mycomponent")
		return err
	})
}
