module l2-ref-ref

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-ref-ref/sdk/go/v12 v12.0.0
)

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3

replace example.com/codeinfra-ref-ref/sdk/go/v12 => /ROOT/artifacts/example.com_codeinfra-ref-ref_sdk_go_v12
