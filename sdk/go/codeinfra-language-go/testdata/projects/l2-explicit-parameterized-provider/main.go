package main

import (
	"example.com/codeinfra-goodbye/sdk/go/v2/goodbye"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		prov, err := goodbye.NewProvider(ctx, "prov", &goodbye.ProviderArgs{
			Text: codeinfra.String("World"),
		})
		if err != nil {
			return err
		}
		// The resource name is based on the parameter value
		res, err := goodbye.NewGoodbye(ctx, "res", nil, codeinfra.Provider(prov))
		if err != nil {
			return err
		}
		ctx.Export("parameterValue", res.ParameterValue)
		return nil
	})
}
