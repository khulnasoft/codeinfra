// package: codegen
// file: codeinfra/codegen/loader.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as codeinfra_codegen_loader_pb from "../codegen/loader_pb";

interface ILoaderService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getSchema: ILoaderService_IGetSchema;
}

interface ILoaderService_IGetSchema extends grpc.MethodDefinition<codeinfra_codegen_loader_pb.GetSchemaRequest, codeinfra_codegen_loader_pb.GetSchemaResponse> {
    path: "/codegen.Loader/GetSchema";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_codegen_loader_pb.GetSchemaRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_codegen_loader_pb.GetSchemaRequest>;
    responseSerialize: grpc.serialize<codeinfra_codegen_loader_pb.GetSchemaResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_codegen_loader_pb.GetSchemaResponse>;
}

export const LoaderService: ILoaderService;

export interface ILoaderServer extends grpc.UntypedServiceImplementation {
    getSchema: grpc.handleUnaryCall<codeinfra_codegen_loader_pb.GetSchemaRequest, codeinfra_codegen_loader_pb.GetSchemaResponse>;
}

export interface ILoaderClient {
    getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
}

export class LoaderClient extends grpc.Client implements ILoaderClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    public getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    public getSchema(request: codeinfra_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
}
