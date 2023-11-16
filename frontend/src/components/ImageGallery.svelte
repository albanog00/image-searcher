<script lang="ts">
  import { imageUrls } from "@/lib/store";
</script>

<section
  class="container m-auto columns-1 gap-4 space-y-4 p-4 sm:columns-2 md:columns-3 lg:columns-4"
  id="gallery"
>
  {#if $imageUrls}
    {#each $imageUrls?.urls as imageUrl}
      <div class="flex">
        <img
          alt="gallery"
          class="h-auto max-w-full rounded-lg shadow-xl drop-shadow-xl transition duration-300 hover:z-10 hover:scale-110 hover:outline-none hover:ring-8 hover:ring-violet-300"
          loading="lazy"
          src={imageUrl}
        />
      </div>
    {/each}
  {/if}
</section>
