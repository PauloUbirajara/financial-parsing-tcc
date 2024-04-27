<script lang="ts">
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";

  import { ListPlaceholder } from "flowbite-svelte";
  import { onMount } from "svelte";

  import ModelList from "../../../components/model/ModelList.svelte";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";
  import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";
  import WalletAddDrawer from "./WalletAddDrawer.svelte";
  import walletModelListInfo from "../../../data/usecases/modelListInfo/wallet";
  import walletSerializer from "../../../data/usecases/modelSerializer/wallet";
  import { ToastType } from "../../../domain/models/toastMessage";
  import { page } from "$app/stores";
  import UserNavbar from "../../../components/UserNavbar.svelte";

  let hideAddDrawer = true;
  function onAdd() {
    hideAddDrawer = false;
  }

  let loading = true;
  let selectedModel: Record<string, any> | null = null;
  onMount(() => {
    loading = false;
  });

  async function onDelete() {
    if (selectedModel === null) {
      console.warn("Cannot delete null wallet!");
      return;
    }

    try {
      const response = await fetch(
        `/api/wallets/${selectedModel?.id}?/delete`,
        {
          method: "POST",
          body: JSON.stringify({}),
        },
      );

      if (response.ok) {
        goto($page.url.toString(), { invalidateAll: true });
        showToast({
          title: "Remover Carteira",
          message: `Carteira "${selectedModel.name}" removida com sucesso`,
          type: ToastType.SUCCESS,
        });
        return;
      }
    } catch (e) {
      showToast({
        title: "Remover Carteira",
        message: `Houve um erro ao remover a carteira "${selectedModel.name}"`,
        type: ToastType.ERROR,
      });
    }
  }

  // model list
  let walletResponse: GetAllModelsRepositoryResponse =
    $page.data.walletResponse;
  let currencyResponse: GetAllModelsRepositoryResponse =
    $page.data.currencyResponse;

  let serializer: IModelSerializer = walletSerializer;
  let modelListInfo: IModelListInfo = walletModelListInfo;

  const breadcrumbs = [{ label: "Carteiras", href: "/api/wallets" }];

  async function onBulkDelete(ids: string[]) {
    const response = await fetch("/api/wallets?/bulk-delete", {
      method: "POST",
      body: JSON.stringify(ids),
    });

    if (response.ok) {
      showToast({
        title: "Remover carteiras",
        message: "Carteiras removidas com sucesso!",
        type: ToastType.SUCCESS,
      });
      goto("/api/wallets", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Remover carteiras",
      message: "Não foi possível remover as carteiras selecionadas.",
      type: ToastType.ERROR,
    });
  }

  async function updateResults(search: string) {
    const urlParams = new URLSearchParams({ search });
    let url = `/api/wallets?${urlParams.toString()}`;
    goto(url, { invalidateAll: true });
  }

  function removeQuery() {
    goto("/api/wallets", { invalidateAll: true });
  }
</script>

{#if loading}
  <div class="container mx-auto">
    <ListPlaceholder
      divClass="p-4 space-y-4 w-full rounded border border-gray-200 divide-y divide-gray-200 shadow animate-pulse dark:divide-gray-700 md:p-6 dark:border-gray-700"
    />
  </div>
{:else}
  <UserNavbar {breadcrumbs} />
  <WalletAddDrawer bind:hideAddDrawer {currencyResponse} />
  <ModelList
    bind:selectedModel
    {updateResults}
    {removeQuery}
    {onAdd}
    {onDelete}
    {serializer}
    {modelListInfo}
    {onBulkDelete}
    response={walletResponse}
  />
{/if}
