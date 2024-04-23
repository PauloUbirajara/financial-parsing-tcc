<script lang="ts">
  import { page } from "$app/stores";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";

  $: isLogged = Boolean($page.data.accessToken);
</script>

<main class="min-h-screen">
  <!-- Navbar -->
  <Navbar {isLogged} />

  <!-- First Content -->
  <div class="" id="first-content"></div>

  <!-- Second Content -->
  <div class="" id="second-content"></div>

  <!-- Third Content -->
  <div class="" id="third-content"></div>

  <!-- Join Content -->
  <div class="" id="join-content"></div>

  <!-- Footer -->
  <Footer />
</main>

<style>
  main {
    display: grid;
    grid-template-rows: auto repeat(3, 1fr) auto;
  }
</style>
