// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cfg := config.New(ctx, "config_missing_go")
		cfg.RequireSecret("notFound")
		return nil
	})
}
