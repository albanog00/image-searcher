import { api } from "./api";
import { SearchState, type SearchPhotoResult } from "./types";

import {
  photos,
  searchImageQuery,
  searchPaginationQuery,
  searchState,
} from "./store";

export async function search(page?: number) {
  const paginationParams = searchPaginationQuery.get();
  try {
    const res = await api<SearchPhotoResult>(
      `/search?q=${searchImageQuery.get()}&page=${
        page ?? paginationParams.page
      }&per_page=${paginationParams.perPage}`,
      {
        method: "GET",
      },
    );

    searchState.set(SearchState.Searching);
    const json = await res.json();

    searchState.set(SearchState.Finished);
    photos.set(json.result?.data);
  } catch (error) {
    console.error(error);

    photos.set(undefined);
    searchState.set(SearchState.Error);
  } finally {
    if (page !== undefined) {
      searchPaginationQuery.set({
        ...searchPaginationQuery.get(),
        page: page,
      });
    }
  }
}
