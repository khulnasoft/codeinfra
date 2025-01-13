module l2-provider-grpc-config-schema-secret

go 1.20

require (
	github.com/khulnasoft/codeinfra/sdk/v3 v3.30.0
	example.com/codeinfra-config-grpc/sdk/go v1.0.0
)

replace example.com/codeinfra-config-grpc/sdk/go => /ROOT/artifacts/example.com_codeinfra-config-grpc_sdk_go

replace github.com/khulnasoft/codeinfra/sdk/v3 => /ROOT/artifacts/github.com_codeinfra_codeinfra_sdk_v3
