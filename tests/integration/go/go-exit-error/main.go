//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		return fmt.Errorf("hello world")
	})
}
