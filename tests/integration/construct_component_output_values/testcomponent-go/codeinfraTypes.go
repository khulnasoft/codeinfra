// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"context"
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Bar struct {
	Tags map[string]string `codeinfra:"tags"`
}

// BarInput is an input type that accepts BarArgs and BarOutput values.
// You can construct a concrete instance of `BarInput` via:
//
//	BarArgs{...}
type BarInput interface {
	codeinfra.Input

	ToBarOutput() BarOutput
	ToBarOutputWithContext(context.Context) BarOutput
}

type BarArgs struct {
	Tags codeinfra.StringMapInput `codeinfra:"tags"`
}

func (BarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Bar)(nil)).Elem()
}

func (i BarArgs) ToBarOutput() BarOutput {
	return i.ToBarOutputWithContext(context.Background())
}

func (i BarArgs) ToBarOutputWithContext(ctx context.Context) BarOutput {
	return codeinfra.ToOutputWithContext(ctx, i).(BarOutput)
}

func (i BarArgs) ToBarPtrOutput() BarPtrOutput {
	return i.ToBarPtrOutputWithContext(context.Background())
}

func (i BarArgs) ToBarPtrOutputWithContext(ctx context.Context) BarPtrOutput {
	return codeinfra.ToOutputWithContext(ctx, i).(BarOutput).ToBarPtrOutputWithContext(ctx)
}

// BarPtrInput is an input type that accepts BarArgs, BarPtr and BarPtrOutput values.
// You can construct a concrete instance of `BarPtrInput` via:
//
//	        BarArgs{...}
//
//	or:
//
//	        nil
type BarPtrInput interface {
	codeinfra.Input

	ToBarPtrOutput() BarPtrOutput
	ToBarPtrOutputWithContext(context.Context) BarPtrOutput
}

type BarOutput struct{ *codeinfra.OutputState }

func (BarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bar)(nil)).Elem()
}

func (o BarOutput) ToBarOutput() BarOutput {
	return o
}

func (o BarOutput) ToBarOutputWithContext(ctx context.Context) BarOutput {
	return o
}

func (o BarOutput) ToBarPtrOutput() BarPtrOutput {
	return o.ToBarPtrOutputWithContext(context.Background())
}

func (o BarOutput) ToBarPtrOutputWithContext(ctx context.Context) BarPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Bar) *Bar {
		return &v
	}).(BarPtrOutput)
}

func (o BarOutput) Tags() codeinfra.StringMapOutput {
	return o.ApplyT(func(v Bar) map[string]string { return v.Tags }).(codeinfra.StringMapOutput)
}

type BarPtrOutput struct{ *codeinfra.OutputState }

func (BarPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bar)(nil)).Elem()
}

func (o BarPtrOutput) ToBarPtrOutput() BarPtrOutput {
	return o
}

func (o BarPtrOutput) ToBarPtrOutputWithContext(ctx context.Context) BarPtrOutput {
	return o
}

func (o BarPtrOutput) Elem() BarOutput {
	return o.ApplyT(func(v *Bar) Bar {
		if v != nil {
			return *v
		}
		var ret Bar
		return ret
	}).(BarOutput)
}

func (o BarPtrOutput) Tags() codeinfra.StringMapOutput {
	return o.ApplyT(func(v *Bar) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(codeinfra.StringMapOutput)
}

type Foo struct {
	Something *string `codeinfra:"something"`
}

// FooInput is an input type that accepts FooArgs and FooOutput values.
// You can construct a concrete instance of `FooInput` via:
//
//	FooArgs{...}
type FooInput interface {
	codeinfra.Input

	ToFooOutput() FooOutput
	ToFooOutputWithContext(context.Context) FooOutput
}

type FooArgs struct {
	Something codeinfra.StringPtrInput `codeinfra:"something"`
}

func (FooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Foo)(nil)).Elem()
}

func (i FooArgs) ToFooOutput() FooOutput {
	return i.ToFooOutputWithContext(context.Background())
}

func (i FooArgs) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return codeinfra.ToOutputWithContext(ctx, i).(FooOutput)
}

func (i FooArgs) ToFooPtrOutput() FooPtrOutput {
	return i.ToFooPtrOutputWithContext(context.Background())
}

func (i FooArgs) ToFooPtrOutputWithContext(ctx context.Context) FooPtrOutput {
	return codeinfra.ToOutputWithContext(ctx, i).(FooOutput).ToFooPtrOutputWithContext(ctx)
}

// FooPtrInput is an input type that accepts FooArgs, FooPtr and FooPtrOutput values.
// You can construct a concrete instance of `FooPtrInput` via:
//
//	        FooArgs{...}
//
//	or:
//
//	        nil
type FooPtrInput interface {
	codeinfra.Input

	ToFooPtrOutput() FooPtrOutput
	ToFooPtrOutputWithContext(context.Context) FooPtrOutput
}

type FooOutput struct{ *codeinfra.OutputState }

func (FooOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Foo)(nil)).Elem()
}

func (o FooOutput) ToFooOutput() FooOutput {
	return o
}

func (o FooOutput) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return o
}

func (o FooOutput) ToFooPtrOutput() FooPtrOutput {
	return o.ToFooPtrOutputWithContext(context.Background())
}

func (o FooOutput) ToFooPtrOutputWithContext(ctx context.Context) FooPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Foo) *Foo {
		return &v
	}).(FooPtrOutput)
}

func (o FooOutput) Something() codeinfra.StringPtrOutput {
	return o.ApplyT(func(v Foo) *string { return v.Something }).(codeinfra.StringPtrOutput)
}

type FooPtrOutput struct{ *codeinfra.OutputState }

func (FooPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Foo)(nil)).Elem()
}

func (o FooPtrOutput) ToFooPtrOutput() FooPtrOutput {
	return o
}

func (o FooPtrOutput) ToFooPtrOutputWithContext(ctx context.Context) FooPtrOutput {
	return o
}

func (o FooPtrOutput) Elem() FooOutput {
	return o.ApplyT(func(v *Foo) Foo {
		if v != nil {
			return *v
		}
		var ret Foo
		return ret
	}).(FooOutput)
}

func (o FooPtrOutput) Something() codeinfra.StringPtrOutput {
	return o.ApplyT(func(v *Foo) *string {
		if v == nil {
			return nil
		}
		return v.Something
	}).(codeinfra.StringPtrOutput)
}

func init() {
	codeinfra.RegisterInputType(reflect.TypeOf((*BarInput)(nil)).Elem(), BarArgs{})
	codeinfra.RegisterInputType(reflect.TypeOf((*BarPtrInput)(nil)).Elem(), BarArgs{})
	codeinfra.RegisterInputType(reflect.TypeOf((*FooInput)(nil)).Elem(), FooArgs{})
	codeinfra.RegisterInputType(reflect.TypeOf((*FooPtrInput)(nil)).Elem(), FooArgs{})
	codeinfra.RegisterOutputType(BarOutput{})
	codeinfra.RegisterOutputType(BarPtrOutput{})
	codeinfra.RegisterOutputType(FooOutput{})
	codeinfra.RegisterOutputType(FooPtrOutput{})
}
