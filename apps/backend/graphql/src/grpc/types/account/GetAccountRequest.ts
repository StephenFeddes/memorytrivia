// Original file: ../protobufs/account.proto

import type { Long } from '@grpc/proto-loader';

export interface GetAccountRequest {
  'id'?: (number | string | Long);
}

export interface GetAccountRequest__Output {
  'id'?: (Long);
}
