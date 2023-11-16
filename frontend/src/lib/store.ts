import { writable } from "svelte/store";
import type { SearchResult } from "./types";

export const imageUrls = writable<SearchResult | undefined>(undefined);