package main

import (
	"example.com/codeinfra-asset-archive/sdk/go/v5/assetarchive"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := assetarchive.NewAssetResource(ctx, "ass", &assetarchive.AssetResourceArgs{
			Value: codeinfra.NewFileAsset("../test.txt"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "arc", &assetarchive.ArchiveResourceArgs{
			Value: codeinfra.NewFileArchive("../archive.tar"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "dir", &assetarchive.ArchiveResourceArgs{
			Value: codeinfra.NewFileArchive("../folder"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "assarc", &assetarchive.ArchiveResourceArgs{
			Value: codeinfra.NewAssetArchive(map[string]interface{}{
				"string":  codeinfra.NewStringAsset("file contents"),
				"file":    codeinfra.NewFileAsset("../test.txt"),
				"folder":  codeinfra.NewFileArchive("../folder"),
				"archive": codeinfra.NewFileArchive("../archive.tar"),
			}),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewAssetResource(ctx, "remoteass", &assetarchive.AssetResourceArgs{
			Value: codeinfra.NewRemoteAsset("https://raw.githubusercontent.com/khulnasoft/codeinfra/master/cmd/codeinfra-test-language/testdata/l2-resource-asset-archive/test.txt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
