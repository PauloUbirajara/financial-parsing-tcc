<script lang="ts">
  import { navigating } from "$app/stores";
  import { A, Badge, Button, P, TextPlaceholder } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    ChevronRightOutline,
    EditSolid,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../../components/model/DeleteModelModal.svelte";
  import type { Transaction } from "../../../../domain/models/transaction";

  function onEdit() {
    goto(`/api/transactions/${transaction.id}/edit`);
  }

  export let transaction: Transaction;
  let showDeleteTransactionModal: boolean = false;

  async function onDelete() {
    const response = await fetch(
      `/api/transactions/${transaction.id}?/delete`,
      {
        method: "POST",
        body: JSON.stringify({}),
      },
    );
    if (response.ok) {
      goto("/api/transactions", { invalidateAll: true });
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
      href="/api/transactions"
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <div class="flex gap-4 items-center">
      <Button on:click={() => onEdit()}>
        <EditSolid />
        Editar
      </Button>
      <Button color="red" on:click={() => (showDeleteTransactionModal = true)}>
        <TrashBinSolid />
        Editar
      </Button>
    </div>
  </div>
  <div class="content">
    <P>Nome</P>
    <P size="3xl" weight="semibold">{transaction.name}</P>
  </div>
  <div class="content">
    <P>Descrição</P>
    <P size="3xl" weight="semibold">{transaction.description}</P>
  </div>
  <div class="content">
    <P>Data da transação</P>
    <P size="3xl" weight="semibold">
      {new Date(transaction.transaction_date).toLocaleString("pt-br", {
        day: "numeric",
        month: "numeric",
        year: "numeric",
      })}
    </P>
  </div>
  <div class="content">
    <P>Valor</P>
    <P size="3xl" weight="semibold">{transaction.value}</P>
  </div>
  <div class="content">
    <P>Carteira</P>
    <div class="flex items-center gap-4">
      <A
        aClass="inline-flex items-center hover:underline text-3xl font-semibold"
        href={`/api/wallets/${transaction.wallet.id}`}
      >
        {transaction.wallet.name}
        <ChevronRightOutline />
      </A>
    </div>
  </div>
  <div class="content">
    <P class="mb-2">Categorias</P>
    <div class="flex flex-wrap p-4 bg-white rounded gap-2">
      {#each transaction.categories as item}
        <Badge color="primary" href={`/api/categories/${item.id}`} large>
          {item.name}
        </Badge>
      {/each}
    </div>
  </div>

  <DeleteModelModal
    bind:showDeleteModal={showDeleteTransactionModal}
    title={`Deseja apagar a transação "${transaction.name}"?`}
    {onDelete}
  />
{/if}
