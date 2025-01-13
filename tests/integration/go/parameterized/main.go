package main

import (
	"example.com/codeinfra-pkg/sdk/go/pkg"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := pkg.NewRandom(ctx, "random", &pkg.RandomArgs{
			Length: codeinfra.Int(8),
		})
		if err != nil {
			return err
		}

		hello := "hello"
		_, err = pkg.DoEcho(ctx, &pkg.DoEchoArgs{
			Echo: &hello,
		})
		if err != nil {
			return err
		}

		_ = pkg.DoEchoOutput(ctx, pkg.DoEchoOutputArgs{
			Echo: codeinfra.String("hello"),
		})

		p, err := pkg.NewEcho(ctx, "echo", &pkg.EchoArgs{})
		if err != nil {
			return err
		}

		_, err = p.DoEchoMethod(ctx, &pkg.EchoDoEchoMethodArgs{
			Echo: codeinfra.String("hello"),
		})
		return err
	})
}
