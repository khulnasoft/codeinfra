import * as codeinfra from "@codeinfra/codeinfra";
import "jest";
import { willThrow } from "."

// Mocks are unused, but previously importing and using the codeinfra package in
// the tests would break source maps
// https://github.com/khulnasoft/codeinfra/issues/9218
codeinfra.runtime.setMocks({
    newResource: (args: codeinfra.runtime.MockResourceArgs) => {
        return {
            id: `${args.name}-id`,
            state: args.inputs,
        };
    },
    call: (args: codeinfra.runtime.MockCallArgs) => {
        return {};
    },
})

it("a failing test so we can inspect the stacktrace reported by jest", async () => {
    willThrow();
});
