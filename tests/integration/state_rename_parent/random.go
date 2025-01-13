package main

import (
	"errors"
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Random struct {
	codeinfra.CustomResourceState

	Length codeinfra.IntOutput    `codeinfra:"length"`
	Result codeinfra.StringOutput `codeinfra:"result"`
}

func NewRandom(ctx *codeinfra.Context,
	name string, args *RandomArgs, opts ...codeinfra.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
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
