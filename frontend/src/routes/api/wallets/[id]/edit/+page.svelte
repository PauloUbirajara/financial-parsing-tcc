<script lang="ts">
  import { navigating, page } from "$app/stores";
  import {
    TextPlaceholder,
    Button,
    Input,
    Label,
    Textarea,
    Select,
  } from "flowbite-svelte";
  import { ArrowLeftOutline, CheckOutline } from "flowbite-svelte-icons";
  import type { Wallet } from "../../../../../domain/models/wallet";
  import type { Currency } from "../../../../../domain/models/currency";

  let wallet: Wallet = $page.data.wallet;
  let currencies: Currency[] = $page.data.currencies.results;
  let updated: Wallet = {
    name: wallet.name,
    currency: wallet.currency,
    description: wallet.description || "",
  };
</script>

<div class="container mx-auto flex flex-col gap-4">
  {#if $navigating}
    <TextPlaceholder divClass="space-y-2.5 animate-pulse mx-auto w-full" />
  {:else}
    <div class="actions flex justify-between">
      <Button
        outline={true}
        class="!p-2"
        size="lg"
        color="primary"
        href={`/api/wallets/${wallet.id}`}
      >
        <ArrowLeftOutline class="w-6" />
      </Button>
    </div>

    <form method="POST" action="?">
      <div class="mb-6">
        <Label for="name" class="block mb-2">Nome*</Label>
        <Input
          id="name"
          name="name"
          value={updated.name}
          required
          placeholder="Digite o nome da carteira"
        />
      </div>
      <div class="mb-6">
        <Label for="description" class="mb-2">Descrição</Label>
        <Textarea
          id="description"
          placeholder="Digite a descrição da carteira"
          rows="4"
          value={updated.description}
          name="description"
        />
      </div>
      <div class="mb-6">
        <Label for="currency" class="mb-2">Moeda*</Label>
        <Select
          id="currency"
          placeholder="Selecione uma moeda"
          rows="4"
          name="currency"
          value={updated.currency.id}
          required
        >
          {#each currencies as item}
            <option value={item.id}>{item.representation} - {item.name}</option>
          {/each}
        </Select>
      </div>
      <Button type="submit" class="w-full" color="green">
        <CheckOutline class="mr-2" />
        Salvar
      </Button>
    </form>
  {/if}
</div>
