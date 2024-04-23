<script lang="ts">
  import { fade } from "svelte/transition";
  import { ListPlaceholder } from "flowbite-svelte";
  import { navigating } from "$app/stores";
</script>

<div
  id="api-layout"
  class="p-4 gap-8 h-full bg-custom-background"
  transition:fade
>
  {#if $navigating}
    <div class="w-full-screen">
      <ListPlaceholder
        divClass="p-4 space-y-4 rounded border border-gray-200 divide-y divide-gray-200 shadow animate-pulse dark:divide-gray-700 md:p-6 dark:border-gray-700"
      />
    </div>
  {:else}
    <slot />
  {/if}
</div>

<style>
  #api-layout {
    display: grid;
    grid-template-rows: auto 1fr;
  }
</style>
