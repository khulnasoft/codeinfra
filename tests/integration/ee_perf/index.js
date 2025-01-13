// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

"use strict";
const codeinfra = require("@codeinfra/codeinfra");

const config = new codeinfra.Config();
const iterations = config.getNumber("count") || 1000;

// Emit many, many diagnostic events from the engine to stress test the
// ability to record those events on the Codeinfra Service.
console.log("Starting to spam a bunch of diagnostic messages...");
for (let i = 0; i < iterations; i++) {
    console.log(`${i}: The current time is ${new Date()}`);
}
console.log("done");
