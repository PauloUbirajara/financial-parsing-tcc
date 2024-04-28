<script lang="ts">
  import "../app.pcss";
  import Toast from "../components/Toast.svelte";
</script>

<div class="bg-custom-background h-100" id="main-layout">
  <slot />
</div>

<div class="toasts">
  <Toast />
</div>

<style>
  .toasts {
    position: fixed;
    top: 1.5rem;
    right: 1.5rem;
    z-index: 50;
  }
</style>
