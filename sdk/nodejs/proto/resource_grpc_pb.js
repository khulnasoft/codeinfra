// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2016-2022, KhulnaSoft Ltd.
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
//
'use strict';
var grpc = require('@grpc/grpc-js');
var codeinfra_resource_pb = require('./resource_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
var codeinfra_provider_pb = require('./provider_pb.js');
var codeinfra_alias_pb = require('./alias_pb.js');
var codeinfra_source_pb = require('./source_pb.js');
var codeinfra_callback_pb = require('./callback_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_CallResponse(arg) {
  if (!(arg instanceof codeinfra_provider_pb.CallResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.CallResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_CallResponse(buffer_arg) {
  return codeinfra_provider_pb.CallResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_Callback(arg) {
  if (!(arg instanceof codeinfra_callback_pb.Callback)) {
    throw new Error('Expected argument of type codeinfrarpc.Callback');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_Callback(buffer_arg) {
  return codeinfra_callback_pb.Callback.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_InvokeResponse(arg) {
  if (!(arg instanceof codeinfra_provider_pb.InvokeResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.InvokeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_InvokeResponse(buffer_arg) {
  return codeinfra_provider_pb.InvokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ReadResourceRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.ReadResourceRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.ReadResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ReadResourceRequest(buffer_arg) {
  return codeinfra_resource_pb.ReadResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ReadResourceResponse(arg) {
  if (!(arg instanceof codeinfra_resource_pb.ReadResourceResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.ReadResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ReadResourceResponse(buffer_arg) {
  return codeinfra_resource_pb.ReadResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RegisterPackageRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.RegisterPackageRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RegisterPackageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RegisterPackageRequest(buffer_arg) {
  return codeinfra_resource_pb.RegisterPackageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RegisterPackageResponse(arg) {
  if (!(arg instanceof codeinfra_resource_pb.RegisterPackageResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.RegisterPackageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RegisterPackageResponse(buffer_arg) {
  return codeinfra_resource_pb.RegisterPackageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RegisterResourceOutputsRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.RegisterResourceOutputsRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RegisterResourceOutputsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RegisterResourceOutputsRequest(buffer_arg) {
  return codeinfra_resource_pb.RegisterResourceOutputsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RegisterResourceRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.RegisterResourceRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RegisterResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RegisterResourceRequest(buffer_arg) {
  return codeinfra_resource_pb.RegisterResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RegisterResourceResponse(arg) {
  if (!(arg instanceof codeinfra_resource_pb.RegisterResourceResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.RegisterResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RegisterResourceResponse(buffer_arg) {
  return codeinfra_resource_pb.RegisterResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ResourceCallRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.ResourceCallRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.ResourceCallRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ResourceCallRequest(buffer_arg) {
  return codeinfra_resource_pb.ResourceCallRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ResourceInvokeRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.ResourceInvokeRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.ResourceInvokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ResourceInvokeRequest(buffer_arg) {
  return codeinfra_resource_pb.ResourceInvokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_SupportsFeatureRequest(arg) {
  if (!(arg instanceof codeinfra_resource_pb.SupportsFeatureRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.SupportsFeatureRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_SupportsFeatureRequest(buffer_arg) {
  return codeinfra_resource_pb.SupportsFeatureRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_SupportsFeatureResponse(arg) {
  if (!(arg instanceof codeinfra_resource_pb.SupportsFeatureResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.SupportsFeatureResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_SupportsFeatureResponse(buffer_arg) {
  return codeinfra_resource_pb.SupportsFeatureResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// ResourceMonitor is the interface a source uses to talk back to the planning monitor orchestrating the execution.
var ResourceMonitorService = exports.ResourceMonitorService = {
  supportsFeature: {
    path: '/codeinfrarpc.ResourceMonitor/SupportsFeature',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.SupportsFeatureRequest,
    responseType: codeinfra_resource_pb.SupportsFeatureResponse,
    requestSerialize: serialize_codeinfrarpc_SupportsFeatureRequest,
    requestDeserialize: deserialize_codeinfrarpc_SupportsFeatureRequest,
    responseSerialize: serialize_codeinfrarpc_SupportsFeatureResponse,
    responseDeserialize: deserialize_codeinfrarpc_SupportsFeatureResponse,
  },
  invoke: {
    path: '/codeinfrarpc.ResourceMonitor/Invoke',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.ResourceInvokeRequest,
    responseType: codeinfra_provider_pb.InvokeResponse,
    requestSerialize: serialize_codeinfrarpc_ResourceInvokeRequest,
    requestDeserialize: deserialize_codeinfrarpc_ResourceInvokeRequest,
    responseSerialize: serialize_codeinfrarpc_InvokeResponse,
    responseDeserialize: deserialize_codeinfrarpc_InvokeResponse,
  },
  streamInvoke: {
    path: '/codeinfrarpc.ResourceMonitor/StreamInvoke',
    requestStream: false,
    responseStream: true,
    requestType: codeinfra_resource_pb.ResourceInvokeRequest,
    responseType: codeinfra_provider_pb.InvokeResponse,
    requestSerialize: serialize_codeinfrarpc_ResourceInvokeRequest,
    requestDeserialize: deserialize_codeinfrarpc_ResourceInvokeRequest,
    responseSerialize: serialize_codeinfrarpc_InvokeResponse,
    responseDeserialize: deserialize_codeinfrarpc_InvokeResponse,
  },
  call: {
    path: '/codeinfrarpc.ResourceMonitor/Call',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.ResourceCallRequest,
    responseType: codeinfra_provider_pb.CallResponse,
    requestSerialize: serialize_codeinfrarpc_ResourceCallRequest,
    requestDeserialize: deserialize_codeinfrarpc_ResourceCallRequest,
    responseSerialize: serialize_codeinfrarpc_CallResponse,
    responseDeserialize: deserialize_codeinfrarpc_CallResponse,
  },
  readResource: {
    path: '/codeinfrarpc.ResourceMonitor/ReadResource',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.ReadResourceRequest,
    responseType: codeinfra_resource_pb.ReadResourceResponse,
    requestSerialize: serialize_codeinfrarpc_ReadResourceRequest,
    requestDeserialize: deserialize_codeinfrarpc_ReadResourceRequest,
    responseSerialize: serialize_codeinfrarpc_ReadResourceResponse,
    responseDeserialize: deserialize_codeinfrarpc_ReadResourceResponse,
  },
  registerResource: {
    path: '/codeinfrarpc.ResourceMonitor/RegisterResource',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.RegisterResourceRequest,
    responseType: codeinfra_resource_pb.RegisterResourceResponse,
    requestSerialize: serialize_codeinfrarpc_RegisterResourceRequest,
    requestDeserialize: deserialize_codeinfrarpc_RegisterResourceRequest,
    responseSerialize: serialize_codeinfrarpc_RegisterResourceResponse,
    responseDeserialize: deserialize_codeinfrarpc_RegisterResourceResponse,
  },
  registerResourceOutputs: {
    path: '/codeinfrarpc.ResourceMonitor/RegisterResourceOutputs',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.RegisterResourceOutputsRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_codeinfrarpc_RegisterResourceOutputsRequest,
    requestDeserialize: deserialize_codeinfrarpc_RegisterResourceOutputsRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Register a resource transform for the stack
registerStackTransform: {
    path: '/codeinfrarpc.ResourceMonitor/RegisterStackTransform',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_callback_pb.Callback,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_codeinfrarpc_Callback,
    requestDeserialize: deserialize_codeinfrarpc_Callback,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Register an invoke transform for the stack
registerStackInvokeTransform: {
    path: '/codeinfrarpc.ResourceMonitor/RegisterStackInvokeTransform',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_callback_pb.Callback,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_codeinfrarpc_Callback,
    requestDeserialize: deserialize_codeinfrarpc_Callback,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  registerPackage: {
    path: '/codeinfrarpc.ResourceMonitor/RegisterPackage',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_resource_pb.RegisterPackageRequest,
    responseType: codeinfra_resource_pb.RegisterPackageResponse,
    requestSerialize: serialize_codeinfrarpc_RegisterPackageRequest,
    requestDeserialize: deserialize_codeinfrarpc_RegisterPackageRequest,
    responseSerialize: serialize_codeinfrarpc_RegisterPackageResponse,
    responseDeserialize: deserialize_codeinfrarpc_RegisterPackageResponse,
  },
};

exports.ResourceMonitorClient = grpc.makeGenericClientConstructor(ResourceMonitorService);
