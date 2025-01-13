// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export class TestProvider extends codeinfra.ProviderResource {
    constructor(name: string) {
        super("testprovider", name);
    }
}
