package main

import (
	"fmt"

	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		var numbers []*random.RandomInteger
		for index := 0; index < 2; index++ {
			key0 := index
			val0 := index
			__res, err := random.NewRandomInteger(ctx, fmt.Sprintf("numbers-%v", key0), &random.RandomIntegerArgs{
				Min:  codeinfra.Int(1),
				Max:  codeinfra.Int(val0),
				Seed: codeinfra.Sprintf("seed%v", val0),
			})
			if err != nil {
				return err
			}
			numbers = append(numbers, __res)
		}
		ctx.Export("first", numbers[0].ID())
		ctx.Export("second", numbers[1].ID())
		return nil
	})
}
