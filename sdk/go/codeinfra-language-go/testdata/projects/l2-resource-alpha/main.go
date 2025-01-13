package main

import (
	"example.com/codeinfra-alpha/sdk/go/v3/alpha"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := alpha.NewResource(ctx, "res", &alpha.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
