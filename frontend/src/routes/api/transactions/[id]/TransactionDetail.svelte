<script lang="ts">
  import { navigating } from "$app/stores";
  import { Button, P, TextPlaceholder } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    EditSolid,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../../components/model/DeleteModelModal.svelte";
  import transactionSerializer from "../../../../data/usecases/modelSerializer/transaction";
  import type { IModelSerializer } from "../../../../domain/usecases/modelSerializer";
  import type { Transaction } from "../../../../domain/models/transaction";

  function onEdit() {
    goto(`/api/transactions/${transaction.id}/edit`);
  }

  export let transaction: Transaction;
  let showDeleteTransactionModal: boolean = false;

  let serializer: IModelSerializer = transactionSerializer;

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
  {#each Object.entries(serializer.serialize(transaction)) as [key, value]}
    <div class="content">
      <P>{key}</P>
      <P size="3xl" weight="semibold">{value}</P>
    </div>
  {/each}

  <DeleteModelModal
    bind:showDeleteModal={showDeleteTransactionModal}
    title={`Deseja apagar a transação "${transaction.name}"?`}
    {onDelete}
  />
{/if}
