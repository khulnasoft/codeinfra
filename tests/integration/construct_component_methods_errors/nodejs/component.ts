// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export class Component extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts, true);
    }

    getMessage(args: Component.GetMessageArgs): codeinfra.Output<Component.GetMessageResult> {
        return codeinfra.runtime.call("testcomponent:index:Component/getMessage", {
            "__self__": this,
            "echo": args.echo,
        }, this);
    }
}

export namespace Component {
    export interface GetMessageArgs {
        echo: codeinfra.Input<string>;
    }

    export interface GetMessageResult {
        message: string;
    }
}
