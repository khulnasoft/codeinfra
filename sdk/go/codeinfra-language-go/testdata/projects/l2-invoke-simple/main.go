package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("hello", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: codeinfra.String("hello"),
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(codeinfra.StringOutput))
		ctx.Export("goodbye", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: codeinfra.String("goodbye"),
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
