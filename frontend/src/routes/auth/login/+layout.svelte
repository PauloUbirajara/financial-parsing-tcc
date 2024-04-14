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
    background-image: url("/auth/login.jpg");
    background-size: cover;
  }
</style>
