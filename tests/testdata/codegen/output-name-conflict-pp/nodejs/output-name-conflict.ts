import * as codeinfra from "@codeinfra/codeinfra";

export = async () => {
    const config = new codeinfra.Config();
    const cidrBlock = config.get("cidrBlock") || "Test config variable";
    return {
        cidrBlock: cidrBlock,
    };
}
