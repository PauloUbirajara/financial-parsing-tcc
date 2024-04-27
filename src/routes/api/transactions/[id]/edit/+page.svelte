<script lang="ts">
  import { navigating, page } from "$app/stores";
  import {
    TextPlaceholder,
    Button,
    Input,
    Label,
    Textarea,
    Select,
    MultiSelect,
  } from "flowbite-svelte";
  import { ArrowLeftOutline, CheckOutline } from "flowbite-svelte-icons";
  import type {
    GetAllModelsRepositoryResponse,
    GetModelByIdRepositoryResponse,
  } from "../../../../../domain/models/modelRepositoryDto";
  import { ToastType } from "../../../../../domain/models/toastMessage";
  import { showToast } from "$lib/toast";
  import { goto } from "$app/navigation";
  import UserNavbar from "../../../../../components/UserNavbar.svelte";

  let transactionResponse: GetModelByIdRepositoryResponse =
    $page.data.transactionResponse;
  let walletResponse: GetAllModelsRepositoryResponse =
    $page.data.walletResponse;
  let categoryResponse: GetAllModelsRepositoryResponse =
    $page.data.categoryResponse;

  let transaction: Record<string, any> = transactionResponse as Record<
    string,
    any
  >;
  let wallets = walletResponse.results;
  let categories = categoryResponse.results;

  let updated = {
    name: transaction.name,
    wallet: transaction.wallet.id,
    description: transaction.description || "",
    value: transaction.value,
    transaction_date: transaction.transaction_date,
    categories: transaction.categories.map((c: any) => c.id),
  };

  const breadcrumbs = [
    { label: "Transações", href: "/api/transactions" },
    { label: transaction.name, href: `/api/transactions/${transaction.id}` },
    { label: "Editar", href: `/api/transactions/${transaction.id}/edit` },
  ];

  async function onUpdate() {
    const id = $page.params.id;
    const response = await fetch(`/api/transactions/${id}/edit`, {
      method: "POST",
      body: JSON.stringify(updated),
    });

    if (response.ok) {
      showToast({
        title: "Atualizar transação",
        message: `Transação "${updated.name}" atualizada com sucesso!`,
        type: ToastType.SUCCESS,
      });
      goto($page.url, { invalidateAll: true });
      return;
    }
    showToast({
      title: "Atualizar transação",
      message: `Houve um erro ao atualizar a transação "${updated.name}".`,
      type: ToastType.ERROR,
    });
  }
</script>

<UserNavbar {breadcrumbs} />
<div class="container mx-auto flex flex-col gap-4">
  {#if $navigating}
    <TextPlaceholder divClass="space-y-2.5 animate-pulse mx-auto w-full" />
  {:else}
    <div class="actions flex justify-between">
      <Button
        outline={true}
        class="!p-2"
        size="md"
        color="primary"
        href={`/api/transactions/${transaction.id}`}
      >
        <ArrowLeftOutline class="me-2" />
        Detalhes
      </Button>
    </div>

    <form on:submit|preventDefault={onUpdate}>
      <div class="mb-6">
        <Label for="name" class="block mb-2">Nome*</Label>
        <Input
          id="name"
          name="name"
          bind:value={updated.name}
          required
          placeholder="Digite o nome da transação"
        />
      </div>
      <div class="mb-6">
        <Label for="description" class="mb-2">Descrição</Label>
        <Textarea
          id="description"
          placeholder="Digite a descrição da transação"
          bind:value={updated.description}
          rows="4"
          name="description"
        />
      </div>
      <div class="mb-6">
        <Label for="transaction_date" class="mb-2">Data da transação*</Label>
        <Input
          type="date"
          id="transaction_date"
          placeholder="Selecione a data da transação"
          bind:value={updated.transaction_date}
          name="transaction_date"
          required
        />
      </div>
      <div class="mb-6">
        <Label for="value" class="block mb-2">Valor*</Label>
        <Input
          type="number"
          step="0.01"
          id="value"
          name="value"
          bind:value={updated.value}
          required
          placeholder="Digite o valor da transação"
        />
      </div>
      <div class="mb-6">
        <Label for="wallet" class="mb-2">Carteira*</Label>
        <Select
          id="wallet"
          placeholder="Selecione uma carteira"
          bind:value={updated.wallet}
          rows="4"
          name="wallet"
          required
        >
          {#each wallets as item}
            <option value={item.id}>
              {item.name} ({item.currency.representation})
            </option>
          {/each}
        </Select>
      </div>
      <div class="mb-6">
        <Label for="categories" class="mb-2">Categorias</Label>
        <MultiSelect
          id="categories"
          placeholder="Selecione uma ou mais categorias"
          name="categories"
          items={categories.map((c) => ({ value: c.id, name: c.name }))}
          bind:value={updated.categories}
        />
      </div>

      <Button type="submit" class="w-full" color="green">
        <CheckOutline class="mr-2" />
        Salvar
      </Button>
    </form>
  {/if}
</div>
