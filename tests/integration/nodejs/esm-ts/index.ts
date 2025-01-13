// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as fs from "fs";
import { x } from "./other.js"; // this is the "by design" way to do this, even in TS

// Use top-level await
await new Promise(r => setTimeout(r, 2000));

export const res = fs.readFileSync("Codeinfra.yaml").toString();
export const otherx = x;
