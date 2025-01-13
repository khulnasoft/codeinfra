package main

import (
	"example.com/codeinfra-primitive-ref/sdk/go/v11/primitiveref"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := primitiveref.NewResource(ctx, "res", &primitiveref.ResourceArgs{
			Data: &primitiveref.DataArgs{
				Boolean: codeinfra.Bool(false),
				Float:   codeinfra.Float64(2.17),
				Integer: codeinfra.Int(-12),
				String:  codeinfra.String("Goodbye"),
				BoolArray: codeinfra.BoolArray{
					codeinfra.Bool(false),
					codeinfra.Bool(true),
				},
				StringMap: codeinfra.StringMap{
					"two":   codeinfra.String("turtle doves"),
					"three": codeinfra.String("french hens"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
