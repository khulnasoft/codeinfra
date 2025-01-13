// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class MyComponent extends codeinfra.ComponentResource {
    constructor(name: string) {
        super("test:index:MyComponent", name);
    }
}

codeinfra.log.debug("A debug message");

new MyComponent("mycomponent");
