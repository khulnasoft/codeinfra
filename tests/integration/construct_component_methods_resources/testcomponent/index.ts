// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";
import { Random } from "./random"

class Component extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts);
    }

    createRandom(length: codeinfra.Input<number>): codeinfra.Output<string> {
        const r = new Random("myrandom", { length }, { parent: this });
        return r.result;
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    constructor() {
        // Register any resources that can come back as resource references that need to be rehydrated.
        codeinfra.runtime.registerResourceModule("testcomponent", "index", {
            version: this.version,
            construct: (name, type, urn) => {
                switch (type) {
                    case "testcomponent:index:Component":
                        return new Component(name, { urn });
                    default:
                        throw new Error(`unknown resource type ${type}`);
                }
            },
        });
    }

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
        if (token != "testcomponent:index:Component/createRandom") {
            throw new Error(`unknown method ${token}`);
        }

        const self: Component = inputs.__self__;
        const result = self.createRandom(inputs.length);
        return { outputs: { result } };
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
