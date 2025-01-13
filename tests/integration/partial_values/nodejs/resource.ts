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
                outs: inputs,
            };
        };
    }
}

export class Resource extends codeinfra.dynamic.Resource {
    public readonly foo: codeinfra.Output<string>;
    public readonly bar: codeinfra.Output<{ value: string, unknown: string }>;
    public readonly baz: codeinfra.Output<any[]>;

    constructor(name: string, props: ResourceProps, opts?: codeinfra.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    foo: codeinfra.Input<string>;
    bar: codeinfra.Input<{ value: codeinfra.Input<string>, unknown: codeinfra.Input<string> }>;
    baz: codeinfra.Input<codeinfra.Input<any>[]>;
}
