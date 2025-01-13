//go:build !all
// +build !all

package main

import (
	"errors"
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		r, err := NewRandom(ctx, "default", &RandomArgs{
			Length: codeinfra.Int(10),
		}, codeinfra.PluginDownloadURL("get.example.test"))
		if err != nil {
			return err
		}

		provider, err := NewProvider(ctx, "explicit",
			codeinfra.PluginDownloadURL("get.codeinfra.test/providers"))
		e, err := NewRandom(ctx, "explicit", &RandomArgs{
			Length: codeinfra.Int(8),
		}, codeinfra.Provider(provider))
		ctx.Export("default provider", r.Result)
		ctx.Export("explicit provider", e.Result)
		return nil
	})
}

type Random struct {
	codeinfra.CustomResourceState

	Length codeinfra.IntOutput    `codeinfra:"length"`
	Result codeinfra.StringOutput `codeinfra:"result"`
}

func NewProvider(ctx *codeinfra.Context, name string,
	opts ...codeinfra.ResourceOption,
) (codeinfra.ProviderResource, error) {
	provider := Provider{}
	err := ctx.RegisterResource("codeinfra:providers:testprovider",
		"provider", nil, &provider, opts...)
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

type Provider struct {
	codeinfra.ProviderResourceState
}

func NewRandom(ctx *codeinfra.Context,
	name string, args *RandomArgs, opts ...codeinfra.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	if args == nil {
		args = &RandomArgs{}
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomArgs struct {
	Length int `codeinfra:"length"`
}

type RandomArgs struct {
	Length codeinfra.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}
