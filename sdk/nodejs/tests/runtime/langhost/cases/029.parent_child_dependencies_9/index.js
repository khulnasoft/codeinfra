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

//                cust1
//                    \
//                    cust2

let cust1 = new MyCustomResource("cust1", {}, {});
let cust2 = new MyCustomResource("cust2", { parentId: cust1.id }, { parent: cust1 });

let res1 = new MyCustomResource("res1", {}, { dependsOn: cust1 });
