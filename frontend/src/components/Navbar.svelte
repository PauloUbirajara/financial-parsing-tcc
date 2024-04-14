<script lang="ts">
  import {
    Navbar,
    NavBrand,
    NavHamburger,
    NavUl,
    NavLi,
  } from "flowbite-svelte";

  export let isLogged: boolean;
</script>

<Navbar>
  <NavBrand href="/">
    <img src="logo.png" class="me-3 h-6 sm:h-9" alt="Financial Parsing Logo" />
    <span
      class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
      >Financial Parsing</span
    >
  </NavBrand>
  <NavHamburger />
  <NavUl>
    {#if isLogged}
      <NavLi href="/dashboard">Dashboard</NavLi>
      <NavLi href="/transactions">Transações</NavLi>
      <NavLi href="/wallets">Carteiras</NavLi>
      <NavLi href="/categories">Categorias</NavLi>
      <NavLi href="/currencies">Moedas</NavLi>
    {:else}
      <NavLi href="/auth/login">Acessar conta</NavLi>
      <NavLi href="/auth/register">Criar conta</NavLi>
    {/if}
  </NavUl>
</Navbar>
