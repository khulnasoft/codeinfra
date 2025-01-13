// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class FailsOnCreate extends codeinfra.CustomResource {
    public readonly value!: codeinfra.Output<number>;
    constructor(name: string) {
        super("testprovider:index:FailsOnCreate", name, { value: undefined });
    }
}

export let xyz = "DEF";

new FailsOnCreate("test");

export let foo = 1;
