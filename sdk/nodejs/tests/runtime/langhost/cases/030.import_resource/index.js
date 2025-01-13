// This tests the basic creation of a single propertyless resource.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

new MyResource("testResource1", { import: "testID" });
