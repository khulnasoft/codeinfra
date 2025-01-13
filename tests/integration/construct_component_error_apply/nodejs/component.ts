// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface ComponentArgs {
    foo: codeinfra.Input<string>;
}

export class Component extends codeinfra.ComponentResource {
    public readonly foo!: codeinfra.Output<string>;

    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["foo"] = args.foo;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

