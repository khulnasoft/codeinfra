// Copyright 2016-2023, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codeinfrax_test // to avoid import cycles

import (
	"reflect"
	"testing"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfrax"
	"github.com/stretchr/testify/require"
)

type testResourceInputs struct {
	PuxString codeinfrax.Input[string]
	PuString  codeinfra.StringInput

	PuxStringPtr codeinfrax.Input[*string]
	PuStringPtr  codeinfra.StringPtrInput

	PuxIntArray codeinfrax.Input[[]int]
	PuIntArray  codeinfra.IntArrayInput

	PuxIntMap codeinfrax.Input[map[string]int]
	PuIntMap  codeinfra.IntMapInput
}

func (*testResourceInputs) ElementType() reflect.Type {
	return reflect.TypeOf((*testResourceArgs)(nil))
}

type testResourceArgs struct {
	PuxString string `codeinfra:"puxString"`
	PuString  string `codeinfra:"puString"`

	PuxStringPtr *string `codeinfra:"puxStringPtr"`
	PuStringPtr  *string `codeinfra:"puStringPtr"`

	PuxIntArray []int `codeinfra:"puxIntArray"`
	PuIntArray  []int `codeinfra:"puIntArray"`

	PuxIntMap map[string]int `codeinfra:"puxIntMap"`
	PuIntMap  map[string]int `codeinfra:"puIntMap"`
}

func TestRegisterResource_inputSerialization(t *testing.T) {
	t.Parallel()

	j := "j"

	tests := []struct {
		desc string
		give codeinfra.Input
		want resource.PropertyMap
	}{
		// --- codeinfrax.Input[string] ---
		{
			desc: "pux.String/pu.String",
			give: &testResourceInputs{PuxString: codeinfra.String("a")},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("a")},
		},
		{
			desc: "pux.String/pu.StringOutput",
			give: &testResourceInputs{PuxString: codeinfra.String("b").ToStringOutput()},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("b")},
		},
		{
			desc: "pux.String/pux.Output[string]",
			give: &testResourceInputs{PuxString: codeinfrax.Val("c")},
			want: resource.PropertyMap{"puxString": resource.NewStringProperty("c")},
		},

		// --- codeinfra.StringInput ---
		{
			desc: "pu.String/pu.String",
			give: &testResourceInputs{PuString: codeinfra.String("d")},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("d")},
		},
		{
			desc: "pu.String/pu.StringOutput",
			give: &testResourceInputs{PuString: codeinfra.String("e").ToStringOutput()},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("e")},
		},
		{
			desc: "pu.String/pux.Output[string] untyped",
			give: &testResourceInputs{PuString: codeinfrax.Val("f").Untyped().(codeinfra.StringOutput)},
			want: resource.PropertyMap{"puString": resource.NewStringProperty("f")},
		},

		// --- codeinfrax.Input[*string] ---
		{
			desc: "pux.StringPtr/pu.String PtrOf",
			give: &testResourceInputs{PuxStringPtr: codeinfrax.PtrOf[string](codeinfra.String("g"))},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("g")},
		},
		{
			desc: "pux.StringPtr/pu.StringPtrOutput",
			give: &testResourceInputs{PuxStringPtr: codeinfra.String("h").ToStringPtrOutput()},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("h")},
		},
		{
			desc: "pux.StringPtr/pux.PtrOutput[string]",
			give: &testResourceInputs{PuxStringPtr: codeinfrax.Ptr("i")},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("i")},
		},
		{
			desc: "pux.StringPtr/pux.Output[*string]",
			give: &testResourceInputs{PuxStringPtr: codeinfrax.Val[*string](&j)},
			want: resource.PropertyMap{"puxStringPtr": resource.NewStringProperty("j")},
		},

		// --- codeinfra.StringPtrInput ---
		{
			desc: "pu.StringPtr/pu.StringPtr",
			give: &testResourceInputs{PuStringPtr: codeinfra.StringPtr("j")},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("j")},
		},
		{
			desc: "pu.StringPtr/pu.String",
			give: &testResourceInputs{PuStringPtr: codeinfra.String("k")},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("k")},
		},
		{
			desc: "pu.StringPtr/pu.StringPtrOutput",
			give: &testResourceInputs{PuStringPtr: codeinfra.String("l").ToStringPtrOutput()},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("l")},
		},
		{
			desc: "pu.StringPtr/pux.PtrOutput[string] untyped",
			give: &testResourceInputs{PuStringPtr: codeinfrax.Ptr("m").Untyped().(codeinfra.StringPtrOutput)},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("m")},
		},
		{
			desc: "pu.StringPtr/pux.GPtrOutput[string] untyped",
			give: &testResourceInputs{
				PuStringPtr: codeinfrax.Cast[codeinfrax.GPtrOutput[string, codeinfra.StringOutput], *string](
					codeinfrax.Ptr("n"),
				).Untyped().(codeinfra.StringPtrOutput),
			},
			want: resource.PropertyMap{"puStringPtr": resource.NewStringProperty("n")},
		},

		// --- codeinfrax.Input[[]int] ---
		{
			desc: "pux.IntArray/pu.IntArray",
			give: &testResourceInputs{
				PuxIntArray: codeinfra.IntArray{
					codeinfra.Int(1),
					codeinfra.Int(2),
					codeinfra.Int(3),
				},
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(1),
						resource.NewNumberProperty(2),
						resource.NewNumberProperty(3),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pu.IntArrayOutput",
			give: &testResourceInputs{
				PuxIntArray: codeinfra.IntArray{
					codeinfra.Int(4),
					codeinfra.Int(5),
					codeinfra.Int(6),
				}.ToIntArrayOutput(),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(4),
						resource.NewNumberProperty(5),
						resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.Output[[]int]",
			give: &testResourceInputs{
				PuxIntArray: codeinfrax.Val([]int{7, 8, 9}),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(7),
						resource.NewNumberProperty(8),
						resource.NewNumberProperty(9),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.ArrayOutput[int]",
			give: &testResourceInputs{
				PuxIntArray: codeinfrax.Cast[codeinfrax.ArrayOutput[int], []int](
					codeinfrax.Val([]int{10, 11, 12}),
				),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(10),
						resource.NewNumberProperty(11),
						resource.NewNumberProperty(12),
					},
				),
			},
		},
		{
			desc: "pux.IntArray/pux.GArrayOutput",
			give: &testResourceInputs{
				PuxIntArray: codeinfrax.Cast[codeinfrax.GArrayOutput[int, codeinfra.IntOutput], []int](
					codeinfrax.Val([]int{13, 14, 15}),
				),
			},
			want: resource.PropertyMap{
				"puxIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(13),
						resource.NewNumberProperty(14),
						resource.NewNumberProperty(15),
					},
				),
			},
		},

		// --- codeinfra.IntArrayInput ---
		{
			desc: "pu.IntArray/pux.Output[[]int] untyped",
			give: &testResourceInputs{
				PuIntArray: codeinfrax.Val([]int{1, 2, 3}).Untyped().(codeinfra.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(1),
						resource.NewNumberProperty(2),
						resource.NewNumberProperty(3),
					},
				),
			},
		},
		{
			desc: "pu.IntArray/pux.ArrayOutput[int] untyped",
			give: &testResourceInputs{
				PuIntArray: codeinfrax.Cast[codeinfrax.ArrayOutput[int], []int](
					codeinfrax.Val([]int{4, 5, 6}),
				).Untyped().(codeinfra.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(4),
						resource.NewNumberProperty(5),
						resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pu.IntArray/pux.GArrayOutput untyped",
			give: &testResourceInputs{
				PuIntArray: codeinfrax.Cast[codeinfrax.GArrayOutput[int, codeinfra.IntOutput], []int](
					codeinfrax.Val([]int{7, 8, 9}),
				).Untyped().(codeinfra.IntArrayOutput),
			},
			want: resource.PropertyMap{
				"puIntArray": resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewNumberProperty(7),
						resource.NewNumberProperty(8),
						resource.NewNumberProperty(9),
					},
				),
			},
		},

		// --- codeinfrax.Input[map[string]int] ---
		{
			desc: "pux.IntMap/pu.IntMap",
			give: &testResourceInputs{
				PuxIntMap: codeinfra.IntMap{"a": codeinfra.Int(1), "b": codeinfra.Int(2)},
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"a": resource.NewNumberProperty(1),
						"b": resource.NewNumberProperty(2),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pu.IntMapOutput",
			give: &testResourceInputs{
				PuxIntMap: codeinfra.IntMap{"c": codeinfra.Int(3), "d": codeinfra.Int(4)}.ToIntMapOutput(),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"c": resource.NewNumberProperty(3),
						"d": resource.NewNumberProperty(4),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.Output[map[string]int]",
			give: &testResourceInputs{
				PuxIntMap: codeinfrax.Val(map[string]int{"e": 5, "f": 6}),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"e": resource.NewNumberProperty(5),
						"f": resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.MapOutput[int]",
			give: &testResourceInputs{
				PuxIntMap: codeinfrax.Cast[codeinfrax.MapOutput[int], map[string]int](
					codeinfrax.Val(map[string]int{"g": 7, "h": 8}),
				),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"g": resource.NewNumberProperty(7),
						"h": resource.NewNumberProperty(8),
					},
				),
			},
		},
		{
			desc: "pux.IntMap/pux.GMapOutput",
			give: &testResourceInputs{
				PuxIntMap: codeinfrax.Cast[codeinfrax.GMapOutput[int, codeinfra.IntOutput], map[string]int](
					codeinfrax.Val(map[string]int{"i": 9, "j": 10}),
				),
			},
			want: resource.PropertyMap{
				"puxIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"i": resource.NewNumberProperty(9),
						"j": resource.NewNumberProperty(10),
					},
				),
			},
		},

		// --- codeinfra.IntMapInput ---
		{
			desc: "pu.IntMap/pu.IntMap",
			give: &testResourceInputs{
				PuIntMap: codeinfra.IntMap{"a": codeinfra.Int(1), "b": codeinfra.Int(2)},
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"a": resource.NewNumberProperty(1),
						"b": resource.NewNumberProperty(2),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.Output[map[string]int] untyped",
			give: &testResourceInputs{
				PuIntMap: codeinfrax.Val(map[string]int{"c": 3, "d": 4}).Untyped().(codeinfra.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"c": resource.NewNumberProperty(3),
						"d": resource.NewNumberProperty(4),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.MapOutput[int] untyped",
			give: &testResourceInputs{
				PuIntMap: codeinfrax.Cast[codeinfrax.MapOutput[int], map[string]int](
					codeinfrax.Val(map[string]int{"e": 5, "f": 6}),
				).Untyped().(codeinfra.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"e": resource.NewNumberProperty(5),
						"f": resource.NewNumberProperty(6),
					},
				),
			},
		},
		{
			desc: "pu.IntMap/pux.GMapOutput untyped",
			give: &testResourceInputs{
				PuIntMap: codeinfrax.Cast[codeinfrax.GMapOutput[int, codeinfra.IntOutput], map[string]int](
					codeinfrax.Val(map[string]int{"g": 7, "h": 8}),
				).Untyped().(codeinfra.IntMapOutput),
			},
			want: resource.PropertyMap{
				"puIntMap": resource.NewObjectProperty(
					resource.PropertyMap{
						"g": resource.NewNumberProperty(7),
						"h": resource.NewNumberProperty(8),
					},
				),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var (
				got resource.PropertyMap
				ok  bool
			)
			err := codeinfra.RunErr(func(ctx *codeinfra.Context) error {
				var res codeinfra.CustomResourceState
				require.NoError(t,
					ctx.RegisterResource("test:index:MyResource", "testResource", tt.give, &res))
				return nil
			}, codeinfra.WithMocks("project", "stack", &mockResourceMonitor{
				CallF: func(args codeinfra.MockCallArgs) (resource.PropertyMap, error) {
					t.Fatalf("unexpected Call(%#v)", args)
					return nil, nil
				},
				NewResourceF: func(args codeinfra.MockResourceArgs) (string, resource.PropertyMap, error) {
					if args.TypeToken == "test:index:MyResource" {
						got = args.Inputs
						ok = true
					} else {
						t.Fatalf("unexpected NewResource(%#v)", args)
					}
					return args.Name + "_id", nil, nil
				},
			}))
			require.NoError(t, err)
			require.True(t, ok)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestRegisterResourceOutputs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		give codeinfra.Input
		// TODO: find a way to intercept RegisterResourceOutput calls.
	}{
		{"pu.String", codeinfra.String("a")},
		{"pu.StringOutput", codeinfra.String("b").ToStringOutput()},
		{"pux.Output[string]", codeinfrax.Val("c")},
		{"pux.Output[string]/untyped", codeinfrax.Val("c").Untyped()},
		{"pux.PtrOf", codeinfrax.PtrOf[string](codeinfra.String("d"))},
		{"pu.StringPtr", codeinfra.StringPtr("e")},
		{"pu.StringPtrOutput", codeinfra.String("f").ToStringPtrOutput()},
		{"pux.Output[*string]", codeinfrax.Ptr("g")},
		{"pux.Output[*string]/untyped", codeinfrax.Ptr("h").Untyped()},
		{"pux.Output[[]int]", codeinfrax.Val([]int{1, 2, 3})},
		{
			"pux.ArrayOutput[int]",
			codeinfrax.Cast[codeinfrax.ArrayOutput[int], []int](
				codeinfrax.Val([]int{4, 5, 6}),
			),
		},
		{
			"pux.GArrayOutput",
			codeinfrax.Cast[codeinfrax.GArrayOutput[int, codeinfra.IntOutput], []int](
				codeinfrax.Val([]int{7, 8, 9}),
			),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			err := codeinfra.RunErr(func(ctx *codeinfra.Context) error {
				ctx.Export("x", tt.give)
				return nil
			}, codeinfra.WithMocks("project", "stack", &mockResourceMonitor{}))
			require.NoError(t, err)
		})
	}
}

type mockResourceMonitor struct {
	CallF        func(codeinfra.MockCallArgs) (resource.PropertyMap, error)
	NewResourceF func(codeinfra.MockResourceArgs) (string, resource.PropertyMap, error)
}

var _ codeinfra.MockResourceMonitor = (*mockResourceMonitor)(nil)

func (m *mockResourceMonitor) Call(args codeinfra.MockCallArgs) (resource.PropertyMap, error) {
	return m.CallF(args)
}

func (m *mockResourceMonitor) NewResource(args codeinfra.MockResourceArgs) (string, resource.PropertyMap, error) {
	return m.NewResourceF(args)
}
