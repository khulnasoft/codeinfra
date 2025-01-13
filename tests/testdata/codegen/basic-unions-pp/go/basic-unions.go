package main

import (
	basicunions "github.com/khulnasoft/codeinfra-basic-unions/sdk/v4/go/basic-unions"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// properties field is bound to union case ServerPropertiesForReplica
		_, err := basicunions.NewExampleServer(ctx, "replica", &basicunions.ExampleServerArgs{
			Properties: &basicunions.ServerPropertiesForReplicaArgs{
				CreateMode: codeinfra.String("Replica"),
				Version:    codeinfra.String("0.1.0-dev"),
			},
		})
		if err != nil {
			return err
		}
		// properties field is bound to union case ServerPropertiesForRestore
		_, err = basicunions.NewExampleServer(ctx, "restore", &basicunions.ExampleServerArgs{
			Properties: &basicunions.ServerPropertiesForRestoreArgs{
				CreateMode:         codeinfra.String("PointInTimeRestore"),
				RestorePointInTime: codeinfra.String("example"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
