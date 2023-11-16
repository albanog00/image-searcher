<script lang="ts">
  import ImageModal from "./ImageModal.svelte";
  export let imageUrl: string;

  let showFullImage = false;

  const openFullImage = (event: Event) => {
    showFullImage = true;
  };

  const closeFullImage = (event: Event) => {
    showFullImage = false;
  };

  const handleKey = (event: KeyboardEvent) => {
    if (event.key === "Escape") {
      showFullImage = false;
    }
  };
</script>

{#if showFullImage}
  <div class="space-y-0">
    <ImageModal {imageUrl} {closeFullImage} {handleKey} />
  </div>
{/if}
<div class="flex">
  <button on:click|preventDefault={openFullImage}>
    <img
      alt="gallery"
      class="relative h-auto max-w-full cursor-pointer rounded-lg shadow-xl drop-shadow-xl transition duration-300 hover:z-10 hover:scale-110 hover:outline-none hover:ring-8 hover:ring-violet-300"
      loading="lazy"
      src={imageUrl}
    />
  </button>
</div>
