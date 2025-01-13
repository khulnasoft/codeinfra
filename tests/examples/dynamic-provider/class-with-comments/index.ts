// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";

class SimpleProvider implements codeinfra.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    // Ensure that the arrow in the following comment does not throw
    //  off how Codeinfra serializes classes/functions.
    // public update: (id: codeinfra.ID, inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
