let assert = require("assert");
let codeinfra = require("../../../../../");

class MyCustomResource extends codeinfra.CustomResource {
    constructor(name, args, opts) {
        super("test:index:MyCustomResource", name, args, opts);
    }
}

class MyComponentResource extends codeinfra.ComponentResource {
    constructor(name, args, opts) {
        super("test:index:MyComponentResource", name, args, opts);
    }
}

//            comp1
//            /   \
//       cust1    cust2

let comp1 = new MyComponentResource("comp1");
let cust1 = new MyCustomResource("cust1", {}, { parent: comp1 });
let cust2 = new MyCustomResource("cust2", {}, { parent: comp1 });
