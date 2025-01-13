// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";
import { Component } from "./component";

let currentID = 0;

class Resource extends dynamic.Resource {
    constructor(name: string, opts?: codeinfra.CustomResourceOptions) {
        const provider = {
            create: async (inputs: any) => ({
                id: (currentID++).toString(),
                outs: undefined,
            }),
        };

        super(provider, name, {}, opts);
    }
}

const resource = new Resource("resource");

const component = new Component("component", {
	message: resource.id.apply(v => `message ${v}`),
	nested: {
		value: resource.id.apply(v => `nested.value ${v}`),
	}
});
