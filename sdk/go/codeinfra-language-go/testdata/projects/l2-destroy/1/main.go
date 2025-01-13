package main

import (
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := simple.NewResource(ctx, "aresource", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
