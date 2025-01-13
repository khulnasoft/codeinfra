package main

import (
	storage "github.com/khulnasoft/codeinfra-azure-native/sdk/go/azure/storage"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		someString := "foobar"
		typeVar := "Block"
		staticwebsite, err := storage.NewStorageAccountStaticWebsite(ctx, "staticwebsite", &storage.StorageAccountStaticWebsiteArgs{
			ResourceGroupName: codeinfra.String(someString),
			AccountName:       codeinfra.String(someString),
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "faviconpng", &storage.BlobArgs{
			ResourceGroupName: codeinfra.String(someString),
			AccountName:       codeinfra.String(someString),
			ContainerName:     codeinfra.String(someString),
			Type:              storage.BlobTypeBlock,
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "_404html", &storage.BlobArgs{
			ResourceGroupName: codeinfra.String(someString),
			AccountName:       codeinfra.String(someString),
			ContainerName:     codeinfra.String(someString),
			Type:              staticwebsite.IndexDocument.ApplyT(func(x *string) storage.BlobType { return storage.BlobType(*x) }).(storage.BlobTypeOutput),
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "another", &storage.BlobArgs{
			ResourceGroupName: codeinfra.String(someString),
			AccountName:       codeinfra.String(someString),
			ContainerName:     codeinfra.String(someString),
			Type:              storage.BlobType(typeVar),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
