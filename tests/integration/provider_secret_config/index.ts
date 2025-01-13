import * as codeinfra from "@codeinfra/codeinfra";

// Regression test for [codeinfra/codeinfra#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends codeinfra.ProviderResource {
    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super("codeinfra-nodejs", name,  { secretProperty: codeinfra.secret("it's a secret to everybody") }, opts);
    }
}

const p = new DynamicProvider("p");
