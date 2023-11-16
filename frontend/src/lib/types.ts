export enum SearchState {
    Default,
    Finished,
    Searching,
    Error,
}

export interface SearchResult {
    urls: string[];
}
