module l2-resource-alpha

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-alpha/sdk/go/v3 v3.0.0-alpha.1.internal
)

replace example.com/codeinfra-alpha/sdk/go/v3 => /ROOT/artifacts/example.com_codeinfra-alpha_sdk_go_v3

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
