// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let codeinfra = require("../../../../../");

(async () => {
    class Provider extends codeinfra.ProviderResource {
        constructor(name, opts) {
            super("test", name, {}, opts);
        }
    }

    const provider = new Provider("p");
    await codeinfra.ProviderResource.register(provider);

    let args = {
        a: "hello",
        b: true,
        c: [0.99, 42, { z: "x" }],
        id: "some-id",
        urn: "some-urn",
    };

    let result1 = await codeinfra.runtime.invoke("test:index:echo", args, { provider });
    for (const key in args) {
        assert.deepStrictEqual(result1[key], args[key]);
    }

    let result2 = codeinfra.runtime.invoke("test:index:echo", args, { provider, async: false });
    result2.then((v) => {
        assert.deepStrictEqual(v, args);
    });
})();
