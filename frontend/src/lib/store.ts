import {
  SearchState,
  type SearchPaginationQuery,
  type SearchPhotoResult,
} from "./types";
import { atom } from "nanostores";

export const photos = atom<SearchPhotoResult | undefined>(undefined);
export const searchImageQuery = atom<string | undefined>(undefined);
export const searchPaginationQuery = atom<SearchPaginationQuery>({
  page: 1,
  perPage: 10,
  autoPaging: false,
  totalPages: 0,
});
export const searchState = atom<SearchState>(SearchState.Default);
