import * as codeinfra from "@codeinfra/codeinfra";
import * as other from "@third-party/codeinfra-other";

const Other = new other.Thing("Other", {idea: "Support Third Party"});
const Question = new other.module.Object("Question", {answer: 42});
const Question2 = new other.module.sub.Object("Question2", {answer: 24});
const Provider = new other.Provider("Provider", {objectProp: {
    prop1: "foo",
    prop2: "bar",
    prop3: "fizz",
}});
