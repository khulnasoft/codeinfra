package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

type SecondArgs struct {
	PetName codeinfra.StringInput
}

type Second struct {
	codeinfra.ResourceState
	PasswordLength codeinfra.AnyOutput
}

func NewSecond(
	ctx *codeinfra.Context,
	name string,
	args *SecondArgs,
	opts ...codeinfra.ResourceOption,
) (*Second, error) {
	var componentResource Second
	err := ctx.RegisterComponentResource("components:index:Second", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	_, err = random.NewRandomPet(ctx, fmt.Sprintf("%s-randomPet", name), &random.RandomPetArgs{
		Length: len(args.PetName),
	}, codeinfra.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	password, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-password", name), &random.RandomPasswordArgs{
		Length:  codeinfra.Int(16),
		Special: codeinfra.Bool(true),
		Numeric: codeinfra.Bool(false),
	}, codeinfra.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterResourceOutputs(&componentResource, codeinfra.Map{
		"passwordLength": password.Length,
	})
	if err != nil {
		return nil, err
	}
	componentResource.PasswordLength = password.Length
	return &componentResource, nil
}
