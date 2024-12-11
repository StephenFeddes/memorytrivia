import path from "path";
import * as grpc from "@grpc/grpc-js";
import * as protoloader from "@grpc/proto-loader";
import { ProtoGrpcType } from "../protobufs/trivia";
import { TriviaServiceClient } from "../protobufs/trivia/TriviaService";
import { GetQuestionRequest } from "../protobufs/trivia/GetQuestionRequest";
import { Question } from "../protobufs/trivia/Question";
import * as dotenv from "dotenv";
dotenv.config();

class TriviaClient {
    private readonly PORT: number = process.env.TRIVIA_PORT
        ? parseInt(process.env.TRIVIA_PORT, 10)
        : 3002;

    private readonly PROTO_FILE: string = path.join(
        process.cwd(),
        "protobuf-schemas/trivia.proto"
    );

    private client: TriviaServiceClient;

    constructor() {
        // Load and resolve the .proto file
        const packageDefinition = protoloader.loadSync(
            path.resolve(__dirname, this.PROTO_FILE)
        );
        const packageObject = grpc.loadPackageDefinition(
            packageDefinition
        ) as unknown as ProtoGrpcType;

        // Create the gRPC client
       this.client = new packageObject.trivia.TriviaService(
           `trivia:${this.PORT}`, // Note: `trivia` is the name of the service defined in `docker-compose.yml`
           grpc.credentials.createInsecure()
       );
    }

    /**
     * Get a trivia question from the server.
     * @param difficulty - The difficulty level of the question.
     * @param category - The category of the question.
     * @returns A promise resolving to the question data.
     */
    public getQuestion(
        difficulty: string,
        category: string
    ): Promise<Question> {
        const request: GetQuestionRequest = { difficulty, category };

        return new Promise((resolve, reject) => {
            this.client.GetQuestion(request, (err, response) => {
                if (err) {
                    reject(err);
                } else if (response) {
                    resolve(response);
                } else {
                    reject(new Error("No response received from the server."));
                }
            });
        });
    }
}

export default TriviaClient;
