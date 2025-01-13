// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as codeinfra from "@codeinfra/codeinfra";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}

(async function() {
    // Just test that basic config works.
    const config = new codeinfra.Config();

    const outsideCapture = await codeinfra.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await codeinfra.runtime.serializeFunction(() => {
        const config = new codeinfra.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();
})()
