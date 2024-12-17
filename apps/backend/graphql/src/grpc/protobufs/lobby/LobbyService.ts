// Original file: ../protobuf-schemas/lobby.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateLobbyRequest as _lobby_CreateLobbyRequest, CreateLobbyRequest__Output as _lobby_CreateLobbyRequest__Output } from '../lobby/CreateLobbyRequest';
import type { CreateLobbyResponse as _lobby_CreateLobbyResponse, CreateLobbyResponse__Output as _lobby_CreateLobbyResponse__Output } from '../lobby/CreateLobbyResponse';
import type { DeleteLobbyRequest as _lobby_DeleteLobbyRequest, DeleteLobbyRequest__Output as _lobby_DeleteLobbyRequest__Output } from '../lobby/DeleteLobbyRequest';
import type { DeleteLobbyResponse as _lobby_DeleteLobbyResponse, DeleteLobbyResponse__Output as _lobby_DeleteLobbyResponse__Output } from '../lobby/DeleteLobbyResponse';
import type { GetLobbyRequest as _lobby_GetLobbyRequest, GetLobbyRequest__Output as _lobby_GetLobbyRequest__Output } from '../lobby/GetLobbyRequest';
import type { Lobby as _lobby_Lobby, Lobby__Output as _lobby_Lobby__Output } from '../lobby/Lobby';
import type { UpdateLobbyRequest as _lobby_UpdateLobbyRequest, UpdateLobbyRequest__Output as _lobby_UpdateLobbyRequest__Output } from '../lobby/UpdateLobbyRequest';
import type { UpdateLobbyResponse as _lobby_UpdateLobbyResponse, UpdateLobbyResponse__Output as _lobby_UpdateLobbyResponse__Output } from '../lobby/UpdateLobbyResponse';

export interface LobbyServiceClient extends grpc.Client {
  CreateLobby(argument: _lobby_CreateLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  CreateLobby(argument: _lobby_CreateLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  CreateLobby(argument: _lobby_CreateLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  CreateLobby(argument: _lobby_CreateLobbyRequest, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  createLobby(argument: _lobby_CreateLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  createLobby(argument: _lobby_CreateLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  createLobby(argument: _lobby_CreateLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  createLobby(argument: _lobby_CreateLobbyRequest, callback: grpc.requestCallback<_lobby_CreateLobbyResponse__Output>): grpc.ClientUnaryCall;
  
  DeleteLobby(argument: _lobby_DeleteLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  DeleteLobby(argument: _lobby_DeleteLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  DeleteLobby(argument: _lobby_DeleteLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  DeleteLobby(argument: _lobby_DeleteLobbyRequest, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  deleteLobby(argument: _lobby_DeleteLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  deleteLobby(argument: _lobby_DeleteLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  deleteLobby(argument: _lobby_DeleteLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  deleteLobby(argument: _lobby_DeleteLobbyRequest, callback: grpc.requestCallback<_lobby_DeleteLobbyResponse__Output>): grpc.ClientUnaryCall;
  
  GetLobby(argument: _lobby_GetLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  GetLobby(argument: _lobby_GetLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  GetLobby(argument: _lobby_GetLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  GetLobby(argument: _lobby_GetLobbyRequest, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  getLobby(argument: _lobby_GetLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  getLobby(argument: _lobby_GetLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  getLobby(argument: _lobby_GetLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  getLobby(argument: _lobby_GetLobbyRequest, callback: grpc.requestCallback<_lobby_Lobby__Output>): grpc.ClientUnaryCall;
  
  UpdateLobby(argument: _lobby_UpdateLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  UpdateLobby(argument: _lobby_UpdateLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  UpdateLobby(argument: _lobby_UpdateLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  UpdateLobby(argument: _lobby_UpdateLobbyRequest, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  updateLobby(argument: _lobby_UpdateLobbyRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  updateLobby(argument: _lobby_UpdateLobbyRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  updateLobby(argument: _lobby_UpdateLobbyRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  updateLobby(argument: _lobby_UpdateLobbyRequest, callback: grpc.requestCallback<_lobby_UpdateLobbyResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface LobbyServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateLobby: grpc.handleUnaryCall<_lobby_CreateLobbyRequest__Output, _lobby_CreateLobbyResponse>;
  
  DeleteLobby: grpc.handleUnaryCall<_lobby_DeleteLobbyRequest__Output, _lobby_DeleteLobbyResponse>;
  
  GetLobby: grpc.handleUnaryCall<_lobby_GetLobbyRequest__Output, _lobby_Lobby>;
  
  UpdateLobby: grpc.handleUnaryCall<_lobby_UpdateLobbyRequest__Output, _lobby_UpdateLobbyResponse>;
  
}

export interface LobbyServiceDefinition extends grpc.ServiceDefinition {
  CreateLobby: MethodDefinition<_lobby_CreateLobbyRequest, _lobby_CreateLobbyResponse, _lobby_CreateLobbyRequest__Output, _lobby_CreateLobbyResponse__Output>
  DeleteLobby: MethodDefinition<_lobby_DeleteLobbyRequest, _lobby_DeleteLobbyResponse, _lobby_DeleteLobbyRequest__Output, _lobby_DeleteLobbyResponse__Output>
  GetLobby: MethodDefinition<_lobby_GetLobbyRequest, _lobby_Lobby, _lobby_GetLobbyRequest__Output, _lobby_Lobby__Output>
  UpdateLobby: MethodDefinition<_lobby_UpdateLobbyRequest, _lobby_UpdateLobbyResponse, _lobby_UpdateLobbyRequest__Output, _lobby_UpdateLobbyResponse__Output>
}
