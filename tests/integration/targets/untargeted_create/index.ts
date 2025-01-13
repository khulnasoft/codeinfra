// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

let currentID = 0;

class Provider implements codeinfra.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}

class Resource extends codeinfra.dynamic.Resource {
    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
