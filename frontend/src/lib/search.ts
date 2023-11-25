import { api } from "./api";
import type { SearchPhotoResult } from "./types";

import { searchImageQuery, searchPaginationQuery } from "./store";

export async function search() {
    const paginationParams = searchPaginationQuery.get();
    return await api<SearchPhotoResult>(
        `/search?q=${searchImageQuery.get()}&page=${paginationParams.page}&per_page=${paginationParams.perPage}`,
        {
            method: "GET",
        },
    );
}