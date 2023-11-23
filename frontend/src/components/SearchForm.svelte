<script lang="ts">
  import { api } from "@/lib/api";
  import {
    imageUrls,
    searchPaginationQuery,
    searchImageQuery,
  } from "@/lib/store";
  import { SearchState, type SearchPhotoResult } from "@/lib/types";

  import LoadingSpinner from "./Icons/LoadingSpinner.svelte";
  import ImageGallery from "./ImageGallery.svelte";
  import FilterPopover from "./FilterPopover.svelte";

  let searchState: SearchState = SearchState.Default;

  async function onSubmit(event: Event) {
    searchState = SearchState.Searching;
    try {
      const res = await api<SearchPhotoResult>(
        `/search?q=${$searchImageQuery}&page=${$searchPaginationQuery?.page}&per_page=${$searchPaginationQuery?.perPage}`,
        {
          method: "GET",
        },
      );

      const json = await res.json();
      searchState = SearchState.Finished;

      imageUrls.set(json.result?.data);
    } catch (error) {
      console.error(error);
      imageUrls.set(undefined);
      searchState = SearchState.Error;
    }
  }
</script>

<div class="flex flex-col items-center justify-center gap-4">
  <div class="flex flex-row gap-2">
    <form class="grid grid-flow-col gap-2" on:submit|preventDefault={onSubmit}>
      <input
        required
        bind:value={$searchImageQuery}
        class="rounded-md border-2 border-gray-300 p-2 font-medium focus:outline-none focus:ring focus:ring-violet-300"
        type="text"
        placeholder="Search images..."
      />

      <button
        disabled={!$searchImageQuery || searchState === SearchState.Searching}
        class="flex w-40 justify-center gap-2 rounded-md border-2 border-gray-300 bg-indigo-700 p-2 font-medium text-white hover:bg-indigo-500 active:bg-indigo-700 active:ring active:ring-violet-300 disabled:cursor-not-allowed disabled:opacity-70"
        type="submit"
      >
        {#if searchState === SearchState.Searching}
          <LoadingSpinner />
          Loading...
        {:else}
          Search
        {/if}
      </button>
      <FilterPopover />
    </form>
  </div>

  {#if searchState === SearchState.Error}
    <span class="text-red-600">
      Error occurred while sending request to the server. Please try again.
    </span>
  {:else}
    <ImageGallery />
  {/if}
</div>
