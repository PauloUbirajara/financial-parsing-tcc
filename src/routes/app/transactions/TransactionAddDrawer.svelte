<script lang="ts">
  import {
    Button,
    CloseButton,
    Drawer,
    Input,
    Label,
    MultiSelect,
    Select,
    Textarea,
  } from "flowbite-svelte";
  import { InfoCircleSolid, PlusOutline } from "flowbite-svelte-icons";
  import { sineIn } from "svelte/easing";
  import type { Wallet } from "../../../domain/models/wallet";
  import type { Category } from "../../../domain/models/category";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";
  import { ToastType } from "../../../domain/models/toastMessage";

  export let hideAddDrawer: boolean = true;
  export let walletResponse: GetAllModelsRepositoryResponse;
  export let categoryResponse: GetAllModelsRepositoryResponse;
  let wallets = walletResponse.results as Wallet[];
  let categories = categoryResponse.results as Category[];

  let transitionParams = {
    x: 320,
    duration: 200,
    easing: sineIn,
  };

  let transaction = {
    name: "",
    description: "",
    categories: [],
    wallet: "",
    value: 0.0,
    transaction_date: "",
  };

  async function onAdd() {
    const response = await fetch("?/create", {
      method: "POST",
      body: JSON.stringify(transaction),
    });
    if (response.ok) {
      showToast({
        title: "Adicionar transação",
        message: `Transação "${transaction.name}" adicionada com sucesso.`,
        type: ToastType.SUCCESS,
      });
      goto("/app/transactions", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Adicionar transação",
      message: `Houve um erro ao adicionar a transação "${transaction.name}".`,
      type: ToastType.ERROR,
    });
  }
</script>

<Drawer
  placement="right"
  transitionType="fly"
  {transitionParams}
  bind:hidden={hideAddDrawer}
  id="sidebar4"
>
  <div class="flex items-center">
    <h5
      id="drawer-label"
      class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
    >
      <InfoCircleSolid class="w-5 h-5 me-2.5" />
      Nova Transação
    </h5>
    <CloseButton
      on:click={() => (hideAddDrawer = true)}
      class="mb-4 dark:text-white"
    />
  </div>
  <form on:submit|preventDefault={onAdd} class="mb-6">
    <div class="mb-6">
      <Label for="name" class="block mb-2">Nome*</Label>
      <Input
        id="name"
        name="name"
        bind:value={transaction.name}
        required
        placeholder="Digite o nome da transação"
      />
    </div>
    <div class="mb-6">
      <Label for="description" class="mb-2">Descrição</Label>
      <Textarea
        id="description"
        placeholder="Digite a descrição da transação"
        bind:value={transaction.description}
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
        bind:value={transaction.transaction_date}
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
        bind:value={transaction.value}
        required
        placeholder="Digite o valor da transação"
      />
    </div>
    <div class="mb-6">
      <Label for="wallet" class="mb-2">Carteira*</Label>
      <Select
        id="wallet"
        placeholder="Selecione uma carteira"
        bind:value={transaction.wallet}
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
        dropdownClass="inline-flex h-100 max-h-[200px]"
        id="wallet"
        placeholder="Selecione uma ou mais categorias"
        name="categories"
        items={categories.map((c) => ({ value: c.id, name: c.name }))}
        bind:value={transaction.categories}
      />
    </div>

    <Button type="submit" class="w-full">
      Salvar transação
    </Button>
  </form>
</Drawer>
