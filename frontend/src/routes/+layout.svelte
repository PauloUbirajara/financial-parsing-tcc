<script lang="ts">
  import "../app.pcss";
  import Toast from "../components/Toast.svelte";
</script>

<div class="bg-custom-background" id="main-layout">
  <slot />
</div>

<div class="toasts">
  <Toast />
</div>

<style>
  #main-layout {
    min-height: 100%;
  }

  .toasts {
    position: fixed;
    top: 1.5rem;
    right: 1.5rem;
  }
</style>
