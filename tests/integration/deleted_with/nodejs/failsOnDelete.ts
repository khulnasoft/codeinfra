// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export class FailsOnDelete extends codeinfra.CustomResource {
    constructor(name: string, opts?: codeinfra.CustomResourceOptions) {
        super("testprovider:index:FailsOnDelete", name, {}, opts);
    }
}
