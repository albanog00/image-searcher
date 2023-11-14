<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  interface $$Props extends HTMLInputAttributes {}
</script>

<input
  {...$$restProps}
  class={`rounded-lg border border-gray-400 p-2 focus:border ${$$restProps.class}`}
/>
