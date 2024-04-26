<script lang="ts">
  import {
    Button,
    CloseButton,
    Drawer,
    Input,
    Label,
    Select,
    Textarea,
  } from "flowbite-svelte";
  import { InfoCircleSolid, PlusOutline } from "flowbite-svelte-icons";
  import { sineIn } from "svelte/easing";
  import type { Currency } from "../../../domain/models/currency";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";
  import { ToastType } from "../../../domain/models/toastMessage";

  export let hideAddDrawer: boolean = true;
  export let currencyResponse: GetAllModelsRepositoryResponse;
  let currencies = currencyResponse.results as Currency[];

  let transitionParams = {
    x: 320,
    duration: 200,
    easing: sineIn,
  };

  let wallet = {
    name: "",
    description: "",
    currency: "",
  };

  async function onAdd() {
    const response = await fetch("/api/wallets?/create", {
      method: "POST",
      body: JSON.stringify(wallet),
    });
    if (response.ok) {
      showToast({
        title: "Adicionar carteira",
        message: `Carteira "${wallet.name}" adicionada com sucesso.`,
        type: ToastType.SUCCESS,
      });
      goto("/api/wallets", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Adicionar carteira",
      message: `Houve um erro ao adicionar a carteira "${wallet.name}".`,
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
      Nova Carteira
    </h5>
    <CloseButton
      on:click={() => (hideAddDrawer = true)}
      class="mb-4 dark:text-white"
    />
  </div>
  <form class="mb-6" on:submit|preventDefault={onAdd}>
    <div class="mb-6">
      <Label for="name" class="block mb-2">Nome*</Label>
      <Input
        id="name"
        name="name"
        required
        bind:value={wallet.name}
        placeholder="Digite o nome da carteira"
      />
    </div>
    <div class="mb-6">
      <Label for="description" class="mb-2">Descrição</Label>
      <Textarea
        id="description"
        placeholder="Digite a descrição da carteira"
        rows="4"
        bind:value={wallet.description}
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
        bind:value={wallet.currency}
        required
      >
        {#each currencies as item}
          <option value={item.id}>{item.representation} - {item.name}</option>
        {/each}
      </Select>
    </div>
    <Button type="submit" class="w-full">
      <PlusOutline class="w-3.5 h-3.5 me-2.5" />
      Criar carteira
    </Button>
  </form>
</Drawer>
