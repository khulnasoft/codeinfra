// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("xyz", codeinfra.String("DEF"))
		return errors.New("program error")
	})
}
