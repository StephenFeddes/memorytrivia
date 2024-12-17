import { accountResolvers } from "./account";

export const resolvers = {
    Query: {
        ...accountResolvers.Query,
    },
    Mutation: {},
};