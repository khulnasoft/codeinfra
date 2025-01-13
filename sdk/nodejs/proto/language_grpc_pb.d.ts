// package: codeinfrarpc
// file: codeinfra/language.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as codeinfra_language_pb from "./language_pb";
import * as codeinfra_codegen_hcl_pb from "./codegen/hcl_pb";
import * as codeinfra_plugin_pb from "./plugin_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

interface ILanguageRuntimeService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    handshake: ILanguageRuntimeService_IHandshake;
    getRequiredPlugins: ILanguageRuntimeService_IGetRequiredPlugins;
    getRequiredPackages: ILanguageRuntimeService_IGetRequiredPackages;
    run: ILanguageRuntimeService_IRun;
    getPluginInfo: ILanguageRuntimeService_IGetPluginInfo;
    installDependencies: ILanguageRuntimeService_IInstallDependencies;
    runtimeOptionsPrompts: ILanguageRuntimeService_IRuntimeOptionsPrompts;
    about: ILanguageRuntimeService_IAbout;
    getProgramDependencies: ILanguageRuntimeService_IGetProgramDependencies;
    runPlugin: ILanguageRuntimeService_IRunPlugin;
    generateProgram: ILanguageRuntimeService_IGenerateProgram;
    generateProject: ILanguageRuntimeService_IGenerateProject;
    generatePackage: ILanguageRuntimeService_IGeneratePackage;
    pack: ILanguageRuntimeService_IPack;
}

interface ILanguageRuntimeService_IHandshake extends grpc.MethodDefinition<codeinfra_language_pb.LanguageHandshakeRequest, codeinfra_language_pb.LanguageHandshakeResponse> {
    path: "/codeinfrarpc.LanguageRuntime/Handshake";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.LanguageHandshakeRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.LanguageHandshakeRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.LanguageHandshakeResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.LanguageHandshakeResponse>;
}
interface ILanguageRuntimeService_IGetRequiredPlugins extends grpc.MethodDefinition<codeinfra_language_pb.GetRequiredPluginsRequest, codeinfra_language_pb.GetRequiredPluginsResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GetRequiredPlugins";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GetRequiredPluginsRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GetRequiredPluginsRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GetRequiredPluginsResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GetRequiredPluginsResponse>;
}
interface ILanguageRuntimeService_IGetRequiredPackages extends grpc.MethodDefinition<codeinfra_language_pb.GetRequiredPackagesRequest, codeinfra_language_pb.GetRequiredPackagesResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GetRequiredPackages";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GetRequiredPackagesRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GetRequiredPackagesRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GetRequiredPackagesResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GetRequiredPackagesResponse>;
}
interface ILanguageRuntimeService_IRun extends grpc.MethodDefinition<codeinfra_language_pb.RunRequest, codeinfra_language_pb.RunResponse> {
    path: "/codeinfrarpc.LanguageRuntime/Run";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.RunRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.RunRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.RunResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.RunResponse>;
}
interface ILanguageRuntimeService_IGetPluginInfo extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, codeinfra_plugin_pb.PluginInfo> {
    path: "/codeinfrarpc.LanguageRuntime/GetPluginInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<codeinfra_plugin_pb.PluginInfo>;
    responseDeserialize: grpc.deserialize<codeinfra_plugin_pb.PluginInfo>;
}
interface ILanguageRuntimeService_IInstallDependencies extends grpc.MethodDefinition<codeinfra_language_pb.InstallDependenciesRequest, codeinfra_language_pb.InstallDependenciesResponse> {
    path: "/codeinfrarpc.LanguageRuntime/InstallDependencies";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<codeinfra_language_pb.InstallDependenciesRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.InstallDependenciesRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.InstallDependenciesResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.InstallDependenciesResponse>;
}
interface ILanguageRuntimeService_IRuntimeOptionsPrompts extends grpc.MethodDefinition<codeinfra_language_pb.RuntimeOptionsRequest, codeinfra_language_pb.RuntimeOptionsResponse> {
    path: "/codeinfrarpc.LanguageRuntime/RuntimeOptionsPrompts";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.RuntimeOptionsRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.RuntimeOptionsRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.RuntimeOptionsResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.RuntimeOptionsResponse>;
}
interface ILanguageRuntimeService_IAbout extends grpc.MethodDefinition<codeinfra_language_pb.AboutRequest, codeinfra_language_pb.AboutResponse> {
    path: "/codeinfrarpc.LanguageRuntime/About";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.AboutRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.AboutRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.AboutResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.AboutResponse>;
}
interface ILanguageRuntimeService_IGetProgramDependencies extends grpc.MethodDefinition<codeinfra_language_pb.GetProgramDependenciesRequest, codeinfra_language_pb.GetProgramDependenciesResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GetProgramDependencies";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GetProgramDependenciesRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GetProgramDependenciesRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GetProgramDependenciesResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GetProgramDependenciesResponse>;
}
interface ILanguageRuntimeService_IRunPlugin extends grpc.MethodDefinition<codeinfra_language_pb.RunPluginRequest, codeinfra_language_pb.RunPluginResponse> {
    path: "/codeinfrarpc.LanguageRuntime/RunPlugin";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<codeinfra_language_pb.RunPluginRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.RunPluginRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.RunPluginResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.RunPluginResponse>;
}
interface ILanguageRuntimeService_IGenerateProgram extends grpc.MethodDefinition<codeinfra_language_pb.GenerateProgramRequest, codeinfra_language_pb.GenerateProgramResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GenerateProgram";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GenerateProgramRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GenerateProgramRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GenerateProgramResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GenerateProgramResponse>;
}
interface ILanguageRuntimeService_IGenerateProject extends grpc.MethodDefinition<codeinfra_language_pb.GenerateProjectRequest, codeinfra_language_pb.GenerateProjectResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GenerateProject";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GenerateProjectRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GenerateProjectRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GenerateProjectResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GenerateProjectResponse>;
}
interface ILanguageRuntimeService_IGeneratePackage extends grpc.MethodDefinition<codeinfra_language_pb.GeneratePackageRequest, codeinfra_language_pb.GeneratePackageResponse> {
    path: "/codeinfrarpc.LanguageRuntime/GeneratePackage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.GeneratePackageRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.GeneratePackageRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.GeneratePackageResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.GeneratePackageResponse>;
}
interface ILanguageRuntimeService_IPack extends grpc.MethodDefinition<codeinfra_language_pb.PackRequest, codeinfra_language_pb.PackResponse> {
    path: "/codeinfrarpc.LanguageRuntime/Pack";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_language_pb.PackRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_language_pb.PackRequest>;
    responseSerialize: grpc.serialize<codeinfra_language_pb.PackResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_language_pb.PackResponse>;
}

export const LanguageRuntimeService: ILanguageRuntimeService;

export interface ILanguageRuntimeServer extends grpc.UntypedServiceImplementation {
    handshake: grpc.handleUnaryCall<codeinfra_language_pb.LanguageHandshakeRequest, codeinfra_language_pb.LanguageHandshakeResponse>;
    getRequiredPlugins: grpc.handleUnaryCall<codeinfra_language_pb.GetRequiredPluginsRequest, codeinfra_language_pb.GetRequiredPluginsResponse>;
    getRequiredPackages: grpc.handleUnaryCall<codeinfra_language_pb.GetRequiredPackagesRequest, codeinfra_language_pb.GetRequiredPackagesResponse>;
    run: grpc.handleUnaryCall<codeinfra_language_pb.RunRequest, codeinfra_language_pb.RunResponse>;
    getPluginInfo: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, codeinfra_plugin_pb.PluginInfo>;
    installDependencies: grpc.handleServerStreamingCall<codeinfra_language_pb.InstallDependenciesRequest, codeinfra_language_pb.InstallDependenciesResponse>;
    runtimeOptionsPrompts: grpc.handleUnaryCall<codeinfra_language_pb.RuntimeOptionsRequest, codeinfra_language_pb.RuntimeOptionsResponse>;
    about: grpc.handleUnaryCall<codeinfra_language_pb.AboutRequest, codeinfra_language_pb.AboutResponse>;
    getProgramDependencies: grpc.handleUnaryCall<codeinfra_language_pb.GetProgramDependenciesRequest, codeinfra_language_pb.GetProgramDependenciesResponse>;
    runPlugin: grpc.handleServerStreamingCall<codeinfra_language_pb.RunPluginRequest, codeinfra_language_pb.RunPluginResponse>;
    generateProgram: grpc.handleUnaryCall<codeinfra_language_pb.GenerateProgramRequest, codeinfra_language_pb.GenerateProgramResponse>;
    generateProject: grpc.handleUnaryCall<codeinfra_language_pb.GenerateProjectRequest, codeinfra_language_pb.GenerateProjectResponse>;
    generatePackage: grpc.handleUnaryCall<codeinfra_language_pb.GeneratePackageRequest, codeinfra_language_pb.GeneratePackageResponse>;
    pack: grpc.handleUnaryCall<codeinfra_language_pb.PackRequest, codeinfra_language_pb.PackResponse>;
}

export interface ILanguageRuntimeClient {
    handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    run(request: codeinfra_language_pb.RunRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    run(request: codeinfra_language_pb.RunRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    run(request: codeinfra_language_pb.RunRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    installDependencies(request: codeinfra_language_pb.InstallDependenciesRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.InstallDependenciesResponse>;
    installDependencies(request: codeinfra_language_pb.InstallDependenciesRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.InstallDependenciesResponse>;
    runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    about(request: codeinfra_language_pb.AboutRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    about(request: codeinfra_language_pb.AboutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    about(request: codeinfra_language_pb.AboutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    runPlugin(request: codeinfra_language_pb.RunPluginRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.RunPluginResponse>;
    runPlugin(request: codeinfra_language_pb.RunPluginRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.RunPluginResponse>;
    generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: codeinfra_language_pb.GenerateProjectRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: codeinfra_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: codeinfra_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    pack(request: codeinfra_language_pb.PackRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    pack(request: codeinfra_language_pb.PackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    pack(request: codeinfra_language_pb.PackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
}

export class LanguageRuntimeClient extends grpc.Client implements ILanguageRuntimeClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    public handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    public handshake(request: codeinfra_language_pb.LanguageHandshakeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.LanguageHandshakeResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPlugins(request: codeinfra_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPackages(request: codeinfra_language_pb.GetRequiredPackagesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetRequiredPackagesResponse) => void): grpc.ClientUnaryCall;
    public run(request: codeinfra_language_pb.RunRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public run(request: codeinfra_language_pb.RunRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public run(request: codeinfra_language_pb.RunRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public installDependencies(request: codeinfra_language_pb.InstallDependenciesRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.InstallDependenciesResponse>;
    public installDependencies(request: codeinfra_language_pb.InstallDependenciesRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.InstallDependenciesResponse>;
    public runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public runtimeOptionsPrompts(request: codeinfra_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public about(request: codeinfra_language_pb.AboutRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public about(request: codeinfra_language_pb.AboutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public about(request: codeinfra_language_pb.AboutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: codeinfra_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public runPlugin(request: codeinfra_language_pb.RunPluginRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.RunPluginResponse>;
    public runPlugin(request: codeinfra_language_pb.RunPluginRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<codeinfra_language_pb.RunPluginResponse>;
    public generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProgram(request: codeinfra_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: codeinfra_language_pb.GenerateProjectRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: codeinfra_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: codeinfra_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: codeinfra_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public pack(request: codeinfra_language_pb.PackRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    public pack(request: codeinfra_language_pb.PackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    public pack(request: codeinfra_language_pb.PackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
}
