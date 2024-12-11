import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { TriviaServiceClient as _trivia_TriviaServiceClient, TriviaServiceDefinition as _trivia_TriviaServiceDefinition } from './trivia/TriviaService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  trivia: {
    GetQuestionRequest: MessageTypeDefinition
    Question: MessageTypeDefinition
    TriviaService: SubtypeConstructor<typeof grpc.Client, _trivia_TriviaServiceClient> & { service: _trivia_TriviaServiceDefinition }
  }
}

