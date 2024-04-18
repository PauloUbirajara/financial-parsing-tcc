<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div class="h-full p-5 flex flex-col justify-center" id="login-layout">
  <Card size="lg" class="max-h-[700px] h-full">
    <slot />
  </Card>
</div>

<style>
  #login-layout {
    background-image: url("/auth/login.jpg");
    background-size: cover;
  }
</style>
