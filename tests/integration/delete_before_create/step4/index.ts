// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Resource } from "./resource";

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: codeinfra.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
