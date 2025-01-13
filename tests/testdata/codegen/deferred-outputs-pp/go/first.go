package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

type FirstArgs struct {
	PasswordLength codeinfra.Float64Input
}

type First struct {
	codeinfra.ResourceState
	PetName  codeinfra.AnyOutput
	Password codeinfra.AnyOutput
}

func NewFirst(
	ctx *codeinfra.Context,
	name string,
	args *FirstArgs,
	opts ...codeinfra.ResourceOption,
) (*First, error) {
	var componentResource First
	err := ctx.RegisterComponentResource("components:index:First", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	randomPet, err := random.NewRandomPet(ctx, fmt.Sprintf("%s-randomPet", name), nil, codeinfra.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	randomPassword, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-randomPassword", name), &random.RandomPasswordArgs{
		Length: args.PasswordLength,
	}, codeinfra.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterResourceOutputs(&componentResource, codeinfra.Map{
		"petName":  randomPet.ID(),
		"password": randomPassword.Result,
	})
	if err != nil {
		return nil, err
	}
	componentResource.PetName = randomPet
	componentResource.Password = randomPassword.Result
	return &componentResource, nil
}
