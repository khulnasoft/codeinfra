module l2-invoke-options

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-simple-invoke/sdk/go/v10 v10.0.0
)

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3

replace example.com/codeinfra-simple-invoke/sdk/go/v10 => /ROOT/artifacts/example.com_codeinfra-simple-invoke_sdk_go_v10
