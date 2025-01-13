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
	"fmt"
	"os"

	"github.com/khulnasoft/codeinfra-tls/sdk/v4/go/tls"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Configurer struct {
	codeinfra.ResourceState

	TlsProviderOutput tls.ProviderOutput `codeinfra:"tlsProvider"`
}

type ConfigurerArgs struct {
	TlsProxy codeinfra.StringInput `codeinfra:"tlsProxy"`
}

func NewConfigurer(
	ctx *codeinfra.Context,
	name string,
	args *ConfigurerArgs,
	opts ...codeinfra.ResourceOption,
) (*Configurer, error) {
	if args == nil {
		return nil, fmt.Errorf("args is required")
	}
	component := &Configurer{}
	err := ctx.RegisterComponentResource(configurerResourceToken, name, component, opts...)
	if err != nil {
		return nil, err
	}

	prov, err := tls.NewProvider(ctx, "tls-p", &tls.ProviderArgs{
		// Due to codeinfra/codeinfra-tls#160 cannot yet set URL here, but can test setting FromEnv.
		Proxy: &tls.ProviderProxyArgs{
			FromEnv: args.TlsProxy.ToStringOutput().ApplyT(func(proxy string) bool {
				if proxy == "FromEnv" {
					return true
				}
				return false
			}).(codeinfra.BoolOutput),
		},
	}, codeinfra.Version("4.10.0"))
	if err != nil {
		return nil, err
	}

	component.TlsProviderOutput = prov.ToProviderOutput()

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"tlsProvider": component.TlsProviderOutput,
	}); err != nil {
		return nil, err
	}
	return component, nil
}

type TlsProviderArgs struct{}

func (c *Configurer) TlsProvider(ctx *codeinfra.Context, args *TlsProviderArgs) (tls.ProviderOutput, error) {
	// The SDKs really do not support receving unknowns plain-resource returning methods, but if desired one can set
	// an UNKNOWNS=true env var to see what happens if the provider was to actually send one, to test the error
	// handling.
	if ctx.DryRun() && os.Getenv("UNKNOWNS") == "true" {
		return codeinfra.UnsafeUnknownOutput(nil).ApplyT(func(x any) *tls.Provider {
			panic("This should not be called")
		}).(tls.ProviderOutput), nil
	}

	return c.TlsProviderOutput, nil
}

type MeaningOfLifeArgs struct{}

type MeaningOfLifeResult struct {
	Result codeinfra.IntOutput `codeinfra:"res"`
}

func (c *Configurer) MeaningOfLife(ctx *codeinfra.Context, args *MeaningOfLifeArgs) (codeinfra.IntOutput, error) {
	return codeinfra.Int(42).ToIntOutputWithContext(ctx.Context()), nil
}

type ObjectMixArgs struct{}

type ObjectMixResult struct {
	Provider      tls.ProviderOutput `codeinfra:"provider"`
	MeaningOfLife codeinfra.IntOutput   `codeinfra:"meaningOfLife"`
}

func (c *Configurer) ObjectMix(ctx *codeinfra.Context, args *ObjectMixArgs) (*ObjectMixResult, error) {
	p, err := c.TlsProvider(ctx, &TlsProviderArgs{})
	if err != nil {
		return nil, err
	}
	m, err := c.MeaningOfLife(ctx, &MeaningOfLifeArgs{})
	if err != nil {
		return nil, err
	}
	return &ObjectMixResult{
		Provider:      p,
		MeaningOfLife: m,
	}, err
}
