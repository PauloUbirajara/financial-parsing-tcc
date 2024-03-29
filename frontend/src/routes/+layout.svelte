<script lang="ts">
  import "../app.pcss";
  import {
    DarkMode,
    Footer,
    FooterBrand,
    FooterCopyright,
    NavBrand,
    NavHamburger,
    NavLi,
    NavUl,
    Navbar,
  } from "flowbite-svelte";

  let isLogged = false;
</script>

<Navbar
  class="px-2 sm:px-4 py-2.5 w-full z-20 top-0 start-0 border-b"
  style="position: sticky"
>
  <NavBrand href="/">
    <img
      src="/images/logo.png"
      class="me-3 h-6 sm:h-9 rounded"
      alt="Financial Parsing"
    />
  </NavBrand>
  <NavHamburger />
  <div class="flex">
    <NavUl>
      {#if isLogged}
        <NavLi href="/dashboard">Dashboard</NavLi>
        <NavLi href="/wallets">Carteiras</NavLi>
        <NavLi href="/transactions">Transações</NavLi>
        <NavLi href="/exports">Exportações</NavLi>
      {:else}
        <NavLi href="/auth/login">Acessar conta</NavLi>
        <NavLi href="/auth/register">Criar conta</NavLi>
      {/if}
    </NavUl>
    <DarkMode
      class="text-primary-500 dark:text-primary-600 border dark:border-gray-800"
    />
  </div>
</Navbar>

<main class="p-10 h-full">
  <slot />
</main>

<Footer footerType="default">
  <FooterBrand
    href="/"
    src="/images/logo.png"
    alt="Financial Parsing Logo"
    imgClass="me-3 h-8 rounded"
  />
  <FooterCopyright href="/" by="Financial Parsing™" />
</Footer>
