<script lang="ts">
  import { Card } from "flowbite-svelte";
</script>

<div class="h-full p-5" id="register-layout">
  <Card>
    <slot />
  </Card>
</div>

<style>
  #register-layout {
    background-image: url("/auth/register.jpg");
    background-size: cover;
  }
</style>
