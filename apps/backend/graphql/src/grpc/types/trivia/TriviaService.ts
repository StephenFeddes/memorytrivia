// Original file: ../protobufs/trivia.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { GetQuestionRequest as _trivia_GetQuestionRequest, GetQuestionRequest__Output as _trivia_GetQuestionRequest__Output } from '../trivia/GetQuestionRequest';
import type { Question as _trivia_Question, Question__Output as _trivia_Question__Output } from '../trivia/Question';

export interface TriviaServiceClient extends grpc.Client {
  GetQuestion(argument: _trivia_GetQuestionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  GetQuestion(argument: _trivia_GetQuestionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  GetQuestion(argument: _trivia_GetQuestionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  GetQuestion(argument: _trivia_GetQuestionRequest, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  getQuestion(argument: _trivia_GetQuestionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  getQuestion(argument: _trivia_GetQuestionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  getQuestion(argument: _trivia_GetQuestionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  getQuestion(argument: _trivia_GetQuestionRequest, callback: grpc.requestCallback<_trivia_Question__Output>): grpc.ClientUnaryCall;
  
}

export interface TriviaServiceHandlers extends grpc.UntypedServiceImplementation {
  GetQuestion: grpc.handleUnaryCall<_trivia_GetQuestionRequest__Output, _trivia_Question>;
  
}

export interface TriviaServiceDefinition extends grpc.ServiceDefinition {
  GetQuestion: MethodDefinition<_trivia_GetQuestionRequest, _trivia_Question, _trivia_GetQuestionRequest__Output, _trivia_Question__Output>
}
