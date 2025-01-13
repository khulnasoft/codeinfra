// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Named extends codeinfra.CustomResource {
    public readonly name!: codeinfra.Output<string>;
    constructor(name: string, resourceName?: string) {
        super("testprovider:index:Named", name, { name: resourceName });
    }
}

export let autoName = new Named("test1").name;
export let explicitName = new Named("test2", "explicit-name").name;
