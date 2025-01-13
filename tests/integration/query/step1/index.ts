// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Resource } from "./resource";

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
