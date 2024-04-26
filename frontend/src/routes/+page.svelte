<script lang="ts">
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import FirstContent from "../components/landing-page/FirstContent.svelte";
  import SecondContent from "../components/landing-page/SecondContent.svelte";
  import ThirdContent from "../components/landing-page/ThirdContent.svelte";
  import JoinContent from "../components/landing-page/JoinContent.svelte";

  $: isLogged = Boolean($page.data.accessToken);

  let screenHeight = 1280;
  onMount(() => {
    screenHeight = window.innerHeight;
  });
</script>

<main class="min-h-screen">
  <Navbar {isLogged} />

  <div class="snap-y snap-proximity">
    <FirstContent {screenHeight} />
    <SecondContent {screenHeight} />
    <ThirdContent {screenHeight} />
    <JoinContent {screenHeight} />
  </div>

  <Footer />
</main>

<style>
  main {
    display: grid;
    grid-template-rows: auto 1fr auto;
  }
</style>
