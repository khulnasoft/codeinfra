module l2-large-string

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-large/sdk/go/v4 v4.3.2
)

replace example.com/codeinfra-large/sdk/go/v4 => /ROOT/artifacts/example.com_codeinfra-large_sdk_go_v4

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
