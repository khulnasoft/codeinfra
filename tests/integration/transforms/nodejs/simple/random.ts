// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface RandomArgs {
    length: codeinfra.Input<number>;
    prefix?: codeinfra.Input<string | undefined>;
}

export class Random extends codeinfra.CustomResource {
    public readonly length!: codeinfra.Output<number>;
    public readonly result!: codeinfra.Output<string>;
    constructor(name: string, args: RandomArgs, opts?: codeinfra.CustomResourceOptions) {
        super("testprovider:index:Random", name, args, opts);
    }

    randomInvoke(args) {
	return codeinfra.runtime.invoke("testprovider:index:returnArgs", args);
    }
}


interface ComponentArgs {
    length: codeinfra.Input<number>;
}

export class Component extends codeinfra.ComponentResource {
    public readonly length!: codeinfra.Output<number>;
    public readonly childId!: codeinfra.Output<string>;
    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testprovider:index:Component", name, args, opts, true);
    }
}

export class TestProvider extends codeinfra.ProviderResource {
    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super("testprovider", name, {}, opts);
    }
}
