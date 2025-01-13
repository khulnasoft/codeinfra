// Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Resource extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #6 - Nested parents changing types
class ComponentSix extends codeinfra.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:ComponentSix-v0", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}

class ComponentSixParent extends codeinfra.ComponentResource {
    child: ComponentSix;
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:module:ComponentSixParent-v0", name, {}, opts);
        this.child = new ComponentSix("child", {parent: this});
    }
}

const comp4 = new ComponentSixParent("comp6");
