import codeinfra

codeinfra.export("strVar", "foo")
codeinfra.export("arrVar", [
    "fizz",
    "buzz",
])
codeinfra.export("readme", (lambda path: open(path).read())("./Codeinfra.README.md"))
