package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		res, err := simple.NewResource(ctx, "res", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("nonSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          codeinfra.String("hello"),
			SecretResponse: codeinfra.Bool(false),
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(codeinfra.StringOutput))
		ctx.Export("firstSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          codeinfra.String("hello"),
			SecretResponse: res.Value,
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(codeinfra.StringOutput))
		ctx.Export("secondSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          codeinfra.ToSecret("goodbye").(codeinfra.StringOutput),
			SecretResponse: codeinfra.Bool(false),
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
