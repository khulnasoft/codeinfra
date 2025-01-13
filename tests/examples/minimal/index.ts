// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import { Config } from "@codeinfra/codeinfra";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

