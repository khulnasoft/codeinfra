package main

import (
	"example.com/codeinfra-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Cover interesting schema shapes.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			String1:     codeinfra.String(""),
			String2:     codeinfra.String("x"),
			String3:     codeinfra.String("{}"),
			Int1:        codeinfra.Int(0),
			Int2:        codeinfra.Int(42),
			Num1:        codeinfra.Float64(0),
			Num2:        codeinfra.Float64(42.42),
			Bool1:       codeinfra.Bool(true),
			Bool2:       codeinfra.Bool(false),
			ListString1: codeinfra.StringArray{},
			ListString2: codeinfra.StringArray{
				codeinfra.String(""),
				codeinfra.String("foo"),
			},
			ListInt1: codeinfra.IntArray{
				codeinfra.Int(1),
				codeinfra.Int(2),
			},
			MapString1: codeinfra.StringMap{},
			MapString2: codeinfra.StringMap{
				"key1": codeinfra.String("value1"),
				"key2": codeinfra.String("value2"),
			},
			MapInt1: codeinfra.IntMap{
				"key1": codeinfra.Int(0),
				"key2": codeinfra.Int(42),
			},
			ObjString1: &configgrpc.Tstring1Args{},
			ObjString2: &configgrpc.Tstring2Args{
				X: codeinfra.String("x-value"),
			},
			ObjInt1: &configgrpc.Tint1Args{
				X: codeinfra.Int(42),
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
