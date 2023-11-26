<script lang="ts">
  import type { Photo } from "@/lib/types";
  import { blur, fade } from "svelte/transition";
  import { onMount } from "svelte";
  import ImageModal from "./ImageModal.svelte";
  import Skeleton from "./Skeleton.svelte";

  let thisImage: HTMLImageElement = new Image();
  let loaded = false;

  onMount(() => {
    thisImage.src = photo.urls.small;

    thisImage.onload = () => {
      loaded = true;
    };
  });

  export let photo: Photo;

  let showFullImage = false;
  const openFullImage = (event: Event) => (showFullImage = true);
  const closeFullImage = (event: Event) => (showFullImage = false);
</script>

{#if loaded}
  {#if showFullImage}
    <div class="space-y-0">
      <ImageModal {photo} {closeFullImage} />
    </div>
  {/if}
  <div class="flex" transition:blur={{ duration: 400 }}>
    <button on:click|preventDefault={openFullImage}>
      <img
        bind:this={thisImage}
        alt="gallery"
        class="relative h-auto max-w-full cursor-pointer rounded-lg border-4 border-white shadow-xl drop-shadow-xl transition duration-300 hover:z-10 hover:scale-110 hover:outline-none hover:ring-8 hover:ring-violet-300"
        loading="lazy"
        src={photo.urls.small}
      />
    </button>
  </div>
{:else}
  <div class="flex" transition:fade={{ duration: 300 }}>
    <Skeleton
      styles="h-[240px] w-[350px] rounded-lg border-4 border-white  shadow-xl drop-shadow-xl"
    />
  </div>
{/if}
