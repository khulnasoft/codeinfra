module l2-parameterized-resource

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-subpackage/sdk/go/v2 v2.0.0
)

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3

replace example.com/codeinfra-subpackage/sdk/go/v2 => /ROOT/artifacts/example.com_codeinfra-subpackage_sdk_go_v2
