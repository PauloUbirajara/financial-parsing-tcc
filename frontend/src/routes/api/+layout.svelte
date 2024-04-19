<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { ListPlaceholder, P } from "flowbite-svelte";
  import Sidebar from "../../components/Sidebar.svelte";
  import { fade } from "svelte/transition";
  import { routes } from "../../domain/models/routes";
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
    <div class="flex items-center gap-4">
      <Sidebar />
      <P>{routes[$page.route.id]}</P>
    </div>
    <slot />
  {/if}
</div>

<style>
  #api-layout {
    display: grid;
    grid-template-rows: auto 1fr;
  }
</style>
