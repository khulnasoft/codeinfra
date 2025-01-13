module.exports = () => {
    let codeinfra = require("../../../../../");

    class MyResource extends codeinfra.CustomResource {
        constructor(name) {
            super("test:index:MyResource", name);
        }
    }

    new MyResource("testResource1");
    return { a: 1 };
};
