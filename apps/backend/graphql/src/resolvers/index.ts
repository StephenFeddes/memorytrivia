import { triviaResolvers } from "./trivia";

export const resolvers = {
    Query: {
        ...triviaResolvers.Query,
    },
    Mutation: {
        
    },
}