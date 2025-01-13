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
        const component = new Component(name, inputs["foo"], options);

	return {
	    urn: component.urn,
	    state: inputs,
	};
    }

    async call(token: string, inputs: codeinfra.Inputs): Promise<provider.InvokeResult> {
	switch (token) {
	    case "testcomponent:index:Component/getMessage":
		throw new codeinfra.InputPropertyError({propertyPath: "foo", reason: "the failure reason"});
	    default:
		throw new Error(`unknown method ${token}`);
	};
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
