import { writable } from "svelte/store";
import type { SearchPhotoResult } from "./types";

export const imageUrls = writable<SearchPhotoResult | undefined>(undefined);