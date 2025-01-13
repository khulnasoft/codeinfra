package main

import (
	"example.com/codeinfra-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// This provider covers scenarios where configuration properties are marked as secret in the schema.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			SecretString1: codeinfra.String("SECRET"),
			SecretInt1:    codeinfra.Int(16),
			SecretNum1:    codeinfra.Float64(123456.789),
			SecretBool1:   codeinfra.Bool(true),
			ListSecretString1: codeinfra.StringArray{
				codeinfra.String("SECRET"),
				codeinfra.String("SECRET2"),
			},
			MapSecretString1: codeinfra.StringMap{
				"key1": codeinfra.String("SECRET"),
				"key2": codeinfra.String("SECRET2"),
			},
			ObjSecretString1: &configgrpc.TsecretString1Args{
				SecretX: codeinfra.String("SECRET"),
			},
		})
		if err != nil {
			return err
		}
		_, err = configgrpc.NewConfigFetcher(ctx, "config", nil, codeinfra.Provider(config_grpc_provider))
		if err != nil {
			return err
		}
		return nil
	})
}
