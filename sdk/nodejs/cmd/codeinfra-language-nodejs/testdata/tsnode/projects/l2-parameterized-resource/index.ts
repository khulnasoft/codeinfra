import * as codeinfra from "@codeinfra/codeinfra";
import * as subpackage from "@codeinfra/subpackage";

// The resource name is based on the parameter value
const example = new subpackage.HelloWorld("example", {});
export const parameterValue = example.parameterValue;
