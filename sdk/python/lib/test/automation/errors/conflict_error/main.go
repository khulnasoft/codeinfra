//go:build !all
// +build !all

package main

import (
	"time"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		time.Sleep(1 * time.Second)
		ctx.Export("exp_static", codeinfra.String("foo"))
		return nil
	})
}
