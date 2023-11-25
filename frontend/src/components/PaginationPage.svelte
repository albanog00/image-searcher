<script lang="ts">
  import { searchPaginationQuery } from "@/lib/store";
  import { blur } from "svelte/transition";

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
</script>

<div
  class="grid grid-flow-col items-center justify-center gap-2"
  transition:blur={{ duration: 300 }}
>
  <button
    disabled={$searchPaginationQuery.page <= 0}
    on:click={prevPage}
    class="rounded-lg border-2 border-gray-300 px-3 py-1 drop-shadow-md hover:bg-white"
    >&larr;
  </button>
  {#each { length: $searchPaginationQuery.totalPages } as _, page}
    <button
      disabled={$searchPaginationQuery.page === page}
      class="rounded-lg border-2 border-gray-300 px-3 py-1 drop-shadow-md hover:bg-white"
      >{page + 1}</button
    >
  {/each}<button
    disabled={$searchPaginationQuery.page >=
      $searchPaginationQuery.totalPages - 1}
    on:click={nextPage}
    class="rounded-lg border-2 border-gray-300 px-3 py-1 drop-shadow-md hover:bg-white"
    >&rarr;
  </button>
</div>
