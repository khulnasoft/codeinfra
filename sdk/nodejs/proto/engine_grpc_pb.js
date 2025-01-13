// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2016-2018, KhulnaSoft Ltd.
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
var codeinfra_engine_pb = require('./engine_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRootResourceRequest(arg) {
  if (!(arg instanceof codeinfra_engine_pb.GetRootResourceRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRootResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRootResourceRequest(buffer_arg) {
  return codeinfra_engine_pb.GetRootResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRootResourceResponse(arg) {
  if (!(arg instanceof codeinfra_engine_pb.GetRootResourceResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRootResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRootResourceResponse(buffer_arg) {
  return codeinfra_engine_pb.GetRootResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_LogRequest(arg) {
  if (!(arg instanceof codeinfra_engine_pb.LogRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.LogRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_LogRequest(buffer_arg) {
  return codeinfra_engine_pb.LogRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_SetRootResourceRequest(arg) {
  if (!(arg instanceof codeinfra_engine_pb.SetRootResourceRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.SetRootResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_SetRootResourceRequest(buffer_arg) {
  return codeinfra_engine_pb.SetRootResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_SetRootResourceResponse(arg) {
  if (!(arg instanceof codeinfra_engine_pb.SetRootResourceResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.SetRootResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_SetRootResourceResponse(buffer_arg) {
  return codeinfra_engine_pb.SetRootResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_StartDebuggingRequest(arg) {
  if (!(arg instanceof codeinfra_engine_pb.StartDebuggingRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.StartDebuggingRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_StartDebuggingRequest(buffer_arg) {
  return codeinfra_engine_pb.StartDebuggingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


// Engine is an auxiliary service offered to language and resource provider plugins. Its main purpose today is
// to serve as a common logging endpoint, but it also serves as a state storage mechanism for language hosts
// that can't store their own global state.
var EngineService = exports.EngineService = {
  // Log logs a global message in the engine, including errors and warnings.
log: {
    path: '/codeinfrarpc.Engine/Log',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_engine_pb.LogRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_codeinfrarpc_LogRequest,
    requestDeserialize: deserialize_codeinfrarpc_LogRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // GetRootResource gets the URN of the root resource, the resource that should be the root of all
// otherwise-unparented resources.
getRootResource: {
    path: '/codeinfrarpc.Engine/GetRootResource',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_engine_pb.GetRootResourceRequest,
    responseType: codeinfra_engine_pb.GetRootResourceResponse,
    requestSerialize: serialize_codeinfrarpc_GetRootResourceRequest,
    requestDeserialize: deserialize_codeinfrarpc_GetRootResourceRequest,
    responseSerialize: serialize_codeinfrarpc_GetRootResourceResponse,
    responseDeserialize: deserialize_codeinfrarpc_GetRootResourceResponse,
  },
  // SetRootResource sets the URN of the root resource.
setRootResource: {
    path: '/codeinfrarpc.Engine/SetRootResource',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_engine_pb.SetRootResourceRequest,
    responseType: codeinfra_engine_pb.SetRootResourceResponse,
    requestSerialize: serialize_codeinfrarpc_SetRootResourceRequest,
    requestDeserialize: deserialize_codeinfrarpc_SetRootResourceRequest,
    responseSerialize: serialize_codeinfrarpc_SetRootResourceResponse,
    responseDeserialize: deserialize_codeinfrarpc_SetRootResourceResponse,
  },
  // StartDebugging indicates to the engine that the program has started under a debugger, and the engine
// should notify the user of how to connect to the debugger.
startDebugging: {
    path: '/codeinfrarpc.Engine/StartDebugging',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_engine_pb.StartDebuggingRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_codeinfrarpc_StartDebuggingRequest,
    requestDeserialize: deserialize_codeinfrarpc_StartDebuggingRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.EngineClient = grpc.makeGenericClientConstructor(EngineService);
