<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div class="min-h-screen p-5 flex flex-col justify-center" id="login-layout">
  <Card size="lg" class="max-h-[700px] h-full">
    <slot />
  </Card>
</div>

<style>
  #login-layout {
    background:
      linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
      url("/auth/login.jpg") fixed no-repeat center/cover;
  }
</style>
