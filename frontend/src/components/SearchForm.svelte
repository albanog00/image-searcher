<script lang="ts">
  import { imageUrls } from "@/lib/store";
  import LoadingSpinner from "./LoadingSpinner.svelte";
  import { convertToObject } from "typescript";

  enum SearchState {
    Default,
    Finished,
    Searching,
    Error,
  }

  let searchState: SearchState = SearchState.Default;

  let searchQuery: string = "";
  let searchController = new AbortController();

  function onAbort(this: AbortSignal, event: Event) {
    searchState = SearchState.Error;
  }
  searchController.signal.addEventListener("abort", onAbort, {
    once: false,
  });

  async function onSubmit(event: Event) {
    searchState = SearchState.Searching;
    try {
      const res = await fetch(`/api/search?q=${searchQuery}`, {
        signal: searchController.signal,
      });

      // Simulate server request
      await new Promise((resolve, reject) => {
        setTimeout(resolve, 1000);
      });

      imageUrls.set((await res.json()) as string[]);
      searchState = SearchState.Finished;
    } catch (error) {
      searchState = SearchState.Error;
      console.error("Request Failed", error);
    }
  }
</script>

<form on:submit|preventDefault={onSubmit}>
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
</form>
