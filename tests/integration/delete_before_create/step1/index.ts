// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 42 });

// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: codeinfra.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property.
const e = new Resource("dependent-4", { state: codeinfra.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
