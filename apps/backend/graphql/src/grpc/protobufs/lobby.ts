import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { LobbyServiceClient as _lobby_LobbyServiceClient, LobbyServiceDefinition as _lobby_LobbyServiceDefinition } from './lobby/LobbyService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  lobby: {
    CreateLobbyRequest: MessageTypeDefinition
    CreateLobbyResponse: MessageTypeDefinition
    DeleteLobbyRequest: MessageTypeDefinition
    DeleteLobbyResponse: MessageTypeDefinition
    GetLobbyRequest: MessageTypeDefinition
    Lobby: MessageTypeDefinition
    LobbyService: SubtypeConstructor<typeof grpc.Client, _lobby_LobbyServiceClient> & { service: _lobby_LobbyServiceDefinition }
    UpdateLobbyRequest: MessageTypeDefinition
    UpdateLobbyResponse: MessageTypeDefinition
  }
}

