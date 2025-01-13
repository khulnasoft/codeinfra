// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export class Component extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts, true);
    }

    createRandom(args: Component.CreateRandomArgs): codeinfra.Output<Component.CreateRandomResult> {
        return codeinfra.runtime.call("testcomponent:index:Component/createRandom", {
            "__self__": this,
            "length": args.length,
        }, this);
    }
}

export namespace Component {
    export interface CreateRandomArgs {
        length: codeinfra.Input<number>;
    }

    export interface CreateRandomResult {
        result: string;
    }
}
