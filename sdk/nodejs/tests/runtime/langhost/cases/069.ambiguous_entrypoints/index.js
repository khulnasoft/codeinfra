// This tests the basic creation of a single propertyless resource.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name) {
        super("test:index:MyResource", name);
    }
}

new MyResource("testResource1");
