// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class DynamicProvider extends codeinfra.ProviderResource {
    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super("codeinfra-nodejs", name, {}, opts);
    }
}

class Provider implements codeinfra.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class Resource extends codeinfra.dynamic.Resource {
    constructor(name: string, provider?: codeinfra.ProviderResource) {
        super(Provider.instance, name, {}, { provider: provider});
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Create an explicit instance of the dynamic provider.
let p = new DynamicProvider("p");

// Create a resource using the explicit dynamic provider instance.
let b = new Resource("b", p);
