<script lang="ts">
  import { imageUrls } from "@/lib/store";
  import { SearchState, type SearchResult } from "@/lib/types";
  import { api } from "@/lib/api";
  import LoadingSpinner from "./LoadingSpinner.svelte";

  let searchState: SearchState = SearchState.Default;
  let searchQuery: string = "";

  async function onSubmit(event: Event) {
    function onAbort(this: AbortSignal, event: Event) {
      searchState = SearchState.Error;
      console.log("Request aborted");
    }

    const searchController = new AbortController();
    searchController.signal.addEventListener("abort", onAbort);

    searchState = SearchState.Searching;
    try {
      const res = await api<SearchResult>(`/search?q=${searchQuery}`, {
        method: "GET",
        signal: searchController.signal,
      });

      // Simulating server request
      await new Promise((resolve, reject) => {
        setTimeout(resolve, 1000);
      });

      const json = await res.json();
      imageUrls.set(json);

      searchState = SearchState.Finished;
    } catch (error) {
      console.error(error);

      searchController.abort();
      imageUrls.set({ urls: [] });

      searchState = SearchState.Error;
    }

    searchController.signal.removeEventListener("abort", onAbort);
  }
</script>

<form
  class="flex flex-col items-center justify-center"
  on:submit|preventDefault={onSubmit}
>
  <div class="grid grid-flow-col gap-2">
    <input
      bind:value={searchQuery}
      class="rounded-md border-2 border-gray-300 p-2 font-medium focus:outline-none focus:ring focus:ring-violet-300"
      type="text"
      placeholder="Search images..."
    />
    <button
      disabled={!searchQuery || searchState === SearchState.Searching}
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
  </div>
  {#if searchState === SearchState.Error}
    <span class="text-red-600">
      Error occurred while sending request to the server. Please try again.
    </span>
  {/if}
</form>
