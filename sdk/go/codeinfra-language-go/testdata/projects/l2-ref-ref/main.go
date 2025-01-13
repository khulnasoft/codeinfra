package main

import (
	"example.com/codeinfra-ref-ref/sdk/go/v12/refref"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := refref.NewResource(ctx, "res", &refref.ResourceArgs{
			Data: &refref.DataArgs{
				InnerData: &refref.InnerDataArgs{
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
				Boolean:   codeinfra.Bool(true),
				Float:     codeinfra.Float64(4.5),
				Integer:   codeinfra.Int(1024),
				String:    codeinfra.String("Hello"),
				BoolArray: codeinfra.BoolArray{},
				StringMap: codeinfra.StringMap{
					"x": codeinfra.String("100"),
					"y": codeinfra.String("200"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
