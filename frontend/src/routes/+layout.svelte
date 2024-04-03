<script lang="ts">
  import "../app.pcss";
  import {
    Footer,
    FooterBrand,
    FooterCopyright,
    NavBrand,
    NavHamburger,
    Button,
    Dropdown,
    DropdownHeader,
    DropdownDivider,
    DropdownItem,
    NavLi,
    DarkMode,
    NavUl,
    Navbar,
  } from "flowbite-svelte";

  export let data
  let isLogged = Boolean(data?.isLogged)
  let username = data?.username
</script>

<main
  class="min-h-screen grid"
  style="grid-template-rows: auto 1fr auto"
>
  <Navbar>
    <NavBrand href="/">
      <img src="/images/logo.png" class="me-3 h-6 sm:h-9 rounded" alt="Financial Parsing Logo" />
      <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Financial Parsing</span>
    </NavBrand>

    <div class="flex items-center md:order-2" >
      {#if isLogged}
        <Button color="alternative" id="profile-dropdown">Olá, {username}</Button>
      {/if}
      <NavHamburger class1="w-full md:flex md:w-auto md:order-1" />
    </div>

    {#if isLogged}
      <Dropdown placement="bottom" triggeredBy="#profile-dropdown">
        <DropdownHeader>
          <span class="block text-sm">{username}</span>
          <span class="block truncate text-sm font-medium">name@flowbite.com</span>
        </DropdownHeader>
        <DropdownItem>Dashboard</DropdownItem>
        <DropdownItem>Configurações</DropdownItem>
        <DropdownDivider />
        <DropdownItem>Sair</DropdownItem>
      </Dropdown>
    {/if}
    <NavUl>
      <NavLi href="/api/dashboard">Dashboard</NavLi>
      <NavLi href="/api/transactions">Transações</NavLi>
      <NavLi href="/api/wallets">Carteiras</NavLi>
      <NavLi href="/api/exports">Exportações</NavLi>
    </NavUl>
  </Navbar>

  <div id="page-container" class="p-10">
    <slot />
  </div>

  <div class="relative">
    <Footer footerType="default" class="fixed w-full bottom-0" style="position: sticky">
      <FooterBrand href="/">
        <img src="/images/logo.png" class="me-3 h-6 sm:h-9 rounded" alt="Financial Parsing Logo" />
        <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Financial Parsing</span>
      </FooterBrand>
      <FooterCopyright href="/" by="Financial Parsing™" />
      <DarkMode />
    </Footer>
  </div>
</main>
