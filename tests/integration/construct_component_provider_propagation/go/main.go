//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		provider, err := NewRandomProvider(ctx, "explicit")
		if err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_default", nil); err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_provider", nil, codeinfra.Provider(provider)); err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_providers", nil, codeinfra.Providers(provider)); err != nil {
			return err
		}

		providerMap := map[string]codeinfra.ProviderResource{
			"testprovider": provider,
		}
		if _, err := NewComponent(ctx, "uses_providers_map", nil, codeinfra.ProviderMap(providerMap)); err != nil {
			return err
		}

		return nil
	})
}

type RandomProvider struct {
	codeinfra.ProviderResourceState
}

func NewRandomProvider(ctx *codeinfra.Context, name string) (*RandomProvider, error) {
	var provider RandomProvider
	err := ctx.RegisterResource("codeinfra:providers:testprovider", "explicit", nil, &provider)
	return &provider, err
}

type Component struct {
	codeinfra.ResourceState

	Result codeinfra.StringOutput `codeinfra:"result"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	return &resource, err
}

type ComponentArgs struct {
	Result codeinfra.StringInput `codeinfra:"result"`
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooComponentArgs)(nil)).Elem()
}

type fooComponentArgs struct {
	Result codeinfra.StringInput `codeinfra:"result"`
}
