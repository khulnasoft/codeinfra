// This tests the basic creation of a single resource with an assets archive property.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name) {
        let archive = new codeinfra.asset.AssetArchive({
            asset: new codeinfra.asset.StringAsset("foo"),
            archive: new codeinfra.asset.AssetArchive({}),
        });
        let archiveP = Promise.resolve(
            new codeinfra.asset.AssetArchive({
                foo: new codeinfra.asset.StringAsset("bar"),
            }),
        );
        let assetP = Promise.resolve(new codeinfra.asset.StringAsset("baz"));
        super("test:index:MyResource", name, {
            archive: archive,
            archiveP: archiveP,
            assetP: assetP,
        });
    }
}

new MyResource("testResource1");
