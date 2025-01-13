import * as codeinfra from "@codeinfra/codeinfra";
import * as goodbye from "@codeinfra/goodbye";

const prov = new goodbye.Provider("prov", {text: "World"});
// The resource name is based on the parameter value
const res = new goodbye.Goodbye("res", {}, {
    provider: prov,
});
export const parameterValue = res.parameterValue;
