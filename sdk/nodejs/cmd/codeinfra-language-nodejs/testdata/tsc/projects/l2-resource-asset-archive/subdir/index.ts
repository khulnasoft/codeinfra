import * as codeinfra from "@codeinfra/codeinfra";
import * as asset_archive from "@codeinfra/asset-archive";

const ass = new asset_archive.AssetResource("ass", {value: new codeinfra.asset.FileAsset("../test.txt")});
const arc = new asset_archive.ArchiveResource("arc", {value: new codeinfra.asset.FileArchive("../archive.tar")});
const dir = new asset_archive.ArchiveResource("dir", {value: new codeinfra.asset.FileArchive("../folder")});
const assarc = new asset_archive.ArchiveResource("assarc", {value: new codeinfra.asset.AssetArchive({
    string: new codeinfra.asset.StringAsset("file contents"),
    file: new codeinfra.asset.FileAsset("../test.txt"),
    folder: new codeinfra.asset.FileArchive("../folder"),
    archive: new codeinfra.asset.FileArchive("../archive.tar"),
})});
const remoteass = new asset_archive.AssetResource("remoteass", {value: new codeinfra.asset.RemoteAsset("https://raw.githubusercontent.com/khulnasoft/codeinfra/master/cmd/codeinfra-test-language/testdata/l2-resource-asset-archive/test.txt")});
