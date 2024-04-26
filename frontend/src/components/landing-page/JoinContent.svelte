<script lang="ts">
  import { Button } from "flowbite-svelte";
  import { fly } from "svelte/transition";

  export let screenHeight: number;
</script>

<div
  class="join-content"
  transition:fly={{ x: -20 }}
  style={`height: ${screenHeight}px; min-height: 300px;`}
>
  <p class="text-4xl font-bold">Experimente agora a nossa plataforma</p>
  <Button href="/auth/login">Juntar-se</Button>
</div>

<style>
  .join-content {
    display: grid;
    grid-template-rows: min-content min-content;
    justify-content: center;
    align-items: center;
    align-content: center;
    gap: 2rem;
  }
</style>
