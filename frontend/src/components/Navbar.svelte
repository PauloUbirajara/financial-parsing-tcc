<script context="module" lang="ts">
  export let data: PageData;
</script>

<script lang="ts">
  import {
    Navbar,
    NavBrand,
    NavHamburger,
    NavUl,
    NavLi,
  } from "flowbite-svelte";
  import type { PageData } from "../routes/$types";

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
      <NavLi href="/api/dashboard">Dashboard</NavLi>
      <NavLi href="/api/transactions">Transações</NavLi>
      <NavLi href="/api/wallets">Carteiras</NavLi>
      <NavLi href="/api/categories">Categorias</NavLi>
      <NavLi href="/api/currencies">Moedas</NavLi>
    {:else}
      <NavLi href="/auth/login">Acessar conta</NavLi>
      <NavLi href="/auth/register">Criar conta</NavLi>
    {/if}
  </NavUl>
</Navbar>
