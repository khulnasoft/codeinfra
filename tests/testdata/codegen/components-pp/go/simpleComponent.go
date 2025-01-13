package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type SimpleComponentArgs struct {
}

type SimpleComponent struct {
	codeinfra.ResourceState
}

func NewSimpleComponent(
	ctx *codeinfra.Context,
	name string,
	args *SimpleComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*SimpleComponent, error) {
	var componentResource SimpleComponent
	err := ctx.RegisterComponentResource("components:index:SimpleComponent", name, &componentResource, opts...)
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
	_, err = random.NewRandomPassword(ctx, fmt.Sprintf("%s-secondPassword", name), &random.RandomPasswordArgs{
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
