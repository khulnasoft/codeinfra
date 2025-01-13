package main

import (
	"github.com/khulnasoft/codeinfra-azure-native/sdk/go/azure/storage"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := storage.NewStorageAccount(ctx, "storageAccounts", &storage.StorageAccountArgs{
			AccountName:       codeinfra.String("sto4445"),
			Kind:              codeinfra.String(storage.KindBlockBlobStorage),
			Location:          codeinfra.String("eastus"),
			ResourceGroupName: codeinfra.String("res9101"),
			Sku: &storage.SkuArgs{
				Name: codeinfra.String(storage.SkuName_Premium_LRS),
			},
			NetworkRuleSet: &storage.NetworkRuleSetArgs{
				DefaultAction: storage.DefaultActionAllow,
				IpRules:       storage.IPRuleArray{},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
