<script lang="ts">
  export let imageUrl: string;
  export let closeFullImage: (event: Event) => void;
  export let handleKey: (event: KeyboardEvent) => void;
</script>

<svelte:window on:keydown={handleKey} />
<div class="fixed bottom-0 left-0 right-0 top-0 z-50">
  <div
    class="flex h-full w-full items-center justify-center bg-gray-400 bg-opacity-50"
  >
    <div class="rounded-md bg-white px-4 pb-4 pt-2 text-end">
      <button
        on:click={closeFullImage}
        class="text-md cursor-pointer font-medium text-black"
      >
        <span>X</span>
      </button>
      <img
        alt="gallery"
        class="max-h-[90vh] max-w-[90vw] rounded-lg shadow-xl drop-shadow-xl"
        loading="lazy"
        src={imageUrl}
      />
    </div>
  </div>
</div>
