// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"
	"os"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
)

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

func main() {
	var schema string
	if _, ok := os.LookupEnv("INCLUDE_SCHEMA"); ok {
		schema = `{"hello": "world"}`
	}
	err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Schema:  []byte(schema),
		Construct: func(ctx *codeinfra.Context, typ, name string,
			inputs codeinfraprovider.ConstructInputs, options codeinfra.ResourceOption,
		) (*codeinfraprovider.ConstructResult, error) {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		},
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}
