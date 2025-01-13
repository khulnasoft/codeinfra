// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface RandomArgs {
    length: codeinfra.Input<number>;
}

export class Random extends codeinfra.CustomResource {
    public readonly length!: codeinfra.Output<number>;
    public readonly result!: codeinfra.Output<string>;
    constructor(name: string, args: RandomArgs, opts?: codeinfra.CustomResourceOptions) {
        super("testprovider:index:Random", name, args, opts);
    }
}
