<script lang="ts">
  import Sidebar from "../../components/Sidebar.svelte";
</script>

<div id="api-layout" class="p-4 flex gap-4 items-center">
  <div>
    <Sidebar />
  </div>
  <slot />
</div>

<style>
  #api-layout {
    display: grid;
    grid-template-rows: auto 1fr;
  }
</style>
