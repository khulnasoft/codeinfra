package main

import (
	"example.com/codeinfra-config/sdk/go/v9/config"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		prov, err := config.NewProvider(ctx, "prov", &config.ProviderArgs{
			Name:              codeinfra.String("my config"),
			PluginDownloadURL: codeinfra.String("not the same as the codeinfra resource option"),
		})
		if err != nil {
			return err
		}
		// Note this isn't _using_ the explicit provider, it's just grabbing a value from it.
		_, err = config.NewResource(ctx, "res", &config.ResourceArgs{
			Text: prov.Version,
		})
		if err != nil {
			return err
		}
		ctx.Export("pluginDownloadURL", prov.PluginDownloadURL)
		return nil
	})
}
