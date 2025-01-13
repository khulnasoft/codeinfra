(sdkgen)=
# SDKs

*[Provider](providers) SDKs* ("software development kits") are generated from a
[Codeinfra Schema](schema) definition. Often referred to as "SDKgen", this process
is used by the myriad providers supported by Codeinfra to expose their resources,
components, and functions in an idiomatic way for a given language. SDKgen is
generally exposed through the [](codeinfrarpc.LanguageRuntime.GeneratePackage)
method of a [language host](language-hosts), which in turn is exposed by the
CLI's [`codeinfra package
gen-sdk`](https://www.codeinfra.com/docs/cli/commands/codeinfra_package_gen-sdk/)
command. At a code level, SDKgen starts with the relevant `GeneratePackage` Go
function in `gen.go` -- see <gh-file:codeinfra#pkg/codegen/nodejs/gen.go> for
NodeJS, <gh-file:codeinfra#pkg/codegen/python/gen.go> for Python, and so on.

:::{note}
The `codeinfra package gen-sdk` command is not really intended to be used by
external users or customers, and instead offers a convenient interface for
generating provider SDKs as part of e.g. the various provider CI jobs used
to automate provider build and release processes.
:::
