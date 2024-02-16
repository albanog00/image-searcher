<script lang="ts">
  import { search } from "@/lib/search";
  import { photos, searchPaginationQuery } from "@/lib/store";
  import { onMount } from "svelte";
  import { blur } from "svelte/transition";

  let pageToLoad = 1;
  let currentPage = searchPaginationQuery.get().page;

  const nextPage = () => {
    searchPaginationQuery.set({
      ...searchPaginationQuery.get(),
      page: $searchPaginationQuery.page + 1,
    });
  };

  const prevPage = () => {
    searchPaginationQuery.set({
      ...searchPaginationQuery.get(),
      page: $searchPaginationQuery.page - 1,
    });
  };

  const jumpToPage = (pageToLoad: number = 0) => {
    if (pageToLoad < 0 || pageToLoad > $photos?.total_pages!) return;
    searchPaginationQuery.set({
      ...searchPaginationQuery.get(),
      page: pageToLoad,
    });
  };

  onMount(() => {
    const removeListener = searchPaginationQuery.listen(async (value) => {
      if (value.page !== currentPage) {
        currentPage = value.page;
        await search();
      }
    });

    return () => {
      removeListener();
    };
  });
</script>

<div class="flex flex-col items-center justify-center gap-2">
  <div
    class="grid grid-flow-col items-center justify-center gap-2"
    transition:blur={{ duration: 300 }}
  >
    <button
      disabled={$searchPaginationQuery.page <= 0}
      on:click={prevPage}
      class="rounded-lg border-2 border-gray-300 bg-gray-200 px-3 py-1 drop-shadow-md hover:bg-white disabled:cursor-not-allowed disabled:opacity-40"
      >&larr;
    </button>
    {#each new Array($photos?.total_pages)
      .fill("")
      .map((_, i) => i + 1) as page}
      <button
        on:click={() => jumpToPage(page)}
        disabled={$searchPaginationQuery.page === page}
        hidden={page > $searchPaginationQuery.page + 2 ||
          page < $searchPaginationQuery.page - 2}
        class="rounded-lg border-2 border-gray-300 bg-gray-200 px-3 py-1 drop-shadow-md hover:bg-white disabled:cursor-not-allowed disabled:opacity-40"
        >{page}</button
      >
    {/each}
    <button
      disabled={$searchPaginationQuery.page >= ($photos?.total_pages ?? 0) - 1}
      on:click={nextPage}
      class="rounded-lg border-2 border-gray-300 bg-gray-200 px-3 py-1 drop-shadow-md hover:bg-white disabled:cursor-not-allowed disabled:opacity-40"
      >&rarr;
    </button>
  </div>
  <div class="flex items-center justify-center">
    <input
      class="w-24 rounded-md border-2 border-gray-300 p-2 font-medium focus:outline-none focus:ring focus:ring-violet-300"
      bind:value={pageToLoad}
      type="number"
    />
    <button
      on:click={() => jumpToPage(pageToLoad)}
      disabled={pageToLoad <= 0 ||
        pageToLoad == currentPage ||
        pageToLoad - 1 > ($photos?.total_pages ?? 0)}
      class="w-32 gap-2 rounded-md border-2 border-gray-300 bg-indigo-700 p-2 font-medium text-white hover:bg-indigo-500 active:bg-indigo-700 active:ring active:ring-violet-300 disabled:cursor-not-allowed disabled:opacity-70"
    >
      Go
    </button>
  </div>
</div>
