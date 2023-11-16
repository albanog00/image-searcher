<script lang="ts">
  import { imageUrls } from "@/lib/store";
  import Image from "./Image.svelte";
</script>

<section
  class="container m-auto columns-1 gap-4 space-y-4 p-4 sm:columns-2 md:columns-3 lg:columns-4"
  id="gallery"
>
  {#if $imageUrls}
    {#each $imageUrls?.urls as imageUrl}
      <Image {imageUrl} />
    {/each}
  {/if}
</section>
