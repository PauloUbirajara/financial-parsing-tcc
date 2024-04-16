<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div
  class="h-full p-5 flex flex-col justify-center"
  id="forgot-password-layout"
>
  <Card size="lg" class="max-h-[700px] h-full">
    <slot />
  </Card>
</div>

<style>
  #forgot-password-layout {
    background-image: url("/auth/forgot-password.jpg");
    background-size: cover;
  }
</style>
