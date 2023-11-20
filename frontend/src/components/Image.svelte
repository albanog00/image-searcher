<script lang="ts">
  import type { Photo } from "@/lib/types";
  import ImageModal from "./ImageModal.svelte";
  export let photo: Photo;

  let showFullImage = false;
  const openFullImage = (event: Event) => (showFullImage = true);
  const closeFullImage = (event: Event) => (showFullImage = false);
</script>

{#if showFullImage}
  <div class="space-y-0">
    <ImageModal {photo} {closeFullImage} />
  </div>
{/if}
<div class="flex">
  <button on:click|preventDefault={openFullImage}>
    <img
      alt="gallery"
      class="relative h-auto max-w-full cursor-pointer rounded-lg shadow-xl drop-shadow-xl transition duration-300 hover:z-10 hover:scale-110 hover:outline-none hover:ring-8 hover:ring-violet-300"
      loading="lazy"
      src={photo.urls.regular}
    />
  </button>
</div>
