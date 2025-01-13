package main

import (
	"example.com/codeinfra-primitive/sdk/go/v7/primitive"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := primitive.NewResource(ctx, "res", &primitive.ResourceArgs{
			Boolean: codeinfra.Bool(true),
			Float:   codeinfra.Float64(3.14),
			Integer: codeinfra.Int(42),
			String:  codeinfra.String("hello"),
			NumberArray: codeinfra.Float64Array{
				codeinfra.Float64(-1),
				codeinfra.Float64(0),
				codeinfra.Float64(1),
			},
			BooleanMap: codeinfra.BoolMap{
				"t": codeinfra.Bool(true),
				"f": codeinfra.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
