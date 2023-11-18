export enum SearchState {
    Default,
    Finished,
    Searching,
    Error,
}

export interface Result<T> {
    message: string,
    status: number,
    result?: {
        data: T,
        count: number
    } | null
}

export interface SearchResult {
    urls: string[];
}
