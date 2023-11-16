export interface ApiResponse<T> extends Response {
    json(): Promise<T>;
}