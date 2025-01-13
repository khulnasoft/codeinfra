// This tests that resources cannot depend on things which are not resources.

let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, deps) {
        super("test:index:MyResource", name, {}, deps);
    }
}

new MyResource("testResource", { dependsOn: codeinfra.output(Promise.resolve([Promise.resolve(1)])) });
