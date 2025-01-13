// This tests simple creation of assets.

let codeinfra = require("../../../../../");

class FileResource extends codeinfra.CustomResource {
    constructor(name, data) {
        super("test:index:FileResource", name, {
            data: data,
        });
    }
}

new FileResource("file1", new codeinfra.asset.FileAsset("./testdata.txt"));
