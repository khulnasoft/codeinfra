package main

import (
	"example.com/codeinfra-secret/sdk/go/v14/secret"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := secret.NewResource(ctx, "res", &secret.ResourceArgs{
			Private: codeinfra.String("closed"),
			Public:  codeinfra.String("open"),
			PrivateData: &secret.DataArgs{
				Private: codeinfra.String("closed"),
				Public:  codeinfra.String("open"),
			},
			PublicData: &secret.DataArgs{
				Private: codeinfra.String("closed"),
				Public:  codeinfra.String("open"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
