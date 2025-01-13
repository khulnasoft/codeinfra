// This tests that invokes cannot depend on things which are not resources.

let codeinfra = require("../../../../../");

const dependsOn = codeinfra.output(Promise.resolve([Promise.resolve(1)]));
codeinfra.runtime.invokeOutput("test:index:echo", {}, { dependsOn });
