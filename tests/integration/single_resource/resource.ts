// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

let currentID = 0;

export class Provider implements codeinfra.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

export class Resource extends codeinfra.dynamic.Resource {
    public readonly state?: any;

    constructor(name: string, props: ResourceProps, opts?: codeinfra.ResourceOptions) {
        super(Provider.instance, name, props, opts);
        this.state = props.state;
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
