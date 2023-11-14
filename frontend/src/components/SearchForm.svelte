<script lang="ts">
  let searchQuery: string = "";
  function onSubmit(event: Event) {
    console.log(searchQuery);
  }
</script>

<form on:submit|preventDefault={onSubmit}>
  <div class="grid grid-flow-col gap-2">
    <input
      bind:value={searchQuery}
      class="rounded-md border-2 border-gray-300 p-2 focus:outline-none focus:ring focus:ring-violet-300"
      type="text"
      placeholder="Search images..."
    />
    <button
      disabled={!searchQuery}
      class="w-40 rounded-md border-2 border-gray-300 bg-indigo-700 p-2 text-white hover:bg-indigo-500 active:bg-indigo-700 active:ring active:ring-violet-300 disabled:cursor-not-allowed disabled:opacity-70"
      type="submit"
    >
      Search
    </button>
  </div>
</form>
