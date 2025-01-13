// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Random } from "./random";

class MyComponent extends codeinfra.ComponentResource {
    child: Random;
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:component:MyComponent", name, {}, opts);
        this.child = new Random(`${name}-child`, { length: 5 }, {
            parent: this,
            additionalSecretOutputs: ["length"],
        });
        this.registerOutputs({});
    }
}

// Scenario #1 - apply a transformation to a CustomResource
const res1 = new Random("res1", { length: 5 }, {
    transformations: [
        ({ props, opts }) => {
            console.log("res1 transformation");
            return {
                props: props,
                opts: codeinfra.mergeOptions(opts, { additionalSecretOutputs: ["result"] }),
            };
        },
    ],
});

// Scenario #2 - apply a transformation to a Component to transform it's children
const res2 = new MyComponent("res2", {
    transformations: [
        ({ type, props, opts }) => {
            console.log("res2 transformation");
            if (type === "testprovider:index:Random") {
                return {
                    props: { prefix: "newDefault", ...props },
                    opts: codeinfra.mergeOptions(opts, { additionalSecretOutputs: ["result"] }),
                };
            }
        },
    ],
});

// Scenario #3 - apply a transformation to the Stack to transform all (future) resources in the stack
codeinfra.runtime.registerStackTransformation(({ type, props, opts }) => {
    console.log("stack transformation");
    if (type === "testprovider:index:Random") {
        return {
            props: { ...props, prefix: "stackDefault" },
            opts: codeinfra.mergeOptions(opts, { additionalSecretOutputs: ["result"] }),
        };
    }
});

const res3 = new Random("res3", { length: 5 });

// Scenario #4 - transformations are applied in order of decreasing specificity
// 1. (not in this example) Child transformation
// 2. First parent transformation
// 3. Second parent transformation
// 4. Stack transformation
const res4 = new MyComponent("res4", {
    transformations: [
        ({ type, props, opts }) => {
            console.log("res4 transformation");
            if (type === "testprovider:index:Random") {
                return {
                    props: { ...props, prefix: "default1" },
                    opts,
                };
            }
        },
        ({ type, props, opts }) => {
            console.log("res4 transformation 2");
            if (type === "testprovider:index:Random") {
                return {
                    props: { ...props, prefix: "default2" },
                    opts,
                };
            }
        },
    ],
});

// Scenario #5 - cross-resource transformations that inject dependencies on one resource into another.
class MyOtherComponent extends codeinfra.ComponentResource {
    child1: Random;
    child2: Random;
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:component:MyOtherComponent", name, {}, opts);
        this.child1 = new Random(`${name}-child1`, { length: 5 }, { parent: this });
        this.child2 = new Random(`${name}-child2`, { length: 5 }, { parent: this });
        this.registerOutputs({});
    }
}

const transformChild1DependsOnChild2: codeinfra.ResourceTransformation = (() => {
    console.log("res5 transformation")

    // Create a promise that wil be resolved once we find child2.  This is needed because we do not
    // know what order we will see the resource registrations of child1 and child2.
    let child2Found: (res: codeinfra.Resource) => void;
    const child2 = new Promise<codeinfra.Resource>((res) => child2Found = res);

    // Return a transformation which will rewrite child1 to depend on the promise for child2, and
    // will resolve that promise when it finds child2.
    return (args: codeinfra.ResourceTransformationArgs) => {
        if (args.name.endsWith("-child2")) {
            // Resolve the child2 promise with the child2 resource.
            child2Found(args.resource);
            return undefined;
        } else if (args.name.endsWith("-child1")) {
            // Overwrite the `prefix` to child2 with a dependency on the `length` from child1.
            const child2Result = codeinfra.output(args.props["length"]).apply(async (input) => {
                if (input !== 5) {
                    // Not strictly necessary - but shows we can confirm invariants we expect to be
                    // true.
                    throw new Error("unexpected input value");
                }
                return child2.then(c2Res => c2Res["result"]);
            });
            // Finally - overwrite the input of child2.
            return {
                props: { ...args.props, prefix: child2Result },
                opts: args.opts,
            };
        }
    };
})();

const res5 = new MyOtherComponent("res5", {
    transformations: [ transformChild1DependsOnChild2 ],
});
