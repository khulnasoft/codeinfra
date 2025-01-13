import * as codeinfra from "@codeinfra/codeinfra";

const config = new codeinfra.Config();
const value = config.require("value");
const tags = config.getObject<Record<string, string>>("tags") || {
    [`interpolated/${value}`]: "value",
};
