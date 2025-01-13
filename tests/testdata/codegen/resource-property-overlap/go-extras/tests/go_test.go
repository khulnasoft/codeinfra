package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"

	"resource-property-overlap/example"
)

// Tests that XArray{x}.ToXArrayOutput().Index(codeinfra.Int(0)) == x.
func TestArrayOutputIndex(t *testing.T) {
	err := codeinfra.RunErr(func(ctx *codeinfra.Context) error {

		r1, err := example.NewRec(ctx, "rec1", &example.RecArgs{})
		if err != nil {
			return err
		}

		r1o := r1.ToRecOutput()

		r2o := example.RecArray{r1o}.ToRecArrayOutput().
			Index(codeinfra.Int(0))

		wg := &sync.WaitGroup{}
		wg.Add(1)

		codeinfra.All(r1o, r2o).ApplyT(func(xs []interface{}) int {
			assert.Equal(t, xs[0], xs[1])
			wg.Done()
			return 0
		})

		wg.Wait()
		return nil
	}, codeinfra.WithMocks("project", "stack", mocks(0)))
	assert.NoError(t, err)
}

type mocks int

func (mocks) NewResource(args codeinfra.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "_id", args.Inputs, nil
}

func (mocks) Call(args codeinfra.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}
