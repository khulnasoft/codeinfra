// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export interface BarArgs {
    tags?: codeinfra.Input<{[key: string]: codeinfra.Input<string>}>;
}

export interface FooArgs {
    something?: codeinfra.Input<string>;
}

export interface ComponentArgs {
    bar?: codeinfra.Input<BarArgs>;
    foo?: FooArgs;
}

export class Component extends codeinfra.ComponentResource {
    constructor(name: string, args?: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true /*remote*/);
    }
}
