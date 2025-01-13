//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

type MyResource struct {
	codeinfra.ResourceState

	Length codeinfra.IntOutput       `codeinfra:"length"`
	Prefix codeinfra.StringPtrOutput `codeinfra:"prefix"`
}

type (
	myResourceArgs struct{}
	MyResourceArgs struct{}
)

func (MyResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*myResourceArgs)(nil)).Elem()
}

func GetResource(ctx *codeinfra.Context, urn codeinfra.URN) (*MyResource, error) {
	var resource MyResource
	err := ctx.RegisterResource("unused:unused:unused", "unused", &MyResourceArgs{}, &resource,
		codeinfra.URN_(string(urn)))
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		c := config.New(ctx, "")
		bar := c.RequireSecret("bar")
		pet, err := random.NewRandomPet(ctx, "cat", &random.RandomPetArgs{
			Length: codeinfra.Int(2),
			Prefix: bar,
		})
		if err != nil {
			return err
		}

		getPetLength := pet.URN().ApplyT(func(urn codeinfra.URN) (codeinfra.IntInput, error) {
			r, err := GetResource(ctx, urn)
			if err != nil {
				return nil, err
			}
			return r.Length, nil
		})
		getPetSecret := pet.URN().ApplyT(func(urn codeinfra.URN) (codeinfra.StringPtrInput, error) {
			r, err := GetResource(ctx, urn)
			if err != nil {
				return nil, err
			}
			return r.Prefix, nil
		})
		ctx.Export("getPetLength", getPetLength)
		ctx.Export("secret", getPetSecret)

		return nil
	})
}
