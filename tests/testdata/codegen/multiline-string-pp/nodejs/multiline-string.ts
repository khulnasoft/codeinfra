import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

const foo = new random.RandomShuffle("foo", {inputs: [
    `just one
newline`,
    `foo
bar
baz
qux
quux
qux`,
    `{
    "a": 1,
    "b": 2,
    "c": [
      "foo",
      "bar",
      "baz",
      "qux",
      "quux"
    ]
}
`,
]});
