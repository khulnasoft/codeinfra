// package: codeinfrarpc.testing
// file: codeinfra/testing/language.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as codeinfra_testing_language_pb from "../testing/language_pb";

interface ILanguageTestService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getLanguageTests: ILanguageTestService_IGetLanguageTests;
    prepareLanguageTests: ILanguageTestService_IPrepareLanguageTests;
    runLanguageTest: ILanguageTestService_IRunLanguageTest;
}

interface ILanguageTestService_IGetLanguageTests extends grpc.MethodDefinition<codeinfra_testing_language_pb.GetLanguageTestsRequest, codeinfra_testing_language_pb.GetLanguageTestsResponse> {
    path: "/codeinfrarpc.testing.LanguageTest/GetLanguageTests";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_testing_language_pb.GetLanguageTestsRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_testing_language_pb.GetLanguageTestsRequest>;
    responseSerialize: grpc.serialize<codeinfra_testing_language_pb.GetLanguageTestsResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_testing_language_pb.GetLanguageTestsResponse>;
}
interface ILanguageTestService_IPrepareLanguageTests extends grpc.MethodDefinition<codeinfra_testing_language_pb.PrepareLanguageTestsRequest, codeinfra_testing_language_pb.PrepareLanguageTestsResponse> {
    path: "/codeinfrarpc.testing.LanguageTest/PrepareLanguageTests";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_testing_language_pb.PrepareLanguageTestsRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_testing_language_pb.PrepareLanguageTestsRequest>;
    responseSerialize: grpc.serialize<codeinfra_testing_language_pb.PrepareLanguageTestsResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_testing_language_pb.PrepareLanguageTestsResponse>;
}
interface ILanguageTestService_IRunLanguageTest extends grpc.MethodDefinition<codeinfra_testing_language_pb.RunLanguageTestRequest, codeinfra_testing_language_pb.RunLanguageTestResponse> {
    path: "/codeinfrarpc.testing.LanguageTest/RunLanguageTest";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<codeinfra_testing_language_pb.RunLanguageTestRequest>;
    requestDeserialize: grpc.deserialize<codeinfra_testing_language_pb.RunLanguageTestRequest>;
    responseSerialize: grpc.serialize<codeinfra_testing_language_pb.RunLanguageTestResponse>;
    responseDeserialize: grpc.deserialize<codeinfra_testing_language_pb.RunLanguageTestResponse>;
}

export const LanguageTestService: ILanguageTestService;

export interface ILanguageTestServer extends grpc.UntypedServiceImplementation {
    getLanguageTests: grpc.handleUnaryCall<codeinfra_testing_language_pb.GetLanguageTestsRequest, codeinfra_testing_language_pb.GetLanguageTestsResponse>;
    prepareLanguageTests: grpc.handleUnaryCall<codeinfra_testing_language_pb.PrepareLanguageTestsRequest, codeinfra_testing_language_pb.PrepareLanguageTestsResponse>;
    runLanguageTest: grpc.handleUnaryCall<codeinfra_testing_language_pb.RunLanguageTestRequest, codeinfra_testing_language_pb.RunLanguageTestResponse>;
}

export interface ILanguageTestClient {
    getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
    runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
    runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
}

export class LanguageTestClient extends grpc.Client implements ILanguageTestClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public getLanguageTests(request: codeinfra_testing_language_pb.GetLanguageTestsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.GetLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public prepareLanguageTests(request: codeinfra_testing_language_pb.PrepareLanguageTestsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.PrepareLanguageTestsResponse) => void): grpc.ClientUnaryCall;
    public runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
    public runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
    public runLanguageTest(request: codeinfra_testing_language_pb.RunLanguageTestRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: codeinfra_testing_language_pb.RunLanguageTestResponse) => void): grpc.ClientUnaryCall;
}
