<script lang="ts">
  import { TextPlaceholder, Button, Footer, Heading, P } from "flowbite-svelte";
  import Navbar from "../components/Navbar.svelte";
  import { page, navigating } from "$app/stores";
  import { fade } from "svelte/transition";

  const isLogged = Boolean($page.data.access);
  const status = $page.status || "Erro";
  const error = $page.error?.message || "Erro ao buscar a página solicitada";
  console.warn({ status, error });
</script>

<Navbar {isLogged} />

{#if $navigating}
  <div class="container mx-auto my-5">
    <TextPlaceholder size="w-full" />
  </div>
{:else}
  <div class="error-page-content h-max" transition:fade>
    <Heading tag="h3" class="text-center">
      Erro ao buscar a página solicitada
    </Heading>
    <Button href="/">Voltar à página inicial</Button>
  </div>
{/if}

<Footer />

<style>
  .error-page-content {
    display: grid;
    justify-content: center;
    padding: 2rem;
    gap: 1.5rem;
  }
</style>
