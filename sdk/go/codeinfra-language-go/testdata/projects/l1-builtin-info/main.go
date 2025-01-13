package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("stackOutput", codeinfra.String(ctx.Stack()))
		ctx.Export("projectOutput", codeinfra.String(ctx.Project()))
		ctx.Export("organizationOutput", codeinfra.String(ctx.Organization()))
		return nil
	})
}
