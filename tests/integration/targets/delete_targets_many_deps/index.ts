// Copyright 2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

interface MyResourceArgs {
    input?: codeinfra.Input<string> | undefined;
}

class MyResource extends codeinfra.dynamic.Resource {
    static latestId: number = 0;

    constructor(name, args?: MyResourceArgs, opts?: codeinfra.CustomResourceOptions) {
        super({
            async create(inputs: any) {
                return { id: (MyResource.latestId++).toString() };
            },
        }, name, args || {}, opts);
    }
}

// Create a chain of resources, such that attempting to delete
// one will fail due to numerous dependency violations. This includes
// both implicit and explicit, as well as parent/child, dependencies.

// A
// B (impl depends on A)
// C (expl depends on A)
// D (impl depends on B)
// E (expl depends on B)
// F (child of A)
// G (child of B)
// H (expl depends on A, B, impl depends on C, D, child of F)

const a = new MyResource("a");
const b = new MyResource("b", { input: a.urn });
const c = new MyResource("c", { }, { dependsOn: a });
const d = new MyResource("d", { input: b.urn });
const e = new MyResource("e", { }, { dependsOn: b });
const f = new MyResource("f", { }, { parent: a });
const g = new MyResource("g", { }, { parent: b });
const h = new MyResource("h",
    { input: codeinfra.all(([c.urn, d.urn])).apply(([curn, _])=>curn) },
    {
        dependsOn: [a, b],
        parent: f,
    },
);
