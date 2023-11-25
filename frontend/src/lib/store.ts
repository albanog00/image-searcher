import type { SearchPaginationQuery, SearchPhotoResult } from "./types";
import { atom } from "nanostores";

export const photos = atom<SearchPhotoResult | undefined>(undefined);
export const searchImageQuery = atom<string | undefined>(undefined);
export const searchPaginationQuery = atom<SearchPaginationQuery>({ page: 0, perPage: 10, autoPaging: false, totalPages: 0 })
