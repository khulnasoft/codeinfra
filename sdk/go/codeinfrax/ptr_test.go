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

package codeinfrax_test

import (
	"context"
	"testing"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/internal"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfrax"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	o := codeinfrax.Ptr(42)
	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, 42, *v.(*int))
}

func TestPtrOf(t *testing.T) {
	t.Parallel()

	vo := codeinfrax.Val("foo")
	o := codeinfrax.PtrOf[string](vo)

	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, "foo", *v.(*string))
}

func TestPtrOutput_Elem(t *testing.T) {
	t.Parallel()

	o := codeinfrax.Ptr(42)
	v, _, _, _, err := internal.AwaitOutput(context.Background(), o.Elem())
	require.NoError(t, err)
	assert.Equal(t, 42, v)
}

func TestPtrOutput_Elem_nil(t *testing.T) {
	t.Parallel()

	o := codeinfrax.Cast[codeinfrax.PtrOutput[string], *string](codeinfrax.Val[*string]((*string)(nil)))

	v, _, _, _, err := internal.AwaitOutput(context.Background(), o.Elem())
	require.NoError(t, err)
	assert.Equal(t, "", v)
}

func TestGPtrOutput(t *testing.T) {
	t.Parallel()

	o := codeinfrax.GPtrOutput[string, codeinfrax.Output[string]](codeinfrax.Ptr("foo"))
	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)

	assert.Equal(t, "foo", *v.(*string))
}

func TestGPtrOutput_Elem(t *testing.T) {
	t.Parallel()

	// Given an Output[*string], we want a pux.GPtrOutput
	// that will produce a codeinfra.StringOutput when Elem is called.
	o := codeinfrax.Cast[
		codeinfrax.GPtrOutput[string, codeinfra.StringOutput],
		*string,
	](codeinfrax.Ptr("foo"))

	so := o.Elem()
	assert.IsType(t, codeinfra.StringOutput{}, so)

	v, _, _, _, err := internal.AwaitOutput(context.Background(), so)
	require.NoError(t, err)
	assert.Equal(t, "foo", v)
}
