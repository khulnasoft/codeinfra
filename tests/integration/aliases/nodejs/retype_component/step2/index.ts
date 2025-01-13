// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Resource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends codeinfra.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");
