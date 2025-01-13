// package: codeinfrarpc
// file: codeinfra/analyzer.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as codeinfra_analyzer_pb from "./analyzer_pb";
import * as codeinfra_plugin_pb from "./plugin_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

interface IAnalyzerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    analyze: IAnalyzerService_IAnalyze;
    analyzeStack: IAnalyzerService_IAnalyzeStack;
    remediate: IAnalyzerService_IRemediate;
    getAnalyzerInfo: IAnalyzerService_IGetAnalyzerInfo;
    getPluginInfo: IAnalyzerService_IGetPluginInfo;
    configure: IAnalyzerService_IConfigure;
}

interface IAnalyzerService_IAnalyze extends grpc.MethodDefinition<codeinfra_analyzer_pb.AnalyzeRequest, codeinfra_analyzer_pb.AnalyzeResponse> {
    path: "/codeinfrarpc.Analyzer/Analyze";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzeRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzeRequest>;
    responseSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzeResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzeResponse>;
}
interface IAnalyzerService_IAnalyzeStack extends grpc.MethodDefinition<codeinfra_analyzer_pb.AnalyzeStackRequest, codeinfra_analyzer_pb.AnalyzeResponse> {
    path: "/codeinfrarpc.Analyzer/AnalyzeStack";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzeStackRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzeStackRequest>;
    responseSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzeResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzeResponse>;
}
interface IAnalyzerService_IRemediate extends grpc.MethodDefinition<codeinfra_analyzer_pb.AnalyzeRequest, codeinfra_analyzer_pb.RemediateResponse> {
    path: "/codeinfrarpc.Analyzer/Remediate";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzeRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzeRequest>;
    responseSerialize: grpc.serialize<codeinfra_analyzer_pb.RemediateResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_analyzer_pb.RemediateResponse>;
}
interface IAnalyzerService_IGetAnalyzerInfo extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, codeinfra_analyzer_pb.AnalyzerInfo> {
    path: "/codeinfrarpc.Analyzer/GetAnalyzerInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<codeinfra_analyzer_pb.AnalyzerInfo>;
    responseDeserialize: grpc.deserialize<codeinfra_analyzer_pb.AnalyzerInfo>;
}
interface IAnalyzerService_IGetPluginInfo extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, codeinfra_plugin_pb.PluginInfo> {
    path: "/codeinfrarpc.Analyzer/GetPluginInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<codeinfra_plugin_pb.PluginInfo>;
    responseDeserialize: grpc.deserialize<codeinfra_plugin_pb.PluginInfo>;
}
interface IAnalyzerService_IConfigure extends grpc.MethodDefinition<codeinfra_analyzer_pb.ConfigureAnalyzerRequest, google_protobuf_empty_pb.Empty> {
    path: "/codeinfrarpc.Analyzer/Configure";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_analyzer_pb.ConfigureAnalyzerRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_analyzer_pb.ConfigureAnalyzerRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const AnalyzerService: IAnalyzerService;

export interface IAnalyzerServer extends grpc.UntypedServiceImplementation {
    analyze: grpc.handleUnaryCall<codeinfra_analyzer_pb.AnalyzeRequest, codeinfra_analyzer_pb.AnalyzeResponse>;
    analyzeStack: grpc.handleUnaryCall<codeinfra_analyzer_pb.AnalyzeStackRequest, codeinfra_analyzer_pb.AnalyzeResponse>;
    remediate: grpc.handleUnaryCall<codeinfra_analyzer_pb.AnalyzeRequest, codeinfra_analyzer_pb.RemediateResponse>;
    getAnalyzerInfo: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, codeinfra_analyzer_pb.AnalyzerInfo>;
    getPluginInfo: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, codeinfra_plugin_pb.PluginInfo>;
    configure: grpc.handleUnaryCall<codeinfra_analyzer_pb.ConfigureAnalyzerRequest, google_protobuf_empty_pb.Empty>;
}

export interface IAnalyzerClient {
    analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class AnalyzerClient extends grpc.Client implements IAnalyzerClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public analyze(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public analyzeStack(request: codeinfra_analyzer_pb.AnalyzeStackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzeResponse) => void): grpc.ClientUnaryCall;
    public remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    public remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    public remediate(request: codeinfra_analyzer_pb.AnalyzeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.RemediateResponse) => void): grpc.ClientUnaryCall;
    public getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    public getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    public getAnalyzerInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_analyzer_pb.AnalyzerInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public configure(request: codeinfra_analyzer_pb.ConfigureAnalyzerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}
