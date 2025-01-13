package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func singleOrNone[T any](elements []T) T {
	if len(elements) != 1 {
		panic(fmt.Errorf("singleOrNone expected input slice to have a single element"))
	}
	return elements[0]
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("result", codeinfra.Float64(singleOrNone([]float64{
			1,
		})))
		return nil
	})
}
