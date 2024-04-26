<script lang="ts">
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";

  import { ListPlaceholder } from "flowbite-svelte";
  import { onMount } from "svelte";

  import ModelList from "../../../components/model/ModelList.svelte";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";
  import categoryModelListInfo from "../../../data/usecases/modelListInfo/category";
  import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";
  import categorySerializer from "../../../data/usecases/modelSerializer/category";
  import CategoryAddDrawer from "./CategoryAddDrawer.svelte";
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
      console.warn("Cannot delete null category!");
      return;
    }

    try {
      const response = await fetch(
        `/api/categories/${selectedModel?.id}?/delete`,
        {
          method: "POST",
          body: JSON.stringify({}),
        },
      );

      if (response.ok) {
        goto($page.url.toString(), { invalidateAll: true });
        showToast({
          title: "Remover Categoria",
          message: `Categoria "${selectedModel.name}" removida com sucesso`,
          type: ToastType.SUCCESS,
        });
        return;
      }
    } catch (e) {
      showToast({
        title: "Remover Categoria",
        message: `Houve um erro ao remover a categoria "${selectedModel.name}"`,
        type: ToastType.ERROR,
      });
    }
  }

  // model list
  let categoryResponse: GetAllModelsRepositoryResponse =
    $page.data.categoryResponse;

  let serializer: IModelSerializer = categorySerializer;
  let modelListInfo: IModelListInfo = categoryModelListInfo;

  const breadcrumbs = [{ label: "Categorias", href: "/api/categories" }];

  async function onBulkDelete(ids: string[]) {
    const response = await fetch("/api/categories?/bulk-delete", {
      method: "POST",
      body: JSON.stringify(ids),
    });

    if (response.ok) {
      showToast({
        title: "Remover categorias",
        message: "Categorias removidas com sucesso!",
        type: ToastType.SUCCESS,
      });
      goto("/api/categories", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Remover categorias",
      message: "Não foi possível remover as categorias selecionadas.",
      type: ToastType.ERROR,
    });
  }

  async function updateResults(search: string) {
    const urlParams = new URLSearchParams({ search });
    let url = `/api/categories?${urlParams.toString()}`;
    goto(url, { invalidateAll: true });
  }

  function removeQuery() {
    goto("/api/categories", { invalidateAll: true });
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
  <CategoryAddDrawer bind:hideAddDrawer />
  <ModelList
    bind:selectedModel
    {updateResults}
    {removeQuery}
    {onAdd}
    {onDelete}
    {serializer}
    {modelListInfo}
    {onBulkDelete}
    response={categoryResponse}
  />
{/if}
