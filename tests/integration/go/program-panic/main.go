package main

import "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		panic("great sadness")
	})
}
