(programgen)=
# Programs

Programs may be generated from a program written in a supported source language.
This process, often referred to as "programgen", exists primarily to support the
[`codeinfra convert`](https://www.codeinfra.com/docs/cli/commands/codeinfra_convert/)
command, though it is also used in other places such as [language conformance
testing](language-conformance-tests). Programgen begins by using a [converter
plugin](plugins) to convert from a source language into [Codeinfra Configuration
Language](pcl) (PCL), before generating a program in the target language from
this PCL definition. The PCL-to-target-language conversion is exposed by a
[language host](language-hosts) through two endpoints:

* [](codeinfrarpc.LanguageRuntime.GenerateProgram), which typically wraps the
  `GenerateProgram` Go function in the relevant `gen_program.go` in this package
  -- see e.g. <gh-file:codeinfra#pkg/codegen/nodejs/gen_program.go> for NodeJS or
  <gh-file:codeinfra#pkg/codegen/python/gen_program.go> for Python.
  `GenerateProgram` only generates source files (e.g. `.ts`, `.py`) etc.
* [](codeinfrarpc.LanguageRuntime.GenerateProject), which again typically wraps the
  `GenerateProject` Go function (itself typically found alongside
  `GenerateProgram`). `GenerateProject` generates a complete project directory,
  including both a `Codeinfra.yaml` file and language-appropriate files such as
  `package.json`, `requirements.txt`, etc.
