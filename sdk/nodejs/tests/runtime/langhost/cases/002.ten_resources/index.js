// This tests the creation of ten propertyless resources.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name) {
        super("test:index:MyResource", name);
    }
}

for (let i = 0; i < 10; i++) {
    new MyResource("testResource" + i);
}
