package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfrax"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		secondPasswordLength, resolveSecondPasswordLength := codeinfra.DeferredOutput[int](ctx)
		first, err := NewFirst(ctx, "first", &FirstArgs{
			PasswordLength: codeinfrax.Cast[codeinfra.IntOutput](secondPasswordLength),
		})
		if err != nil {
			return err
		}
		second, err := NewSecond(ctx, "second", &SecondArgs{
			PetName: first.PetName,
		})
		if err != nil {
			return err
		}
		resolveSecondPasswordLength(second.PasswordLength)
		loopingOverMany, resolveLoopingOverMany := codeinfra.DeferredOutput[[]int](ctx)
		another, err := NewFirst(ctx, "another", &FirstArgs{
			PasswordLength: codeinfrax.Cast[codeinfra.IntOutput](loopingOverMany.ApplyT(func(loopingOverMany []int) (codeinfra.Int, error) {
				return codeinfra.Int(len(loopingOverMany)), nil
			}).(codeinfra.IntOutput)),
		})
		if err != nil {
			return err
		}
		var many []*Second
		for index := 0; index < 10; index++ {
			key0 := index
			_ := index
			__res, err := NewSecond(ctx, fmt.Sprintf("many-%v", key0), &SecondArgs{
				PetName: another.PetName,
			})
			if err != nil {
				return err
			}
			many = append(many, __res)
		}
		resolveLoopingOverMany("TODO: For expression")
		return nil
	})
}
