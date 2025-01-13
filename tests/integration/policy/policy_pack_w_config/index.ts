// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

import * as policy from "@codeinfra/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
