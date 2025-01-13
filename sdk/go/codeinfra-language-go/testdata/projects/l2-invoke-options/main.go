package main

import (
	"example.com/codeinfra-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		explicitProvider, err := simpleinvoke.NewProvider(ctx, "explicitProvider", nil)
		if err != nil {
			return err
		}
		data := simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: codeinfra.String("hello"),
		}, codeinfra.Provider(explicitProvider), codeinfra.Parent(explicitProvider), codeinfra.Version("10.0.0"), codeinfra.PluginDownloadURL("https://example.com/github/example"))
		ctx.Export("hello", data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
			return data.Result, nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
