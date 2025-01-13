// Copyright 2016-2021, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:deadcode
package codeinfra

import (
	"context"
	_ "unsafe" // unsafe is needed to use go:linkname

	codeinfrarpc "github.com/khulnasoft/codeinfra/sdk/v3/proto/go"

	"google.golang.org/grpc"
)

// We want the public provider-related APIs to be exported from the provider package, but need to make use of unexported
// functionality in this package for their implementations. To achieve this, go:linkname is used to make the following
// functions available in the provider package.

//go:linkname linkedConstruct github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedConstruct
func linkedConstruct(ctx context.Context, req *codeinfrarpc.ConstructRequest, engineConn *grpc.ClientConn,
	constructF constructFunc,
) (*codeinfrarpc.ConstructResponse, error) {
	return construct(ctx, req, engineConn, constructF)
}

//go:linkname linkedConstructInputsMap github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedConstructInputsMap
func linkedConstructInputsMap(ctx *Context, inputs map[string]interface{}) (Map, error) {
	return constructInputsMap(ctx, inputs)
}

//go:linkname linkedConstructInputsCopyTo github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedConstructInputsCopyTo
func linkedConstructInputsCopyTo(ctx *Context, inputs map[string]interface{}, args interface{}) error {
	return constructInputsCopyTo(ctx, inputs, args)
}

//go:linkname linkedNewConstructResult github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedNewConstructResult
func linkedNewConstructResult(resource ComponentResource) (URNInput, Input, error) {
	return newConstructResult(resource)
}

//go:linkname linkedCall github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedCall
func linkedCall(ctx context.Context, req *codeinfrarpc.CallRequest, engineConn *grpc.ClientConn,
	callF callFunc,
) (*codeinfrarpc.CallResponse, error) {
	return call(ctx, req, engineConn, callF)
}

//go:linkname linkedCallArgsCopyTo github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedCallArgsCopyTo
func linkedCallArgsCopyTo(ctx *Context, source map[string]interface{}, args interface{}) (Resource, error) {
	return callArgsCopyTo(ctx, source, args)
}

//go:linkname linkedCallArgsSelf github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedCallArgsSelf
func linkedCallArgsSelf(ctx *Context, source map[string]interface{}) (Resource, error) {
	return callArgsSelf(ctx, source)
}

//go:linkname linkedNewCallResult github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedNewCallResult
func linkedNewCallResult(result interface{}) (Input, error) {
	return newCallResult(result)
}

//go:linkname linkedNewCallFailure github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider.linkedNewCallFailure
func linkedNewCallFailure(property, reason string) interface{} {
	return newCallFailure(property, reason)
}
