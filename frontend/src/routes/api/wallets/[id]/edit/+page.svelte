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
  import Sidebar from "../../../../../components/Sidebar.svelte";
  import Breadcrumb from "../../../../../components/Breadcrumb.svelte";
  import type {
    GetAllModelsRepositoryResponse,
    GetModelByIdRepositoryResponse,
  } from "../../../../../domain/models/modelRepositoryDto";
  import { ToastType } from "../../../../../domain/models/toastMessage";
  import { showToast } from "$lib/toast";
  import { goto } from "$app/navigation";

  let walletResponse: GetModelByIdRepositoryResponse =
    $page.data.walletResponse;
  let currencyResponse: GetAllModelsRepositoryResponse =
    $page.data.currencyResponse;
  let wallet: Record<string, any> = walletResponse as Record<string, any>;
  let currencies = currencyResponse.results;
  let updated: Wallet = {
    name: wallet.name,
    currency: wallet.currency.id,
    description: wallet.description || "",
  };

  const breadcrumbs = [
    { label: "Carteiras", href: "/api/wallets" },
    { label: wallet.name, href: `/api/wallets/${wallet.id}` },
    { label: "Editar", href: `/api/wallets/${wallet.id}/edit` },
  ];

  async function onUpdate() {
    const id = $page.params.id;
    const response = await fetch(`/api/wallets/${id}/edit`, {
      method: "POST",
      body: JSON.stringify(updated),
    });

    if (response.ok) {
      showToast({
        title: "Atualizar carteira",
        message: `Carteira "${updated.name}" atualizada com sucesso!`,
        type: ToastType.SUCCESS,
      });
      goto($page.url, { invalidateAll: true });
      return;
    }
    showToast({
      title: "Atualizar carteira",
      message: `Houve um erro ao atualizar a carteira "${updated.name}".`,
      type: ToastType.ERROR,
    });
  }
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>
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
        href={`/api/wallets/${wallet.id}`}
      >
        <ArrowLeftOutline class="me-2" />
        Detalhes
      </Button>
    </div>

    <form method="POST" on:submit|preventDefault={onUpdate}>
      <div class="mb-6">
        <Label for="name" class="block mb-2">Nome*</Label>
        <Input
          id="name"
          name="name"
          bind:value={updated.name}
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
          bind:value={updated.description}
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
          bind:value={updated.currency}
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
