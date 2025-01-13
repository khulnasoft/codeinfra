// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

let currentID = 0;

class Provider implements codeinfra.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Resource extends codeinfra.dynamic.Resource {
    constructor(name: string, parent?: codeinfra.Resource) {
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
