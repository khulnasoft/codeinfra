//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewComponent(ctx, "foo")
		return err
	})
}

type Component struct {
	codeinfra.ResourceState
}

func NewComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Component, error) {
	var component Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &component, opts...)
	return &component, err
}
