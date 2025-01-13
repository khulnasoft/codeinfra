# Codeinfra Golang SDK

This directory contains support for writing Codeinfra programs in the Go language.  There are two aspects to this:

* `codeinfra/` contains the client language bindings Codeinfra program's code directly against;
* `codeinfra-language-go/` contains the language host plugin that the Codeinfra engine uses to orchestrate updates.

To author a Codeinfra program in Go, simply say so in your `Codeinfra.yaml`

    name: <my-project>
    runtime: go

and ensure you have `codeinfra-language-go` on your path (it is distributed in the Codeinfra download automatically).

By default, the language plugin will use your project's name, `<my-project>`, as the executable that it loads.  This too
must be on your path for the language provider to load it when you run `codeinfra preview` or `codeinfra up`.
