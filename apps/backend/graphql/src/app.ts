import path from "path";
import { ApolloServer } from "@apollo/server";
import { loadFilesSync } from "@graphql-tools/load-files";
import { mergeTypeDefs } from "@graphql-tools/merge";
import { resolvers } from "./resolvers";
import { createServer } from "http";
import { expressMiddleware } from "@apollo/server/express4";
import express from "express";
import * as dotenv from "dotenv";
dotenv.config();

const loadedSchemas = loadFilesSync(
    path.join(__dirname, "schemas/**/*.graphql")
);

// Combine schemas
export const graphqlSchema = mergeTypeDefs(loadedSchemas);

const app = express();
const httpServer = createServer(app);

const server = new ApolloServer({
    typeDefs: graphqlSchema,
    resolvers,
});

const PORT = process.env.PORT || "3000";

// Start the server and apply middleware
httpServer.listen({ port: PORT }, async () => {
    await server.start();

    // Add JSON parsing middleware
    app.use(express.json());

    // Apply the Apollo Server middleware
    app.use("/graphql", expressMiddleware(server));

    console.log(`Server ready at http://localhost:${PORT}/graphql`);
});
