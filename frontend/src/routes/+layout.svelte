<script lang="ts">
  import "../app.pcss";
  import {
    Footer,
    FooterBrand,
    FooterCopyright,
    NavBrand,
    NavHamburger,
    NavLi,
    NavUl,
    Navbar,
  } from "flowbite-svelte";

  export let data
  let isLogged = Boolean(data?.isLogged)
</script>

<main
  class="min-h-screen grid"
  style="grid-template-rows: auto 1fr auto"
>
  <Navbar
    class="px-2 sm:px-4 py-2.5 w-full border-b"
  >
    <NavBrand href="/">
      <img
        src="/images/logo.png"
        class="me-3 h-6 sm:h-9 rounded"
        alt="Financial Parsing"
      />
    </NavBrand>
    <NavHamburger />
    <NavUl>
      {#if isLogged}
        <NavLi href="/api/dashboard">Dashboard</NavLi>
        <NavLi href="/api/wallets">Carteiras</NavLi>
        <NavLi href="/api/transactions">Transações</NavLi>
        <NavLi href="/api/exports">Exportações</NavLi>
      {/if}
      {#if !isLogged}
        <NavLi href="/auth/login">Acessar conta</NavLi>
        <NavLi href="/auth/register">Criar conta</NavLi>
      {/if}
    </NavUl>
  </Navbar>
  <!-- <DarkMode -->
  <!--   class="aspect-square text-primary-500 dark:text-primary-600 border dark:border-gray-800" -->
  <!-- /> -->

  <div id="page-container" class="p-10">
    <slot />
  </div>

  <div class="relative">
    <Footer footerType="default" class="fixed w-full bottom-0" style="position: sticky">
      <FooterBrand
        href="/"
        src="/images/logo.png"
        alt="Financial Parsing Logo"
        imgClass="me-3 h-8 rounded"
      />
      <FooterCopyright href="/" by="Financial Parsing™" />
    </Footer>
  </div>
</main>
