// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

let currentID = 0;

export class Provider implements codeinfra.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: undefined,
        };
    }
}

export class Resource extends codeinfra.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__codeinfraType === "codeinfra-nodejs:dynamic:Resource";
    }

    constructor(name: string, props: codeinfra.Inputs, opts?: codeinfra.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
