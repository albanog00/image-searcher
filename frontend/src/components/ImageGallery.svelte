<script lang="ts">
  import { photos, searchPaginationQuery } from "@/lib/store";
  import Image from "./Image.svelte";
  import Separator from "./Separator.svelte";
  import PaginationPage from "./PaginationPage.svelte";
  import { blur } from "svelte/transition";
</script>

{#if $photos}
  {#if $photos.total > 0}
    <section
      class="container m-auto columns-1 gap-4 space-y-4 p-4 sm:columns-2 md:columns-3 lg:columns-4"
      id="gallery"
    >
      {#each $photos.results as photo}
        <Image {photo} />
      {/each}
    </section>
    {#if $searchPaginationQuery.autoPaging}
      <div
        class="m-auto w-min animate-bounce cursor-pointer drop-shadow-lg"
        transition:blur={{ duration: 300 }}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          height="3em"
          viewBox="0 0 512 512"
          ><path
            d="M256 0a256 256 0 1 0 0 512A256 256 0 1 0 256 0zM127 281c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l71 71L232 136c0-13.3 10.7-24 24-24s24 10.7 24 24l0 182.1 71-71c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9L273 393c-9.4 9.4-24.6 9.4-33.9 0L127 281z"
          /></svg
        >
      </div>
    {:else}
      <Separator />
      <PaginationPage />
    {/if}
  {:else}
    No results found
  {/if}
{/if}
