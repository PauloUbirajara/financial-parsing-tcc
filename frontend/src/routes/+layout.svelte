<script lang="ts">
  import "../app.pcss";
</script>

<div class="bg-custom-background" id="main-layout">
  <slot />
</div>

<style>
  #main-layout {
    min-height: 100%;
  }
</style>
