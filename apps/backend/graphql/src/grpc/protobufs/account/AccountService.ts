// Original file: ../protobuf-schemas/account.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Account as _account_Account, Account__Output as _account_Account__Output } from '../account/Account';
import type { CreateAccountRequest as _account_CreateAccountRequest, CreateAccountRequest__Output as _account_CreateAccountRequest__Output } from '../account/CreateAccountRequest';
import type { GetAccountRequest as _account_GetAccountRequest, GetAccountRequest__Output as _account_GetAccountRequest__Output } from '../account/GetAccountRequest';

export interface AccountServiceClient extends grpc.Client {
  CreateAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  
  GetAccount(argument: _account_GetAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  GetAccount(argument: _account_GetAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  GetAccount(argument: _account_GetAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  GetAccount(argument: _account_GetAccountRequest, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  getAccount(argument: _account_GetAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  getAccount(argument: _account_GetAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  getAccount(argument: _account_GetAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  getAccount(argument: _account_GetAccountRequest, callback: grpc.requestCallback<_account_Account__Output>): grpc.ClientUnaryCall;
  
}

export interface AccountServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateAccount: grpc.handleUnaryCall<_account_CreateAccountRequest__Output, _account_Account>;
  
  GetAccount: grpc.handleUnaryCall<_account_GetAccountRequest__Output, _account_Account>;
  
}

export interface AccountServiceDefinition extends grpc.ServiceDefinition {
  CreateAccount: MethodDefinition<_account_CreateAccountRequest, _account_Account, _account_CreateAccountRequest__Output, _account_Account__Output>
  GetAccount: MethodDefinition<_account_GetAccountRequest, _account_Account, _account_GetAccountRequest__Output, _account_Account__Output>
}
