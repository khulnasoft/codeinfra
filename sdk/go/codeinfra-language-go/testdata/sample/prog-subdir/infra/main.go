package main

import (
	"github.com/codeinfra/go-dependency-testdata/dep"
	"github.com/codeinfra/go-dependency-testdata/plugin"
)

func main() {
	plugin.Foo()
	dep.Bar()
}
