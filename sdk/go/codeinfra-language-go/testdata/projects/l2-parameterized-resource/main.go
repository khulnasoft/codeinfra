package main

import (
	"example.com/codeinfra-subpackage/sdk/go/v2/subpackage"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// The resource name is based on the parameter value
		example, err := subpackage.NewHelloWorld(ctx, "example", nil)
		if err != nil {
			return err
		}
		ctx.Export("parameterValue", example.ParameterValue)
		return nil
	})
}
