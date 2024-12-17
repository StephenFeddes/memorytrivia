import { User } from "../../../types/user";
import { ApiClient } from "../ApiClient";

export class UserApi {

    apiClient: ApiClient;

    constructor() {
        this.apiClient = new ApiClient();
    }

    public async signInUser(data: {
        username: string;
        email: string;
        password: string;
    }): Promise<{ jwtToken: string; user: User }> {
        const response = await this.apiClient.post<{
            jwtToken: string;
            user: User;
        }>("/users/tokens", data);
        return response
    }

    public async createUser(data: {
        username: string;
        email: string;
        password: string;
    }): Promise<{ jwtToken: string; user: User }> {
        const response = await this.apiClient.post<{
            jwtToken: string;
            user: User;
        }>("/users", data);
        return response;
    }
}
