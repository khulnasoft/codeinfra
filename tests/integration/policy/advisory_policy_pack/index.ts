import { PolicyPack, validateResourceOfType } from "@codeinfra/policy";

new PolicyPack("typescript", {
    policies: [{
        name: "advisory-policy-pack",
        description: "Failing advisory policy pack for testing",
        enforcementLevel: "advisory",
        validateStack: (stack: any, reportViolation: any) => {
	    reportViolation("foobar");
        },
    }],
});
