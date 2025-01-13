// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2016-2023, KhulnaSoft Ltd.
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
var codeinfra_callback_pb = require('./callback_pb.js');

function serialize_codeinfrarpc_CallbackInvokeRequest(arg) {
  if (!(arg instanceof codeinfra_callback_pb.CallbackInvokeRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.CallbackInvokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_CallbackInvokeRequest(buffer_arg) {
  return codeinfra_callback_pb.CallbackInvokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_CallbackInvokeResponse(arg) {
  if (!(arg instanceof codeinfra_callback_pb.CallbackInvokeResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.CallbackInvokeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_CallbackInvokeResponse(buffer_arg) {
  return codeinfra_callback_pb.CallbackInvokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Callbacks is a service for invoking functions in one runtime from other processes.
var CallbacksService = exports.CallbacksService = {
  // Invoke invokes a given callback, identified by its token.
invoke: {
    path: '/codeinfrarpc.Callbacks/Invoke',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_callback_pb.CallbackInvokeRequest,
    responseType: codeinfra_callback_pb.CallbackInvokeResponse,
    requestSerialize: serialize_codeinfrarpc_CallbackInvokeRequest,
    requestDeserialize: deserialize_codeinfrarpc_CallbackInvokeRequest,
    responseSerialize: serialize_codeinfrarpc_CallbackInvokeResponse,
    responseDeserialize: deserialize_codeinfrarpc_CallbackInvokeResponse,
  },
};

exports.CallbacksClient = grpc.makeGenericClientConstructor(CallbacksService);
