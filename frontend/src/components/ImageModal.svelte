<script lang="ts">
  import type { Photo } from "@/lib/types";

  export let photo: Photo;
  export let closeFullImage: (event: Event) => void;

  const handleKey = (event: KeyboardEvent) => {
    if (event.key === "Escape") {
      closeFullImage(event);
    }
  };
</script>

<svelte:window on:keydown={handleKey} />

<div class="fixed bottom-0 left-0 right-0 top-0 z-40">
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div
    on:click={closeFullImage}
    class="flex h-full w-full items-center justify-center bg-gray-600 bg-opacity-70"
  >
    <div
      on:click|stopPropagation
      class="rounded-md bg-white px-4 pb-4 pt-2 text-end"
    >
      <div class="flex justify-between">
        <span class="self-start text-start italic text-gray-500"
          >by <a href={photo.user.links.self} class="font-semibold text-black">
            {photo.user.username}
          </a></span
        >
        <button
          on:click={closeFullImage}
          class="text-md cursor-pointer font-medium text-black"
        >
          <span>X</span>
        </button>
      </div>
      <img
        alt="gallery"
        class="max-h-[90vh] max-w-[90vw] rounded-lg shadow-xl drop-shadow-xl"
        loading="eager"
        src={photo.urls.full}
      />
    </div>
  </div>
</div>
