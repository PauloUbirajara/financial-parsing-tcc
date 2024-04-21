<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { Button, P, TextPlaceholder } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    EditSolid,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../../components/model/DeleteModelModal.svelte";

  function onEdit() {
    goto(`/api/wallets/${wallet.id}/edit`);
  }

  let wallet = $page.data.wallet;
  console.log({ wallet });
  let showDeleteWalletModal: boolean = false;

  let fields = {
    Nome: wallet.name,
    Descrição: wallet.description || "-",
    Moeda: `${wallet.currency.name} (${wallet.currency.representation})`,
  };

  async function onDelete() {
    const response = await fetch(`/api/wallets/${wallet.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (response.ok) {
      goto("/api/wallets", { invalidateAll: true });
    }
  }
</script>

{#if $navigating}
  <TextPlaceholder divClass="space-y-2.5 animate-pulse mx-auto w-full" />
{:else}
  <div class="actions flex justify-between">
    <Button
      outline={true}
      class="!p-2"
      size="lg"
      color="primary"
      href="/api/wallets"
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <div class="flex gap-4 items-center">
      <Button on:click={() => onEdit()}>
        <EditSolid />
        Editar
      </Button>
      <Button color="red" on:click={() => (showDeleteWalletModal = true)}>
        <TrashBinSolid />
        Editar
      </Button>
    </div>
  </div>
  {#each Object.entries(fields) as [key, value]}
    <div class="content">
      <P>{key}</P>
      <P size="3xl" weight="semibold">{value}</P>
    </div>
  {/each}

  <DeleteModelModal
    bind:showDeleteModal={showDeleteWalletModal}
    title={`Deseja apagar a carteira "${wallet.name}"?`}
    {onDelete}
  />
{/if}
