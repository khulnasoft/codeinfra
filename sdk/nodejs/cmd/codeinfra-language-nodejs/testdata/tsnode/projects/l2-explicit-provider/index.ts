import * as codeinfra from "@codeinfra/codeinfra";
import * as simple from "@codeinfra/simple";

const prov = new simple.Provider("prov", {});
const res = new simple.Resource("res", {value: true}, {
    provider: prov,
});
