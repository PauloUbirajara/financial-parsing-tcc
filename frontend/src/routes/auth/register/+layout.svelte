<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div class="h-full p-5 flex flex-col justify-center" id="register-layout">
  <Card size="lg" class="max-h-[700px] h-full">
    <slot />
  </Card>
</div>

<style>
  #register-layout {
    background:
      linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
      url("/auth/register.jpg") fixed no-repeat center/cover;
  }
</style>
