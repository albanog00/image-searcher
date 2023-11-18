import type { Result } from "./types";

// Create an interface that extends Response that takes Generic type T and Result
export interface ApiResponse<T, K = Result<T>> extends Response {
    json(): Promise<K>;
}