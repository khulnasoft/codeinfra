// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

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
