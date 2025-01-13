// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

const version = "0.0.1";

class Provider extends codeinfra.ProviderResource {
    public readonly message!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super("testcomponent", name, { "message": undefined }, opts);
    }
}

class Component extends codeinfra.ComponentResource {
    public message!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);
    }

    protected async initialize(args: codeinfra.Inputs) {
        const provider = this.getProvider("testcomponent::");
        if (!(provider instanceof Provider)) {
            throw new Error("provider is not an instance of Provider");
        }
        this.message = provider.message;
        this.registerOutputs({
            message: provider.message,
        });
        return undefined;
    }
}

class ProviderServer implements provider.Provider {
    public readonly version = version;

    constructor() {
        codeinfra.runtime.registerResourcePackage("testcomponent", {
            version,
            constructProvider: (name: string, type: string, urn: string): codeinfra.ProviderResource => {
                if (type !== "codeinfra:providers:testcomponent") {
                    throw new Error(`unknown provider type ${type}`);
                }
                return new Provider(name, { urn });
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
            state: {
                message: component.message,
            },
        };
    }
}

export function main(args: string[]) {
    return provider.main(new ProviderServer(), args);
}

main(process.argv.slice(2));
