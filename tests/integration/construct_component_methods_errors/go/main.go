// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		component, err := NewComponent(ctx, "component")
		if err != nil {
			return err
		}
		_, err = component.GetMessage(ctx, &ComponentGetMessageArgs{
			Echo: codeinfra.String("hello"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
