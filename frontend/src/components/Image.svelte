<script lang="ts">
  import type { Photo } from "@/lib/types";
  import ImageModal from "./ImageModal.svelte";
  import { blur } from "svelte/transition";
  import { onMount } from "svelte";

  let thisImage: HTMLImageElement = new Image();
  let loaded = false;

  onMount(() => {
    thisImage.src = photo.urls.regular;

    thisImage.onload = () => {
      loaded = true;
    };
  });

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
{#if loaded}
  <div class="flex" transition:blur={{ duration: 300 }}>
    <button on:click|preventDefault={openFullImage}>
      <img
        bind:this={thisImage}
        alt="gallery"
        class="relative h-auto max-w-full cursor-pointer rounded-lg border-4 border-white shadow-xl drop-shadow-xl transition duration-300 hover:z-10 hover:scale-110 hover:outline-none hover:ring-8 hover:ring-violet-300"
        loading="lazy"
        src={photo.urls.regular}
      />
    </button>
  </div>
{/if}
