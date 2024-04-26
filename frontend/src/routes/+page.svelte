<script lang="ts">
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import FirstContent from "../components/landing-page/FirstContent.svelte";
  import SecondContent from "../components/landing-page/SecondContent.svelte";
  import ThirdContent from "../components/landing-page/ThirdContent.svelte";
  import JoinContent from "../components/landing-page/JoinContent.svelte";
  import { ImagePlaceholder, Spinner } from "flowbite-svelte";

  $: isLogged = Boolean($page.data.accessToken);

  let screenHeight: number | null = null;
  onMount(() => {
    screenHeight = window.innerHeight;
  });
</script>

<main class="min-h-screen">
  <Navbar {isLogged} />

  <div class="landing-page-content">
    {#if !screenHeight}
      <div class="container mx-auto my-5">
        <ImagePlaceholder />
      </div>
    {:else}
      <FirstContent {screenHeight} />
      <SecondContent {screenHeight} />
      <ThirdContent {screenHeight} />
      <JoinContent {screenHeight} />
    {/if}
  </div>

  <Footer />
</main>

<style>
  main {
    display: grid;
    grid-template-rows: auto 1fr auto;
  }
</style>
