// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface ComponentArgs {
    message: codeinfra.Input<string>;
    nested: codeinfra.Input<{
        value: codeinfra.Input<string>;
    }>;
}

export class Component extends codeinfra.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true);
    }
}
