<script lang="ts">
  import type { ChangeEventHandler } from "svelte/elements";
  import { searchPaginationQuery } from "@/lib/store";

  import IconMenu from "./Icons/IconMenu.svelte";
  import Popover from "./Popover.svelte";
  import { slide, fade } from "svelte/transition";

  let autoPaging: boolean;
  let perPage: number;

  searchPaginationQuery.subscribe((query) => {
    perPage = query.perPage;
    autoPaging = query.autoPaging;
  });

  const onChangePerPage: ChangeEventHandler<HTMLInputElement> = (event) => {
    const value = Number(event.currentTarget.value);
    if (value > 30) {
      searchPaginationQuery.set({
        ...searchPaginationQuery.get(),
        perPage: 30,
      });
    } else if (value < 10) {
      searchPaginationQuery.set({
        ...searchPaginationQuery.get(),
        perPage: 10,
      });
    }
  };
</script>

<Popover>
  <IconMenu slot="trigger" />

  <div
    class="z-10 flex flex-col gap-2 rounded-lg border-2 border-gray-300 bg-white p-4"
    in:slide={{ duration: 250 }}
    out:fade={{ duration: 150 }}
  >
    <div class="flex flex-col">
      <label class="" for="per_page">Images per page</label>
      <input
        id="per_page"
        on:keyup={onChangePerPage}
        bind:value={perPage}
        max="30"
        min="10"
        class="rounded-md border-2 border-gray-300 p-2 font-medium focus:outline-none focus:ring focus:ring-violet-300"
        type="number"
      />
    </div>
    <div class="flex flex-row gap-2">
      <label class="" for="auto_paging">Auto-paging</label>
      <input
        id="auto_paging"
        bind:checked={autoPaging}
        on:change={function () {
          searchPaginationQuery.set({
            ...searchPaginationQuery.get(),
            autoPaging: autoPaging,
          });
        }}
        class="rounded-md border-2 border-gray-300 p-2 font-medium focus:outline-none focus:ring focus:ring-violet-300"
        type="checkbox"
      />
    </div>
  </div>
</Popover>
