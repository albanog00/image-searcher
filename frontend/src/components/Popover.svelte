<script lang="ts">
  let openPopover: boolean = false;

  const closePopover = (event: KeyboardEvent) => {
    if (event.key == "Escape") {
      openPopover = false;
    }
  };
</script>

<svelte:window on:keydown={closePopover} />
<div>
  <button
    on:click|preventDefault={function () {
      openPopover = !openPopover;
    }}
    class="m-auto cursor-pointer rounded-md border-2 border-gray-300 px-3 py-2 drop-shadow-lg"
  >
    <slot name="trigger" />
  </button>
  <div class="absolute z-10 mt-1 shadow-md">
    {#if openPopover}
      <slot />
    {/if}
  </div>
</div>
