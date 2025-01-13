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
	"errors"
	"reflect"
	"testing"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/internal"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfrax"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	o := codeinfrax.Flatten[string, codeinfrax.Output[string]](codeinfrax.Val(codeinfrax.Val("a")))
	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.False(t, secret)
	assert.Empty(t, deps)
	assert.Equal(t, "a", v)
}

func TestFlatten_secret(t *testing.T) {
	t.Parallel()

	o := codeinfrax.Flatten[string, codeinfra.StringOutput](
		codeinfrax.Val(
			codeinfra.ToSecret(codeinfra.String("a")).(codeinfra.StringOutput),
		),
	)

	v, known, secret, deps, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)
	assert.True(t, known)
	assert.True(t, secret)
	assert.Empty(t, deps)
	assert.Equal(t, "a", v)
}

func TestFlatten_failedOutput(t *testing.T) {
	t.Parallel()

	in := codeinfrax.Output[codeinfrax.Output[string]]{
		OutputState: internal.NewOutputState(nil, reflect.TypeOf(codeinfrax.Output[string]{})),
	}

	giveErr := errors.New("great sadness")
	internal.RejectOutput(in, giveErr)

	o := codeinfrax.Flatten[string, codeinfrax.Output[string]](in)
	_, _, _, _, err := internal.AwaitOutput(context.Background(), o)
	assert.ErrorIs(t, err, giveErr)
}

func TestAll(t *testing.T) {
	t.Parallel()

	o := codeinfrax.All(
		codeinfrax.Val("a").AsAny(),
		codeinfrax.Val(1).AsAny(),
		codeinfrax.Val(true).AsAny(),
		codeinfrax.Array[string]{codeinfrax.Val("b"), codeinfrax.Val("c")}.AsAny(),
		codeinfrax.Map[int]{"d": codeinfrax.Val(3), "e": codeinfrax.Val(4)}.AsAny(),
	)
	v, _, _, _, err := internal.AwaitOutput(context.Background(), o)
	require.NoError(t, err)

	assert.Equal(t, []any{
		"a",
		1,
		true,
		[]string{"b", "c"},
		map[string]int{"d": 3, "e": 4},
	}, v)
}
