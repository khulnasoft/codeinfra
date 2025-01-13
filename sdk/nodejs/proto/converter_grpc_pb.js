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
var codeinfra_converter_pb = require('./converter_pb.js');
var codeinfra_codegen_hcl_pb = require('./codegen/hcl_pb.js');

function serialize_codeinfrarpc_ConvertProgramRequest(arg) {
  if (!(arg instanceof codeinfra_converter_pb.ConvertProgramRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.ConvertProgramRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ConvertProgramRequest(buffer_arg) {
  return codeinfra_converter_pb.ConvertProgramRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ConvertProgramResponse(arg) {
  if (!(arg instanceof codeinfra_converter_pb.ConvertProgramResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.ConvertProgramResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ConvertProgramResponse(buffer_arg) {
  return codeinfra_converter_pb.ConvertProgramResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ConvertStateRequest(arg) {
  if (!(arg instanceof codeinfra_converter_pb.ConvertStateRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.ConvertStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ConvertStateRequest(buffer_arg) {
  return codeinfra_converter_pb.ConvertStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_ConvertStateResponse(arg) {
  if (!(arg instanceof codeinfra_converter_pb.ConvertStateResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.ConvertStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_ConvertStateResponse(buffer_arg) {
  return codeinfra_converter_pb.ConvertStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Converter is a service for converting between other ecosystems and Codeinfra.
// This is currently unstable and experimental.
var ConverterService = exports.ConverterService = {
  // ConvertState converts state from the target ecosystem into a form that can be imported into Codeinfra.
convertState: {
    path: '/codeinfrarpc.Converter/ConvertState',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_converter_pb.ConvertStateRequest,
    responseType: codeinfra_converter_pb.ConvertStateResponse,
    requestSerialize: serialize_codeinfrarpc_ConvertStateRequest,
    requestDeserialize: deserialize_codeinfrarpc_ConvertStateRequest,
    responseSerialize: serialize_codeinfrarpc_ConvertStateResponse,
    responseDeserialize: deserialize_codeinfrarpc_ConvertStateResponse,
  },
  // ConvertProgram converts a program from the target ecosystem into a form that can be used with Codeinfra.
convertProgram: {
    path: '/codeinfrarpc.Converter/ConvertProgram',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_converter_pb.ConvertProgramRequest,
    responseType: codeinfra_converter_pb.ConvertProgramResponse,
    requestSerialize: serialize_codeinfrarpc_ConvertProgramRequest,
    requestDeserialize: deserialize_codeinfrarpc_ConvertProgramRequest,
    responseSerialize: serialize_codeinfrarpc_ConvertProgramResponse,
    responseDeserialize: deserialize_codeinfrarpc_ConvertProgramResponse,
  },
};

exports.ConverterClient = grpc.makeGenericClientConstructor(ConverterService);
