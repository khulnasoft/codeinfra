// Stress-test the engine handling many resources with many aliases.
//go:build !all
// +build !all

package main

import (
	"fmt"

	random "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		conf := config.New(ctx, "")
		mode := conf.Require("mode")
		n := conf.RequireInt("n")

		parent, err := makeResource(ctx, nil, nil, 0, "parent")
		if err != nil {
			return err
		}

		var prev []*random.RandomInteger
		for i := 0; i < n; i++ {
			r, err := makeResource(ctx, parent, prev, i, mode)
			if err != nil {
				return err
			}
			// uncomment for quadratic deps
			// prev = append(prev, r)
			prev = []*random.RandomInteger{r}
		}
		return nil
	})
}

func nameResource(i int, mode string) string {
	return fmt.Sprintf("resource-%s-%d", mode, i)
}

func makeResource(
	ctx *codeinfra.Context,
	parent *random.RandomInteger,
	prev []*random.RandomInteger,
	i int,
	mode string,
) (*random.RandomInteger, error) {
	name := nameResource(i, mode)
	opts := []codeinfra.ResourceOption{}
	if len(prev) != 0 {
		deps := []codeinfra.Resource{}
		for _, p := range prev {
			deps = append(deps, p)
		}
		opts = append(opts, codeinfra.DependsOn(deps))
	} else if parent != nil {
		opts = append(opts, codeinfra.DependsOn([]codeinfra.Resource{parent}))
	}
	if mode == "alias" {
		alias := codeinfra.Alias{
			Name:     codeinfra.String(nameResource(i, "new")),
			NoParent: codeinfra.Bool(true),
		}
		opts = append(opts, codeinfra.Aliases([]codeinfra.Alias{alias}))
		opts = append(opts, codeinfra.Parent(parent))
	}

	if len(prev) != 0 {
		ints := []interface{}{}
		for _, p := range prev {
			ints = append(ints, p.Result)
		}

		var derived codeinfra.IntOutput = codeinfra.All(ints...).ApplyT(func(data []interface{}) int {
			s := 10
			return s
		}).(codeinfra.IntOutput)

		return random.NewRandomInteger(ctx,
			name,
			&random.RandomIntegerArgs{
				Min: derived,
				Max: derived,
			},
			opts...)

	} else {
		return random.NewRandomInteger(ctx,
			name,
			&random.RandomIntegerArgs{
				Min: codeinfra.Int(0),
				Max: codeinfra.Int(100),
			},
			opts...)
	}
}
