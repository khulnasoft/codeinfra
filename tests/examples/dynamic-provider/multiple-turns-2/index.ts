// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: codeinfra.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: codeinfra.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<codeinfra.Output<string>> {
    await sleep(1000);

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
