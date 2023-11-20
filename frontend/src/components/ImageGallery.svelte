<script lang="ts">
  import { imageUrls as photos } from "@/lib/store";
  import Image from "./Image.svelte";
</script>

<section
  class="container m-auto columns-1 gap-4 space-y-4 p-4 sm:columns-2 md:columns-3 lg:columns-4"
  id="gallery"
>
  {#if $photos}
    {#each $photos.results as photo}
      <Image {photo} />
    {/each}
  {/if}
</section>
