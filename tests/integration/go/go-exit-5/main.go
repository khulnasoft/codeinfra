//go:build !all
// +build !all

package main

import (
	"os"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		os.Exit(5)
		return nil
	})
}
