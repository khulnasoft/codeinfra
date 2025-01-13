import * as codeinfra from "@codeinfra/codeinfra";
import * as splat from "@codeinfra/splat";

const allKeys = splat.getSshKeys({});
const main = new splat.Server("main", {sshKeys: allKeys.then(allKeys => allKeys.sshKeys.map(__item => __item.name))});
