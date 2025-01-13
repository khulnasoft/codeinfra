//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		newComponent := func(name string, opts ...codeinfra.ResourceOption) (*Component, error) {
			return NewComponent(ctx, name, &ComponentArgs{Echo: codeinfra.String(name)}, opts...)
		}

		dep1, err := newComponent("Dep1")
		if err != nil {
			return err
		}
		dep2, err := newComponent("Dep2")
		if err != nil {
			return err
		}
		_, err = newComponent("DependsOn", codeinfra.DependsOn([]codeinfra.Resource{dep1, dep2}))
		if err != nil {
			return err
		}

		_, err = newComponent("Protect", codeinfra.Protect(true))
		if err != nil {
			return err
		}

		_, err = newComponent("AdditionalSecretOutputs", codeinfra.AdditionalSecretOutputs([]string{"foo"}))
		if err != nil {
			return err
		}

		_, err = newComponent("CustomTimeouts", codeinfra.Timeouts(&codeinfra.CustomTimeouts{
			Create: ("1m"),
			Update: ("2m"),
			Delete: ("3m"),
		}))
		if err != nil {
			return err
		}

		getDeletedWithMe, err := newComponent("getDeletedWithMe")
		if err != nil {
			return err
		}
		_, err = newComponent("DeletedWith", codeinfra.DeletedWith(getDeletedWithMe))
		if err != nil {
			return err
		}

		_, err = newComponent("RetainOnDelete", codeinfra.RetainOnDelete(true))
		if err != nil {
			return err
		}

		return nil
	})
}

type Component struct {
	codeinfra.ResourceState

	Echo codeinfra.StringOutput `codeinfra:"echo"`
	Foo  codeinfra.StringOutput `codeinfra:"foo"`
	Bar  codeinfra.StringOutput `codeinfra:"bar"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	return &resource, err
}

type ComponentArgs struct {
	Echo codeinfra.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type componentArgs struct {
	Echo string `codeinfra:"echo"`
}
