import path from "path";
import * as grpc from "@grpc/grpc-js";
import * as protoloader from "@grpc/proto-loader";
import { ProtoGrpcType } from "../protobufs/account";
import { Question } from "../protobufs/trivia/Question";
import * as dotenv from "dotenv";
import { AccountServiceClient } from "../protobufs/account/AccountService";
import { GetAccountRequest } from "../protobufs/account/GetAccountRequest";
dotenv.config();

export class AccountClient {
    private readonly PORT: number = process.env.ACCOUNT_PORT
        ? parseInt(process.env.ACCOUNT_PORT, 10)
        : 50051;

    private readonly PROTO_FILE: string = path.join(
        process.cwd(),
        "protobuf-schemas/account.proto"
    );

    private client: AccountServiceClient;

    constructor() {
        // Load and resolve the .proto file
        const packageDefinition = protoloader.loadSync(
            path.resolve(__dirname, this.PROTO_FILE)
        );
        const packageObject = grpc.loadPackageDefinition(
            packageDefinition
        ) as unknown as ProtoGrpcType;

        // Create the gRPC client
        this.client = new packageObject.account.AccountService(
            `account:${this.PORT}`, // Note: `account` is the name of the service defined in `docker-compose.yml`
            grpc.credentials.createInsecure()
        );
    }

    /**
     * Get a trivia question from the server.
     * @param difficulty - The difficulty level of the question.
     * @param category - The category of the question.
     * @returns A promise resolving to the question data.
     */
    public getAccount(
        id: number
    ): Promise<Question> {
        const request: GetAccountRequest = { id: id};

        return new Promise((resolve, reject) => {
            this.client.GetAccount(request, (err, response) => {
                if (err) {
                    reject(err);
                } else if (response) {
                    console.log(response);
                    resolve(response);
                } else {
                    reject(new Error("No response received from the server."));
                }
            });
        });
    }
}
