module l2-resource-config

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-config/sdk/go/v9 v9.0.0
)

replace example.com/codeinfra-config/sdk/go/v9 => /ROOT/artifacts/example.com_codeinfra-config_sdk_go_v9

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
