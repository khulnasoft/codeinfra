// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";
import * as provider from "@codeinfra/codeinfra/provider";

// The component has a child resource that takes a long time to be created.
// We want to ensure the component's slow child resource will actually be created when the
// component is created inside `construct`.

const CREATION_DELAY = 15 * 1000; // 15 second delay in milliseconds

let currentID = 0;

class SlowResource extends dynamic.Resource {
    constructor(name: string, opts?: codeinfra.CustomResourceOptions) {
        const provider = {
            // Return the result after a delay to simulate a resource that takes a long time
            // to be created.
            create: async (inputs: any) => {
                await delay(CREATION_DELAY);
                return {
                    id: (currentID++).toString(),
                    outs: undefined,
                };
            },
        };
        super(provider, name, {}, opts);
    }
}

function delay(timeout: number): Promise<void> {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve();
        }, timeout);
    });
}

class Component extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);
        // Create a child resource that takes a long time in the provider to be created.
        new SlowResource(`child-${name}`, {parent: this});
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        // Create the component with a slow child resource.
        const component = new Component(name, options);
        return Promise.resolve({
            urn: component.urn,
            state: {},
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
