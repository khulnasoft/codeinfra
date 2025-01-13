//go:build !all
// +build !all

package main

import (
	"strings"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", codeinfra.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
