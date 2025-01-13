// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let codeinfra = require("../../../../../");

class MyResource extends codeinfra.CustomResource {
    constructor(name, args, opts) {
        super(
            "test:index:MyResource",
            name,
            Object.assign(args, {
                outprop: undefined,
            }),
            opts,
        );
    }
}

let resA = new MyResource("resA", {});
let resB = new MyResource("resB", {}, { dependsOn: [resA] });
let resC = new MyResource("resC", {
    propA: resA.outprop,
    propB: resB.outprop,
    propC: "foo",
});
let resD = new MyResource("resD", {
    propA: codeinfra.all([resA.outprop, resB.outprop]).apply(([a, b]) => `${a} ${b}`),
    propB: resC.outprop,
    propC: "bar",
});
let resE = new MyResource(
    "resE",
    {
        propA: resC.outprop,
        propB: codeinfra.all([resA.outprop, resB.outprop]).apply(([a, b]) => `${a} ${b}`),
        propC: "baz",
    },
    { dependsOn: [resD] },
);
