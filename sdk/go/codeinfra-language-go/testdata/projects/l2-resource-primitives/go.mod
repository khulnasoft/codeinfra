module l2-resource-primitives

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-primitive/sdk/go/v7 v7.0.0
)

replace example.com/codeinfra-primitive/sdk/go/v7 => /ROOT/artifacts/example.com_codeinfra-primitive_sdk_go_v7

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
