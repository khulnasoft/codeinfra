// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import { Resource } from "./resource";

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
