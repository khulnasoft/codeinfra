//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		c := config.New(ctx, "")
		ctx.Export("exp_static", codeinfra.String("foo"))
		ctx.Export("exp_cfg", codeinfra.String(c.Get("bar")))
		ctx.Export("exp_secret", c.GetSecret("buzz"))
		return nil
	})
}
