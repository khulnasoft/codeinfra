// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";
import * as grpc from "@grpc/grpc-js";

class Component extends codeinfra.ComponentResource {
    public readonly foo: codeinfra.Output<string>;

    constructor(name: string, foo: codeinfra.Input<string>, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts);

        this.foo = codeinfra.output(foo);

        this.registerOutputs({
            foo: this.foo,
        })
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    async construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, inputs["foo"], options);

	throw new codeinfra.InputPropertiesError({
	    message: "failing for a reason",
	    errors: [{propertyPath: "foo", reason: "the failure reason"}]});
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
