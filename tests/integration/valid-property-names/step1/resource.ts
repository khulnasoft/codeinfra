// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
import * as codeinfra from "@codeinfra/codeinfra";

let currentID = 0;

export class Provider implements codeinfra.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    constructor() {}

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: inputs,
        };
    }

    public async delete(id: codeinfra.ID, props: any) {}

    public async diff(id: codeinfra.ID, olds: any, news: any) { return {}; }

    public async update(id: codeinfra.ID, olds: any, news: any) {
        return news;
    }
}

export class Resource extends codeinfra.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: codeinfra.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
