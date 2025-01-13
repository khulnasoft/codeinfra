// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

class Component extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts);
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    async construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, options);
        return {
            urn: component.urn,
            state: inputs,
        };
    }

    async call(token: string, inputs: codeinfra.Inputs): Promise<provider.InvokeResult> {
        switch (token) {
            case "testcomponent:index:Component/getMessage":
                return {
                    failures: [{ property: "the failure property", reason: "the failure reason" }],
                };

            default:
                throw new Error(`unknown method ${token}`);
        }
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
