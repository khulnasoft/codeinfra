package main

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := random.NewRandomShuffle(ctx, "foo", &random.RandomShuffleArgs{
			Inputs: codeinfra.StringArray{
				codeinfra.String("just one\nnewline"),
				codeinfra.String("foo\nbar\nbaz\nqux\nquux\nqux"),
				codeinfra.String(`{
    "a": 1,
    "b": 2,
    "c": [
      "foo",
      "bar",
      "baz",
      "qux",
      "quux"
    ]
}
`),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
