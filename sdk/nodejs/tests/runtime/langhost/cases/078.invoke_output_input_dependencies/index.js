// Test the dependsOn invoke option with components

const assert = require("assert");
const codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, opts) {
        super("test:index:MyResource", name, {}, opts);
    }
}

const dep = new MyResource("dep");

const argWithResourceDependency = new codeinfra.Output(
    new Set([dep]),
    Promise.resolve(0),
    Promise.resolve(true), // isKnown
    Promise.resolve(false), // isSecret
);

codeinfra.runtime.invokeOutput("test:index:echo", { argWithResourceDependency });
