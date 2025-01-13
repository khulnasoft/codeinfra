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
var codeinfra_language_pb = require('./language_pb.js');
var codeinfra_codegen_hcl_pb = require('./codegen/hcl_pb.js');
var codeinfra_plugin_pb = require('./plugin_pb.js');
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

function serialize_codeinfrarpc_AboutRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.AboutRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.AboutRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_AboutRequest(buffer_arg) {
  return codeinfra_language_pb.AboutRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_AboutResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.AboutResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.AboutResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_AboutResponse(buffer_arg) {
  return codeinfra_language_pb.AboutResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GeneratePackageRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GeneratePackageRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GeneratePackageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GeneratePackageRequest(buffer_arg) {
  return codeinfra_language_pb.GeneratePackageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GeneratePackageResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GeneratePackageResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GeneratePackageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GeneratePackageResponse(buffer_arg) {
  return codeinfra_language_pb.GeneratePackageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GenerateProgramRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GenerateProgramRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GenerateProgramRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GenerateProgramRequest(buffer_arg) {
  return codeinfra_language_pb.GenerateProgramRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GenerateProgramResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GenerateProgramResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GenerateProgramResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GenerateProgramResponse(buffer_arg) {
  return codeinfra_language_pb.GenerateProgramResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GenerateProjectRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GenerateProjectRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GenerateProjectRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GenerateProjectRequest(buffer_arg) {
  return codeinfra_language_pb.GenerateProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GenerateProjectResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GenerateProjectResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GenerateProjectResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GenerateProjectResponse(buffer_arg) {
  return codeinfra_language_pb.GenerateProjectResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetProgramDependenciesRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetProgramDependenciesRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GetProgramDependenciesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetProgramDependenciesRequest(buffer_arg) {
  return codeinfra_language_pb.GetProgramDependenciesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetProgramDependenciesResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetProgramDependenciesResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GetProgramDependenciesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetProgramDependenciesResponse(buffer_arg) {
  return codeinfra_language_pb.GetProgramDependenciesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRequiredPackagesRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetRequiredPackagesRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRequiredPackagesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRequiredPackagesRequest(buffer_arg) {
  return codeinfra_language_pb.GetRequiredPackagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRequiredPackagesResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetRequiredPackagesResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRequiredPackagesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRequiredPackagesResponse(buffer_arg) {
  return codeinfra_language_pb.GetRequiredPackagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRequiredPluginsRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetRequiredPluginsRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRequiredPluginsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRequiredPluginsRequest(buffer_arg) {
  return codeinfra_language_pb.GetRequiredPluginsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_GetRequiredPluginsResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.GetRequiredPluginsResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.GetRequiredPluginsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_GetRequiredPluginsResponse(buffer_arg) {
  return codeinfra_language_pb.GetRequiredPluginsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_InstallDependenciesRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.InstallDependenciesRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.InstallDependenciesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_InstallDependenciesRequest(buffer_arg) {
  return codeinfra_language_pb.InstallDependenciesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_InstallDependenciesResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.InstallDependenciesResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.InstallDependenciesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_InstallDependenciesResponse(buffer_arg) {
  return codeinfra_language_pb.InstallDependenciesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_LanguageHandshakeRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.LanguageHandshakeRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.LanguageHandshakeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_LanguageHandshakeRequest(buffer_arg) {
  return codeinfra_language_pb.LanguageHandshakeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_LanguageHandshakeResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.LanguageHandshakeResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.LanguageHandshakeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_LanguageHandshakeResponse(buffer_arg) {
  return codeinfra_language_pb.LanguageHandshakeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_PackRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.PackRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.PackRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_PackRequest(buffer_arg) {
  return codeinfra_language_pb.PackRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_PackResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.PackResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.PackResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_PackResponse(buffer_arg) {
  return codeinfra_language_pb.PackResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_PluginInfo(arg) {
  if (!(arg instanceof codeinfra_plugin_pb.PluginInfo)) {
    throw new Error('Expected argument of type codeinfrarpc.PluginInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_PluginInfo(buffer_arg) {
  return codeinfra_plugin_pb.PluginInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RunPluginRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.RunPluginRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RunPluginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RunPluginRequest(buffer_arg) {
  return codeinfra_language_pb.RunPluginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RunPluginResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.RunPluginResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.RunPluginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RunPluginResponse(buffer_arg) {
  return codeinfra_language_pb.RunPluginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RunRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.RunRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RunRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RunRequest(buffer_arg) {
  return codeinfra_language_pb.RunRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RunResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.RunResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.RunResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RunResponse(buffer_arg) {
  return codeinfra_language_pb.RunResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RuntimeOptionsRequest(arg) {
  if (!(arg instanceof codeinfra_language_pb.RuntimeOptionsRequest)) {
    throw new Error('Expected argument of type codeinfrarpc.RuntimeOptionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RuntimeOptionsRequest(buffer_arg) {
  return codeinfra_language_pb.RuntimeOptionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_codeinfrarpc_RuntimeOptionsResponse(arg) {
  if (!(arg instanceof codeinfra_language_pb.RuntimeOptionsResponse)) {
    throw new Error('Expected argument of type codeinfrarpc.RuntimeOptionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_codeinfrarpc_RuntimeOptionsResponse(buffer_arg) {
  return codeinfra_language_pb.RuntimeOptionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// LanguageRuntime is the interface that the planning monitor uses to drive execution of an interpreter responsible
// for confguring and creating resource objects.
var LanguageRuntimeService = exports.LanguageRuntimeService = {
  // `Handshake` is the first call made by the engine to a language host. It is used to pass the 
// engine's address to the language host so that it may establish its own connections back,
// and to establish protocol configuration that will be used to communicate between the two parties. 
handshake: {
    path: '/codeinfrarpc.LanguageRuntime/Handshake',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.LanguageHandshakeRequest,
    responseType: codeinfra_language_pb.LanguageHandshakeResponse,
    requestSerialize: serialize_codeinfrarpc_LanguageHandshakeRequest,
    requestDeserialize: deserialize_codeinfrarpc_LanguageHandshakeRequest,
    responseSerialize: serialize_codeinfrarpc_LanguageHandshakeResponse,
    responseDeserialize: deserialize_codeinfrarpc_LanguageHandshakeResponse,
  },
  // GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
getRequiredPlugins: {
    path: '/codeinfrarpc.LanguageRuntime/GetRequiredPlugins',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GetRequiredPluginsRequest,
    responseType: codeinfra_language_pb.GetRequiredPluginsResponse,
    requestSerialize: serialize_codeinfrarpc_GetRequiredPluginsRequest,
    requestDeserialize: deserialize_codeinfrarpc_GetRequiredPluginsRequest,
    responseSerialize: serialize_codeinfrarpc_GetRequiredPluginsResponse,
    responseDeserialize: deserialize_codeinfrarpc_GetRequiredPluginsResponse,
  },
  // GetRequiredPackages computes the complete set of anticipated packages required by a program.
getRequiredPackages: {
    path: '/codeinfrarpc.LanguageRuntime/GetRequiredPackages',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GetRequiredPackagesRequest,
    responseType: codeinfra_language_pb.GetRequiredPackagesResponse,
    requestSerialize: serialize_codeinfrarpc_GetRequiredPackagesRequest,
    requestDeserialize: deserialize_codeinfrarpc_GetRequiredPackagesRequest,
    responseSerialize: serialize_codeinfrarpc_GetRequiredPackagesResponse,
    responseDeserialize: deserialize_codeinfrarpc_GetRequiredPackagesResponse,
  },
  // Run executes a program and returns its result.
run: {
    path: '/codeinfrarpc.LanguageRuntime/Run',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.RunRequest,
    responseType: codeinfra_language_pb.RunResponse,
    requestSerialize: serialize_codeinfrarpc_RunRequest,
    requestDeserialize: deserialize_codeinfrarpc_RunRequest,
    responseSerialize: serialize_codeinfrarpc_RunResponse,
    responseDeserialize: deserialize_codeinfrarpc_RunResponse,
  },
  // GetPluginInfo returns generic information about this plugin, like its version.
getPluginInfo: {
    path: '/codeinfrarpc.LanguageRuntime/GetPluginInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: codeinfra_plugin_pb.PluginInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_codeinfrarpc_PluginInfo,
    responseDeserialize: deserialize_codeinfrarpc_PluginInfo,
  },
  // InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects.
installDependencies: {
    path: '/codeinfrarpc.LanguageRuntime/InstallDependencies',
    requestStream: false,
    responseStream: true,
    requestType: codeinfra_language_pb.InstallDependenciesRequest,
    responseType: codeinfra_language_pb.InstallDependenciesResponse,
    requestSerialize: serialize_codeinfrarpc_InstallDependenciesRequest,
    requestDeserialize: deserialize_codeinfrarpc_InstallDependenciesRequest,
    responseSerialize: serialize_codeinfrarpc_InstallDependenciesResponse,
    responseDeserialize: deserialize_codeinfrarpc_InstallDependenciesResponse,
  },
  // RuntimeOptionsPrompts returns a list of additional prompts to ask during `codeinfra new`.
runtimeOptionsPrompts: {
    path: '/codeinfrarpc.LanguageRuntime/RuntimeOptionsPrompts',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.RuntimeOptionsRequest,
    responseType: codeinfra_language_pb.RuntimeOptionsResponse,
    requestSerialize: serialize_codeinfrarpc_RuntimeOptionsRequest,
    requestDeserialize: deserialize_codeinfrarpc_RuntimeOptionsRequest,
    responseSerialize: serialize_codeinfrarpc_RuntimeOptionsResponse,
    responseDeserialize: deserialize_codeinfrarpc_RuntimeOptionsResponse,
  },
  // About returns information about the runtime for this language.
about: {
    path: '/codeinfrarpc.LanguageRuntime/About',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.AboutRequest,
    responseType: codeinfra_language_pb.AboutResponse,
    requestSerialize: serialize_codeinfrarpc_AboutRequest,
    requestDeserialize: deserialize_codeinfrarpc_AboutRequest,
    responseSerialize: serialize_codeinfrarpc_AboutResponse,
    responseDeserialize: deserialize_codeinfrarpc_AboutResponse,
  },
  // GetProgramDependencies returns the set of dependencies required by the program.
getProgramDependencies: {
    path: '/codeinfrarpc.LanguageRuntime/GetProgramDependencies',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GetProgramDependenciesRequest,
    responseType: codeinfra_language_pb.GetProgramDependenciesResponse,
    requestSerialize: serialize_codeinfrarpc_GetProgramDependenciesRequest,
    requestDeserialize: deserialize_codeinfrarpc_GetProgramDependenciesRequest,
    responseSerialize: serialize_codeinfrarpc_GetProgramDependenciesResponse,
    responseDeserialize: deserialize_codeinfrarpc_GetProgramDependenciesResponse,
  },
  // RunPlugin executes a plugin program and returns its result asynchronously.
runPlugin: {
    path: '/codeinfrarpc.LanguageRuntime/RunPlugin',
    requestStream: false,
    responseStream: true,
    requestType: codeinfra_language_pb.RunPluginRequest,
    responseType: codeinfra_language_pb.RunPluginResponse,
    requestSerialize: serialize_codeinfrarpc_RunPluginRequest,
    requestDeserialize: deserialize_codeinfrarpc_RunPluginRequest,
    responseSerialize: serialize_codeinfrarpc_RunPluginResponse,
    responseDeserialize: deserialize_codeinfrarpc_RunPluginResponse,
  },
  // GenerateProgram generates a given PCL program into a program for this language.
generateProgram: {
    path: '/codeinfrarpc.LanguageRuntime/GenerateProgram',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GenerateProgramRequest,
    responseType: codeinfra_language_pb.GenerateProgramResponse,
    requestSerialize: serialize_codeinfrarpc_GenerateProgramRequest,
    requestDeserialize: deserialize_codeinfrarpc_GenerateProgramRequest,
    responseSerialize: serialize_codeinfrarpc_GenerateProgramResponse,
    responseDeserialize: deserialize_codeinfrarpc_GenerateProgramResponse,
  },
  // GenerateProject generates a given PCL program into a project for this language.
generateProject: {
    path: '/codeinfrarpc.LanguageRuntime/GenerateProject',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GenerateProjectRequest,
    responseType: codeinfra_language_pb.GenerateProjectResponse,
    requestSerialize: serialize_codeinfrarpc_GenerateProjectRequest,
    requestDeserialize: deserialize_codeinfrarpc_GenerateProjectRequest,
    responseSerialize: serialize_codeinfrarpc_GenerateProjectResponse,
    responseDeserialize: deserialize_codeinfrarpc_GenerateProjectResponse,
  },
  // GeneratePackage generates a given codeinfra package into a package for this language.
generatePackage: {
    path: '/codeinfrarpc.LanguageRuntime/GeneratePackage',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.GeneratePackageRequest,
    responseType: codeinfra_language_pb.GeneratePackageResponse,
    requestSerialize: serialize_codeinfrarpc_GeneratePackageRequest,
    requestDeserialize: deserialize_codeinfrarpc_GeneratePackageRequest,
    responseSerialize: serialize_codeinfrarpc_GeneratePackageResponse,
    responseDeserialize: deserialize_codeinfrarpc_GeneratePackageResponse,
  },
  // Pack packs a package into a language specific artifact.
pack: {
    path: '/codeinfrarpc.LanguageRuntime/Pack',
    requestStream: false,
    responseStream: false,
    requestType: codeinfra_language_pb.PackRequest,
    responseType: codeinfra_language_pb.PackResponse,
    requestSerialize: serialize_codeinfrarpc_PackRequest,
    requestDeserialize: deserialize_codeinfrarpc_PackRequest,
    responseSerialize: serialize_codeinfrarpc_PackResponse,
    responseDeserialize: deserialize_codeinfrarpc_PackResponse,
  },
};

exports.LanguageRuntimeClient = grpc.makeGenericClientConstructor(LanguageRuntimeService);
