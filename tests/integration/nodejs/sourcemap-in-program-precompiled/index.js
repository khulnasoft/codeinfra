"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.willThrow = willThrow;
var codeinfra = require("@codeinfra/codeinfra");
function willThrow() {
    if (true) {
        codeinfra.log.error("Oh no!");
        throw new Error("this is a test error");
    }
}
willThrow();
//# sourceMappingURL=index.js.map