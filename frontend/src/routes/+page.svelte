<script lang="ts">
  import Navbar from "$components/Navbar.svelte";
  import Footer from "$components/Footer.svelte";
  import type { PageData } from "./$types";

  export let data: PageData;
  $: isLogged = data.isLogged;
</script>

<main>
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
    height: 100%;
    display: grid;
    grid-template-rows: auto repeat(3, 1fr) auto;
  }
</style>
