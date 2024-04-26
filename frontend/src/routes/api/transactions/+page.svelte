<script lang="ts">
  import { goto } from "$app/navigation";
  import { ListPlaceholder } from "flowbite-svelte";
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import { showToast } from "$lib/toast";

  import ModelList from "../../../components/model/ModelList.svelte";
  import TransactionAddDrawer from "./TransactionAddDrawer.svelte";
  import transactionModelListInfo from "../../../data/usecases/modelListInfo/transaction";
  import transactionserializer from "../../../data/usecases/modelSerializer/transaction";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";
  import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";
  import { ToastType } from "../../../domain/models/toastMessage";
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
      console.warn("Cannot delete null transaction!");
      return;
    }

    try {
      const response = await fetch(
        `/api/transactions/${selectedModel?.id}?/delete`,
        {
          method: "POST",
          body: JSON.stringify({}),
        },
      );

      if (response.ok) {
        goto($page.url.toString(), { invalidateAll: true });
        showToast({
          title: "Remover Transação",
          message: `Transação "${selectedModel.name}" removida com sucesso`,
          type: ToastType.SUCCESS,
        });
        return;
      }
    } catch (e) {
      showToast({
        title: "Remover Transação",
        message: `Houve um erro ao remover a transação "${selectedModel.name}"`,
        type: ToastType.ERROR,
      });
    }
  }

  // model list
  let transactionResponse: GetAllModelsRepositoryResponse =
    $page.data.transactionResponse;
  let walletResponse: GetAllModelsRepositoryResponse =
    $page.data.walletResponse;
  let categoryResponse: GetAllModelsRepositoryResponse =
    $page.data.categoryResponse;

  let serializer: IModelSerializer = transactionserializer;
  let modelListInfo: IModelListInfo = transactionModelListInfo;

  const breadcrumbs = [{ label: "Transações", href: "/api/transactions" }];

  async function onBulkDelete(ids: string[]) {
    const response = await fetch("/api/transactions?/bulk-delete", {
      method: "POST",
      body: JSON.stringify(ids),
    });

    if (response.ok) {
      showToast({
        title: "Remover transações",
        message: "Transações removidas com sucesso!",
        type: ToastType.SUCCESS,
      });
      goto("/api/transactions", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Remover transações",
      message: "Não foi possível remover as transações selecionadas.",
      type: ToastType.ERROR,
    });
  }

  async function updateResults(search: string) {
    const urlParams = new URLSearchParams({ search });
    let url = `/api/transactions?${urlParams.toString()}`;
    goto(url, { invalidateAll: true });
  }

  function removeQuery() {
    goto("/api/transactions", { invalidateAll: true });
  }
</script>

<UserNavbar {breadcrumbs} />
{#if loading}
  <div class="container mx-auto">
    <ListPlaceholder
      divClass="p-4 space-y-4 w-full rounded border border-gray-200 divide-y divide-gray-200 shadow animate-pulse dark:divide-gray-700 md:p-6 dark:border-gray-700"
    />
  </div>
{:else}
  <TransactionAddDrawer
    bind:hideAddDrawer
    {walletResponse}
    {categoryResponse}
  />
  <ModelList
    bind:selectedModel
    {updateResults}
    {removeQuery}
    {onAdd}
    {onDelete}
    {serializer}
    {modelListInfo}
    {onBulkDelete}
    response={transactionResponse}
  />
{/if}
