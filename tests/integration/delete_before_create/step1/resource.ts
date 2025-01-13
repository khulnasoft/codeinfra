// Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";
import {v4 as uuidv4} from "uuid";

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        };
    }

    public async diff(id: codeinfra.ID, olds: any, news: any): Promise<dynamic.DiffResult> {
        if (olds.state !== news.state) {
            return {
                changes: true,
                replaces: ["state"],
                deleteBeforeReplace: news.noDBR ? false : true,
            };
        }

        if (olds.noReplace !== news.noReplace) {
            return {
                changes: true,
            }
        }

        return {
            changes: false,
        };
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        return {
            id: uuidv4(),
            outs: inputs,
        };
    }
}

export class Resource extends codeinfra.dynamic.Resource {
    public uniqueKey?: codeinfra.Output<number>;
    public state: codeinfra.Output<number>;
    public noReplace?: codeinfra.Output<number>;

    constructor(name: string, props: ResourceProps, opts?: codeinfra.CustomResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    readonly uniqueKey?: codeinfra.Input<number>;
    readonly state: codeinfra.Input<number>;
    readonly noReplace?: codeinfra.Input<number>;
    readonly noDBR?: codeinfra.Input<boolean>;
}
