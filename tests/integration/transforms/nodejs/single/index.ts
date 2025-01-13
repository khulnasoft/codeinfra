// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Random } from "./random";

codeinfra.runtime.registerStackTransform(async ({ type, props, opts }) => {
    console.log("stack transform");
    return undefined;
});

new Random("res1", { length: codeinfra.secret(5) });