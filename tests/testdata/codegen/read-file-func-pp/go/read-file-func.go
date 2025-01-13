package main

import (
	"os"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func readFileOrPanic(path string) codeinfra.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return codeinfra.String(string(data))
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		key := readFileOrPanic("key.pub")
		ctx.Export("result", key)
		return nil
	})
}
