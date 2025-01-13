module l2-failed-create-continue-on-error

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-fail_on_create/sdk/go/v4 v4.0.0
	example.com/codeinfra-simple/sdk/go/v2 v2.0.0
)

replace example.com/codeinfra-fail_on_create/sdk/go/v4 => /ROOT/artifacts/example.com_codeinfra-fail_on_create_sdk_go_v4

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3

replace example.com/codeinfra-simple/sdk/go/v2 => /ROOT/artifacts/example.com_codeinfra-simple_sdk_go_v2
