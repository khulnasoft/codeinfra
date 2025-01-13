module l2-resource-secret

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-secret/sdk/go/v14 v14.0.0
)

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3

replace example.com/codeinfra-secret/sdk/go/v14 => /ROOT/artifacts/example.com_codeinfra-secret_sdk_go_v14
