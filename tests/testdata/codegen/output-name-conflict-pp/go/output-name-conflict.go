package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cfg := config.New(ctx, "")
		cidrBlock := "Test config variable"
		if param := cfg.Get("cidrBlock"); param != "" {
			cidrBlock = param
		}
		ctx.Export("cidrBlock", cidrBlock)
		return nil
	})
}
