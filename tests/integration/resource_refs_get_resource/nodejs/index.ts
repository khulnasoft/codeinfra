// Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.

import * as assert from "assert";

import * as codeinfra from "@codeinfra/codeinfra";

class Child extends codeinfra.ComponentResource {
    public readonly message!: codeinfra.Output<string>;
    constructor(name: string, message?: string, opts?: codeinfra.ResourceOptions) {
        const args = { message }
        super("test:index:Child", name, args, opts);
        if (opts?.urn) {
            return;
        }
        this.registerOutputs(args);
    }
}

class Container extends codeinfra.ComponentResource {
    public readonly child!: codeinfra.Output<Child>;
    constructor(name: string, child?: Child, opts?: codeinfra.ResourceOptions) {
        const args = { child };
        super("test:index:Container", name, args, opts);
        if (opts?.urn) {
            return;
        }
        this.registerOutputs(args);
    }
}

async function waitForContainer(container: Container): Promise<void> {
    // Wait for a maximum of 500ms for the resource outputs to be registered.
    const end = Date.now() + 500;
    for (let i = 0; ; i++) {
	let foundURN = false;
	let success = new Promise((resolve, reject) => {
	    container.urn.apply(urn => {
		const roundTrippedContainer = new Container("mycontainer", undefined, { urn })
		roundTrippedContainer.child.apply(c => {
		    if (c != undefined) {
			foundURN = true;
		    }
		    resolve();
		});
	    });
	});
	await success;
	if (foundURN) {
	    break;
	} else if (Date.now() > end) {
	    throw new Error("resource outputs not registered correctly");
	}
    }
}

codeinfra.runtime.registerResourceModule("test", "index", {
    version: "0.0.1",
    construct: (name: string, type: string, urn: string) => {
        switch (type) {
            case "test:index:Child":
                return new Child(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type: ${type}`);
        }
    },
});

const child = new Child("mychild", "hello world!");
const container = new Container("mycontainer", child);

// Wait to make sure registerOutputs has actually finished registering the resource outputs.
//
// registerOutputs works asynchronously, as does registering a component resource.  This
// means registerOutputs is inheritly racy with the resource being read later.  This test explicitly
// tests roundtripping a container component resource, which means we need to read the outputs registered
// outputs are registered before we return the container.  Ideally we should find a way to make this non-racy
// (see the issue linked below)
//
// TODO: make RegisterResourceOutputs not racy [codeinfra/codeinfra#16896]
waitForContainer(container).then(() => {
    console.log(child, container);

    codeinfra.all([child.urn, container.urn]).apply(([childUrn, urn]) => {
	const roundTrippedContainer = new Container("mycontainer", undefined, { urn })
	const roundTrippedContainerChildUrn = roundTrippedContainer.child.apply(c => c.urn);
	const roundTrippedContainerChildMessage = roundTrippedContainer.child.apply(c => c.message);
	return codeinfra.all([childUrn, roundTrippedContainerChildUrn, roundTrippedContainerChildMessage])
	    .apply(([expectedUrn, actualUrn, actualMessage]) => {
		assert.strictEqual(actualUrn, expectedUrn);
		assert.strictEqual(actualMessage, "hello world!");
		return expectedUrn;
	    });
    });
});
