package main

import (
	usingdashes "example.com/codeinfra-using-dashes/sdk/go/using-dashes"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := usingdashes.NewDash(ctx, "main", &usingdashes.DashArgs{
			Stack: codeinfra.String("dev"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
