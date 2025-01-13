module l2-primitive-ref

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-primitive-ref/sdk/go/v11 v11.0.0
)

replace example.com/codeinfra-primitive-ref/sdk/go/v11 => /ROOT/artifacts/example.com_codeinfra-primitive-ref_sdk_go_v11

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
