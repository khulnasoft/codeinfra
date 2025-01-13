// Copyright 2016-2023, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/khulnasoft/codeinfra-tls/sdk/v4/go/tls"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
	"github.com/khulnasoft/codeinfra/tests/testdata/codegen/methods-return-plain-resource/go/metaprovider"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cfg := config.New(ctx, "")
		proxy := cfg.Require("proxy")

		configurer, err := metaprovider.NewConfigurer(ctx, "configurer", &metaprovider.ConfigurerArgs{
			TlsProxy: unknownIfDryRun(ctx, proxy),
		})
		if err != nil {
			return err
		}

		prov, err := configurer.TlsProvider(ctx)
		if err != nil {
			return err
		}

		key, err := tls.NewPrivateKey(ctx, "my-private-key", &tls.PrivateKeyArgs{
			Algorithm:  codeinfra.String("ECDSA"),
			EcdsaCurve: codeinfra.String("P384"),
		}, codeinfra.Provider(prov))
		if err != nil {
			return err
		}

		var n int
		n, err = configurer.MeaningOfLife(ctx)
		if err != nil {
			return err
		}

		mix, err := configurer.ObjectMix(ctx)
		if err != nil {
			return err
		}

		ctx.Export("meaningOfLife", codeinfra.Int(n))
		ctx.Export("keyAlgo", key.Algorithm)
		if mix.MeaningOfLife != nil {
			ctx.Export("meaningOfLife2", codeinfra.Int(*mix.MeaningOfLife))
		}

		key2, err := tls.NewPrivateKey(ctx, "my-private-key-2", &tls.PrivateKeyArgs{
			Algorithm:  codeinfra.String("ECDSA"),
			EcdsaCurve: codeinfra.String("P384"),
		}, codeinfra.Provider(mix.Provider))
		if err != nil {
			return err
		}

		ctx.Export("keyAlgo2", key2.Algorithm)
		return nil
	})
}

func unknownIfDryRun(ctx *codeinfra.Context, value string) codeinfra.StringOutput {
	if ctx.DryRun() {
		return codeinfra.UnsafeUnknownOutput(nil).ApplyT(func(_ any) string {
			panic("impossible")
		}).(codeinfra.StringOutput)
	}
	return codeinfra.String(value).ToStringOutput()
}
