// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface ComponentArgs {
    first: codeinfra.Input<string>;
    second: codeinfra.Input<string>;
}

export class Component extends codeinfra.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true);
    }

    getMessage(args: Component.GetMessageArgs): codeinfra.Output<Component.GetMessageResult> {
        return codeinfra.runtime.call("testcomponent:index:Component/getMessage", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export namespace Component {
    export interface GetMessageArgs {
        name: codeinfra.Input<string>;
    }

    export interface GetMessageResult {
        message: string;
    }
}
