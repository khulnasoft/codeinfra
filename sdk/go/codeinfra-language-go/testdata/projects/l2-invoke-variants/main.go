package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		res, err := simpleinvoke.NewStringResource(ctx, "res", &simpleinvoke.StringResourceArgs{
			Text: codeinfra.String("hello"),
		})
		if err != nil {
			return err
		}
		ctx.Export("outputInput", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: res.Text,
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(codeinfra.StringOutput))
		ctx.Export("unit", simpleinvoke.UnitOutput(ctx, simpleinvoke.UnitOutputArgs{}, nil).ApplyT(func(invoke simpleinvoke.UnitResult) (string, error) {
			return invoke.Result, nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
