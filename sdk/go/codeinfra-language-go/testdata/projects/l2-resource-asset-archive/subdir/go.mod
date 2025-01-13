module l2-resource-asset-archive

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-asset-archive/sdk/go/v5 v5.0.0
)

replace example.com/codeinfra-asset-archive/sdk/go/v5 => /ROOT/artifacts/example.com_codeinfra-asset-archive_sdk_go_v5

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
