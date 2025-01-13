let codeinfra = require("../../../../../");

codeinfra.log.info("info message");
codeinfra.log.warn("warning message");
codeinfra.log.error("error message");

class FakeResource extends codeinfra.CustomResource {
    constructor(name) {
        super("test:index:FakeResource", name);
    }
}

const res = new FakeResource("test");
codeinfra.log.info("attached to resource", res);
codeinfra.log.info("with streamid", res, 42);
