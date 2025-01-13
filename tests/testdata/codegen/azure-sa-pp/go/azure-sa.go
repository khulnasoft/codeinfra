package main

import (
	"github.com/khulnasoft/codeinfra-azure/sdk/v4/go/azure/core"
	"github.com/khulnasoft/codeinfra-azure/sdk/v4/go/azure/storage"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cfg := config.New(ctx, "")
		// The name of the storage account
		storageAccountNameParam := cfg.Require("storageAccountNameParam")
		// The name of the resource group
		resourceGroupNameParam := cfg.Require("resourceGroupNameParam")
		resourceGroupVar, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
			Name: resourceGroupNameParam,
		}, nil)
		if err != nil {
			return err
		}
		locationParam := resourceGroupVar.Location
		if param := cfg.Get("locationParam"); param != "" {
			locationParam = param
		}
		storageAccountTierParam := "Standard"
		if param := cfg.Get("storageAccountTierParam"); param != "" {
			storageAccountTierParam = param
		}
		storageAccountTypeReplicationParam := "LRS"
		if param := cfg.Get("storageAccountTypeReplicationParam"); param != "" {
			storageAccountTypeReplicationParam = param
		}
		storageAccountResource, err := storage.NewAccount(ctx, "storageAccountResource", &storage.AccountArgs{
			Name:                   codeinfra.String(storageAccountNameParam),
			AccountKind:            codeinfra.String("StorageV2"),
			Location:               codeinfra.String(locationParam),
			ResourceGroupName:      codeinfra.String(resourceGroupNameParam),
			AccountTier:            codeinfra.String(storageAccountTierParam),
			AccountReplicationType: codeinfra.String(storageAccountTypeReplicationParam),
		})
		if err != nil {
			return err
		}
		ctx.Export("storageAccountNameOut", storageAccountResource.Name)
		return nil
	})
}
