import { AccountClient } from "../grpc/clients/AccountClient";

export const accountResolvers = {
    Query: {
        getAccount: async (
            _: any,
            { id }: { id: number }
        ) => {
            console.log("Getting account with id: ", id);
            return new AccountClient().getAccount(id);
        },
    },
};
