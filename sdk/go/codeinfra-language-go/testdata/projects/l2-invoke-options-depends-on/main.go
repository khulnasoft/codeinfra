package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := simpleinvoke.NewProvider(ctx, "explicitProvider", nil)
		if err != nil {
			return err
		}
		first, err := simpleinvoke.NewStringResource(ctx, "first", &simpleinvoke.StringResourceArgs{
			Text: codeinfra.String("first hello"),
		})
		if err != nil {
			return err
		}
		data := simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: codeinfra.String("hello"),
		}, codeinfra.DependsOn([]codeinfra.Resource{
			codeinfra.Resource(first),
		}))
		_, err = simpleinvoke.NewStringResource(ctx, "second", &simpleinvoke.StringResourceArgs{
			Text: data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
				return data.Result, nil
			}).(codeinfra.StringOutput),
		})
		if err != nil {
			return err
		}
		ctx.Export("hello", data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
			return data.Result, nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
