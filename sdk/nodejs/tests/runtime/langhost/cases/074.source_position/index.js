let assert = require("assert");
let codeinfra = require("../../../../../");

class MyCustomResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyCustomResource", name, {}, opts);
    }
}

class MyComponentResource extends codeinfra.ComponentResource {
    constructor(name, opts) {
        super("test:index:MyComponentResource", name, {}, opts);
    }
}

const custom = new MyCustomResource("custom");
const component = new MyComponentResource("component");
