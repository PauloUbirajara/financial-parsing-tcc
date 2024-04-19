<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { Button, P, TextPlaceholder } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    EditSolid,
    TrashBinSolid,
  } from "flowbite-svelte-icons";

  export let onEdit: Function;

  $: wallet = $page.data.wallet;

  let fields = {
    Nome: wallet.name,
    Descrição: wallet.description,
    Moeda: `${wallet.currency.name} (${wallet.currency.representation})`,
  };
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
      <Button color="red">
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
{/if}
