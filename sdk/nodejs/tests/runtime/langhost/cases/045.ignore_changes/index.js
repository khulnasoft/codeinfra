// This tests the creation of ten propertyless resources.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

new MyResource("testResource", { ignoreChanges: ["ignoredProperty"] });
