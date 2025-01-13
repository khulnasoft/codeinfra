// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Resource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});
