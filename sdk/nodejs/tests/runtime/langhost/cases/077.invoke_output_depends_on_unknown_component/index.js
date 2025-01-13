// Test the dependsOn invoke option with components

const assert = require("assert");
const codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

class MyComponent extends codeinfra.ComponentResource {
    constructor(name) {
        super("test:index:MyComponent", name);
        const dep = new MyResource("dep", { parent: this });
    }
}

const comp = new MyComponent("comp");
const dependsOn = [comp];

codeinfra.runtime.invokeOutput("test:index:echo", {}, { dependsOn });
