import TriviaClient from "../grpc/clients/TriviaClient";

export const triviaResolvers = {
    Query: {
        getQuestion: async (
            _: any,
            { difficulty, category }: { difficulty: string; category: string }
        ) => {
            return new TriviaClient().getQuestion(difficulty, category);
        },
    },
};
