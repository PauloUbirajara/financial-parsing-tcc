<script lang="ts">
  import { fly } from "svelte/transition";

  export let screenHeight: number;
</script>

<div
  class="third-content snap-center"
  transition:fly={{ x: -20 }}
  style={`height: ${screenHeight}px`}
>
  <p class="text-4xl font-bold">
    Exporte suas transações em formatos convenientes
  </p>
</div>

<style>
</style>
