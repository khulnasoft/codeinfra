package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type AnotherComponentArgs struct {
}

type AnotherComponent struct {
	codeinfra.ResourceState
}

func NewAnotherComponent(
	ctx *codeinfra.Context,
	name string,
	args *AnotherComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*AnotherComponent, error) {
	var componentResource AnotherComponent
	err := ctx.RegisterComponentResource("components:index:AnotherComponent", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	_, err = random.NewRandomPassword(ctx, fmt.Sprintf("%s-firstPassword", name), &random.RandomPasswordArgs{
		Length:  codeinfra.Int(16),
		Special: codeinfra.Bool(true),
	}, codeinfra.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterResourceOutputs(&componentResource, codeinfra.Map{})
	if err != nil {
		return nil, err
	}
	return &componentResource, nil
}
