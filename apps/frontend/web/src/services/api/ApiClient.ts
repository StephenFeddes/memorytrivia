const API_URL = import.meta.env.VITE_API_URL;

export class ApiClient {
    private baseUrl: string;
    private headers: Record<string, string>;

    constructor(
        baseUrl: string = API_URL,
        headers: Record<string, string> = {
            "Content-Type": "application/json",
        }
    ) {
        this.baseUrl = baseUrl;
        this.headers = headers;
    }

    private async request<T>(
        endpoint: string,
        options: RequestInit = {}
    ): Promise<T> {
        const url = `${this.baseUrl}${endpoint}`;

        try {
            const response = await fetch(url, { ...options, ...this.headers });
            if (!response.ok) {
                const errorBody = await response.text();
                throw new Error(`Error: ${response.status}, ${errorBody}`);
            }
            return response.json() as Promise<T>;
        } catch (error) {
            console.error("API Request Error:", error);
            throw error;
        }
    }

    public get<T>(endpoint: string, params?: Record<string, any>): Promise<T> {
        const query = params
            ? `?${new URLSearchParams(params as any).toString()}`
            : "";
        return this.request<T>(`${endpoint}${query}`, {
            method: "GET",
        });
    }

    public post<T>(
        endpoint: string,
        body?: any,
        headers?: HeadersInit
    ): Promise<T> {
        return this.request<T>(endpoint, {
            method: "POST",
            body: JSON.stringify(body),
            headers: { "Content-Type": "application/json", ...headers },
        });
    }

    public put<T>(
        endpoint: string,
        body?: any,
        headers?: HeadersInit
    ): Promise<T> {
        return this.request<T>(endpoint, {
            method: "PUT",
            body: JSON.stringify(body),
            headers: { "Content-Type": "application/json", ...headers },
        });
    }

    public delete<T>(endpoint: string, headers?: HeadersInit): Promise<T> {
        return this.request<T>(endpoint, {
            method: "DELETE",
            headers,
        });
    }
}
