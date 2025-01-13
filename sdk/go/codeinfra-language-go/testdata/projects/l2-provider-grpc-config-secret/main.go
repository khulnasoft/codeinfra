package main

import (
	"example.com/codeinfra-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// This provider covers scenarios where user passes secret values to the provider.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			String1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				String1: codeinfra.String("SECRET"),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
				return invoke.String1, nil
			}).(codeinfra.StringOutput),
			Int1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Int1: codeinfra.Int(1234567890),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (int, error) {
				return invoke.Int1, nil
			}).(codeinfra.IntOutput),
			Num1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Num1: codeinfra.Float64(123456.789),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (float64, error) {
				return invoke.Num1, nil
			}).(codeinfra.Float64Output),
			Bool1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Bool1: codeinfra.Bool(true),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (bool, error) {
				return invoke.Bool1, nil
			}).(codeinfra.BoolOutput),
			ListString1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				ListString1: codeinfra.StringArray{
					codeinfra.String("SECRET"),
					codeinfra.String("SECRET2"),
				},
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) ([]string, error) {
				return invoke.ListString1, nil
			}).(codeinfra.StringArrayOutput),
			ListString2: codeinfra.StringArray{
				codeinfra.String("VALUE"),
				configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: codeinfra.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(codeinfra.StringOutput),
			},
			MapString2: codeinfra.StringMap{
				"key1": codeinfra.String("value1"),
				"key2": configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: codeinfra.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(codeinfra.StringOutput),
			},
			ObjString2: &configgrpc.Tstring2Args{
				X: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: codeinfra.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(codeinfra.StringOutput),
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
