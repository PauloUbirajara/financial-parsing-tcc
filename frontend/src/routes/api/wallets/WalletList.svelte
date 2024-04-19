<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { Dropdown, DropdownItem, TextPlaceholder } from "flowbite-svelte";
  import type { Wallet } from "../../../domain/models/wallet";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../components/DeleteModelModal.svelte";
  import ModelList from "../../../components/ModelList.svelte";

  export let onAdd: Function;

  let showDeleteWalletModal = false;

  async function onDelete() {
    if (selectedWallet === null) {
      console.warn("Cannot delete null wallet");
      return;
    }

    const response = await fetch(`/api/wallets/${selectedWallet?.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (response.ok) {
      goto("/api/wallets", { invalidateAll: true });
    }
  }

  $: wallets = $page.data.wallets as Wallet[];
  let selectedWallet: Wallet | null;
  let searchTerm: string = "";
  $: filteredItems = wallets.filter(
    (item) => item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1,
  );
  let fields = {
    Nome: (w: Wallet) => w.name,
    Descrição: (w: Wallet) => w.description,
    Moeda: (w: Wallet) => w.currency.representation,
  };
</script>

{#if $navigating}
  <TextPlaceholder
    divClass="space-y-2.5 animate-pulse w-full mx-auto container"
  />
{:else}
  <ModelList
    {onAdd}
    {onDelete}
    {fields}
    bind:searchTerm
    bind:selectedModel={selectedWallet}
    bind:filteredItems
  />

  <Dropdown triggeredBy=".dots-menu">
    <DropdownItem href={`/api/wallets/${selectedWallet?.id}`}>Ver</DropdownItem>
    <DropdownItem href={`/api/wallets/${selectedWallet?.id}/edit`}>
      Editar
    </DropdownItem>
    <DropdownItem slot="footer" on:click={() => (showDeleteWalletModal = true)}>
      Apagar
    </DropdownItem>
  </Dropdown>

  {#if selectedWallet !== null}
    <DeleteModelModal
      bind:showDeleteModal={showDeleteWalletModal}
      title={`Deseja apagar a carteira "${selectedWallet?.name}"?`}
      {onDelete}
    />
  {/if}
{/if}
