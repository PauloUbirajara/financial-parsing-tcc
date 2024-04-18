<script lang="ts">
  import {
    Drawer,
    Button,
    CloseButton,
    Sidebar,
    SidebarDropdownItem,
    SidebarDropdownWrapper,
    SidebarGroup,
    SidebarItem,
    SidebarWrapper,
  } from "flowbite-svelte";
  import {
    ChartPieSolid,
    CartSolid,
    GridSolid,
    MailBoxSolid,
    UsersSolid,
    ShoppingBagSolid,
    ArrowRightToBracketOutline,
    EditOutline,
    BarsOutline,
    WalletSolid,
    FileChartBarSolid,
    TagSolid,
  } from "flowbite-svelte-icons";
  import { sineIn } from "svelte/easing";

  let hideSidebar: boolean = true;
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn,
  };
</script>

<Button class="!p-2" on:click={() => (hideSidebar = false)}>
  <BarsOutline class="w-6 h-6" /></Button
>
<Drawer
  transitionType="fly"
  {transitionParams}
  bind:hidden={hideSidebar}
  id="sidebar2"
>
  <div class="flex items-center">
    <h5
      id="drawer-navigation-label-3"
      class="text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
    >
      Menu
    </h5>
    <CloseButton
      on:click={() => (hideSidebar = true)}
      class="mb-4 dark:text-white"
    />
  </div>
  <Sidebar>
    <SidebarWrapper
      divClass="overflow-y-auto py-4 px-3 rounded dark:bg-gray-800"
    >
      <SidebarGroup>
        <SidebarItem label="Dashboard" href="/api/dashboard">
          <svelte:fragment slot="icon">
            <ChartPieSolid
              class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
            />
          </svelte:fragment>
        </SidebarItem>

        <SidebarItem label="Transações" href="/api/transactions">
          <svelte:fragment slot="icon">
            <FileChartBarSolid
              class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
            />
          </svelte:fragment>
        </SidebarItem>

        <SidebarItem label="Carteiras" href="/api/wallets">
          <svelte:fragment slot="icon">
            <WalletSolid
              class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
            />
          </svelte:fragment>
        </SidebarItem>

        <SidebarItem label="Categorias" href="/api/categories">
          <svelte:fragment slot="icon">
            <TagSolid
              class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
            />
          </svelte:fragment>
        </SidebarItem>
      </SidebarGroup>
    </SidebarWrapper>
  </Sidebar>
</Drawer>
