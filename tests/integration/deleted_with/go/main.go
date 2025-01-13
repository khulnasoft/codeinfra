// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		rand, err := NewRandom(ctx, "random", &RandomArgs{Length: codeinfra.Int(10)})
		if err != nil {
			return err
		}

		_, err = NewFailsOnDelete(ctx, "failsondelete", codeinfra.DeletedWith(rand))
		if err != nil {
			return err
		}

		return nil
	})
}
