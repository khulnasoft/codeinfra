// package: codeinfrarpc
// file: codeinfra/callback.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as codeinfra_callback_pb from "./callback_pb";

interface ICallbacksService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    invoke: ICallbacksService_IInvoke;
}

interface ICallbacksService_IInvoke extends grpc.MethodDefinition<codeinfra_callback_pb.CallbackInvokeRequest, codeinfra_callback_pb.CallbackInvokeResponse> {
    path: "/codeinfrarpc.Callbacks/Invoke";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_callback_pb.CallbackInvokeRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_callback_pb.CallbackInvokeRequest>;
    responseSerialize: grpc.serialize<codeinfra_callback_pb.CallbackInvokeResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_callback_pb.CallbackInvokeResponse>;
}

export const CallbacksService: ICallbacksService;

export interface ICallbacksServer extends grpc.UntypedServiceImplementation {
    invoke: grpc.handleUnaryCall<codeinfra_callback_pb.CallbackInvokeRequest, codeinfra_callback_pb.CallbackInvokeResponse>;
}

export interface ICallbacksClient {
    invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
}

export class CallbacksClient extends grpc.Client implements ICallbacksClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: codeinfra_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
}
