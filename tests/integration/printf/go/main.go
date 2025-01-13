// Copyright 2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"
	"os"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		for i := 0; i < 10; i++ {
			fmt.Printf("Line %d\n", i)
			fmt.Fprintf(os.Stderr, "Errln %d\n", i+10)
		}
		fmt.Printf("Line 10")
		fmt.Fprintf(os.Stderr, "Errln 20")
		return nil
	})
}
