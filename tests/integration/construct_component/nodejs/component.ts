// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface ComponentArgs {
    echo: codeinfra.Input<any>;
}

export class Component extends codeinfra.ComponentResource {
    public readonly echo!: codeinfra.Output<any>;
    public readonly childId!: codeinfra.Output<codeinfra.ID>;

    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;
        inputs["secret"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

