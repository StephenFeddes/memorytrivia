import React, { createContext, useContext } from "react";
import { UserApi } from "../services/api/endpoints/UserApi";
import { User } from "../types/user";

export interface ApiContextType {
    createUser: (data: { username: string; email: string, password: string }) => Promise<{
        jwtToken: string;
        user: User
    }>;
}

const ApiContext = createContext<ApiContextType | undefined>(undefined);

export const ApiProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const userApi = new UserApi();
    const createUser = async (data: { username: string; email: string, password: string }) => userApi.createUser(data);

    return (
        <ApiContext.Provider value={{ createUser }}>
            {children}
        </ApiContext.Provider>
    );
};

export const useApi = (): ApiContextType => {
    const context = useContext(ApiContext);
    if (!context) {
        throw new Error("useApi must be used within an ApiProvider");
    }
    return context;
};
