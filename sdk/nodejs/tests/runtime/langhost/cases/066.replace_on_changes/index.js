// This tests the replaceOnChanges ResourceOption.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

new MyResource("testResource", { replaceOnChanges: ["foo"] });
