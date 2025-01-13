// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

// Step 3: Run a query during `codeinfra query`.
codeinfra.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")
    .groupBy<string, codeinfra.Resource>(r => (<any>r).__codeinfraType)
    .all(async function(group) {
        const count = await group.count();
        if (group.key === "codeinfra-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);
        return (
            group.key === "codeinfra-nodejs:dynamic:Resource" ||
            group.key === "codeinfra:providers:codeinfra-nodejs" ||
            group.key === "codeinfra:codeinfra:Stack"
        );
    })
    .then(res => {
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
