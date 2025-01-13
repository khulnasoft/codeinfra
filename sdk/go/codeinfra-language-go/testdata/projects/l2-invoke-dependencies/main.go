package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		first, err := simple.NewResource(ctx, "first", &simple.ResourceArgs{
			Value: codeinfra.Bool(false),
		})
		if err != nil {
			return err
		}
		// assert that resource second depends on resource first
		// because it uses .secret from the invoke which depends on first
		_, err = simple.NewResource(ctx, "second", &simple.ResourceArgs{
			Value: simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
				Value:          codeinfra.String("hello"),
				SecretResponse: first.Value,
			}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (bool, error) {
				return invoke.Secret, nil
			}).(codeinfra.BoolOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
