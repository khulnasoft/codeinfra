package main

import (
	"os"
	"time"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		os.WriteFile("ready", []byte("✅"), 0644)
		time.Sleep(10 * time.Second)
		return nil
	})
}
