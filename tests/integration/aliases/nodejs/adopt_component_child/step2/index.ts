// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class FooResource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:FooResource", name, {}, opts);
    }
}

class ComponentResource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:ComponentResource", name, {}, opts);
        new FooResource("child", {
            parent: this,
            aliases: [{ parent: codeinfra.rootStackResource }],
        });
    }
}

new ComponentResource("comp");
