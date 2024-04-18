<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div class="h-full p-5 flex flex-col justify-center" id="reset-password-layout">
  <Card size="lg" class="max-h-[700px] h-full">
    <slot />
  </Card>
</div>

<style>
  #reset-password-layout {
    background-image: url("/auth/reset-password.jpg");
    background-size: cover;
  }
</style>
