import codeinfra
import codeinfra_asset_archive as asset_archive

ass = asset_archive.AssetResource("ass", value=codeinfra.FileAsset("../test.txt"))
arc = asset_archive.ArchiveResource("arc", value=codeinfra.FileArchive("../archive.tar"))
dir = asset_archive.ArchiveResource("dir", value=codeinfra.FileArchive("../folder"))
assarc = asset_archive.ArchiveResource("assarc", value=codeinfra.AssetArchive({
    "string": codeinfra.StringAsset("file contents"),
    "file": codeinfra.FileAsset("../test.txt"),
    "folder": codeinfra.FileArchive("../folder"),
    "archive": codeinfra.FileArchive("../archive.tar"),
}))
remoteass = asset_archive.AssetResource("remoteass", value=codeinfra.RemoteAsset("https://raw.githubusercontent.com/khulnasoft/codeinfra/master/cmd/codeinfra-test-language/testdata/l2-resource-asset-archive/test.txt"))
