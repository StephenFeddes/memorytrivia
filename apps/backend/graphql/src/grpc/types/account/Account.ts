// Original file: ../protobufs/account.proto

import type { Long } from '@grpc/proto-loader';

export interface Account {
  'id'?: (number | string | Long);
  'username'?: (string);
  'email'?: (string);
}

export interface Account__Output {
  'id'?: (Long);
  'username'?: (string);
  'email'?: (string);
}
