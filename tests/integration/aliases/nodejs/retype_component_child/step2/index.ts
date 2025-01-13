// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class FooResource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        const aliasOpts = { aliases: [{ type: "my:module:FooResource" }] };
        opts = codeinfra.mergeOptions(opts, aliasOpts);
        super("my:module:FooResourceNew", name, {}, opts);
    }
}

class ComponentResource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:ComponentResource", name, {}, opts);
        new FooResource("child", { parent: this });
    }
}

new ComponentResource("comp");
