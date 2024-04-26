<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { TextPlaceholder, Button, Input, Label } from "flowbite-svelte";
  import { ArrowLeftOutline, CheckOutline } from "flowbite-svelte-icons";
  import type { Category } from "../../../../../domain/models/category";
  import Sidebar from "../../../../../components/Sidebar.svelte";
  import Breadcrumb from "../../../../../components/Breadcrumb.svelte";
  import type { GetModelByIdRepositoryResponse } from "../../../../../domain/models/modelRepositoryDto";
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";
  import { ToastType } from "../../../../../domain/models/toastMessage";

  let categoryResponse: GetModelByIdRepositoryResponse =
    $page.data.categoryResponse;
  let category: Record<string, any> = categoryResponse as Record<string, any>;
  let updated: Category = {
    name: category.name,
  };

  const breadcrumbs = [
    { label: "Categorias", href: "/api/categories" },
    { label: category.name, href: `/api/categories/${category.id}` },
    { label: "Editar", href: `/api/categories/${category.id}/edit` },
  ];

  async function onUpdate() {
    const id = $page.params.id;
    const response = await fetch(`/api/categories/${id}/edit`, {
      method: "POST",
      body: JSON.stringify(updated),
    });

    if (response.ok) {
      showToast({
        title: "Atualizar categoria",
        message: `Categoria "${updated.name}" atualizada com sucesso!`,
        type: ToastType.SUCCESS,
      });
      goto($page.url, { invalidateAll: true });
      return;
    }
    showToast({
      title: "Atualizar categoria",
      message: `Houve um erro ao atualizar a categoria "${updated.name}".`,
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
        href={`/api/categories/${category.id}`}
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
          placeholder="Digite o nome da categoria"
        />
      </div>
      <Button type="submit" class="w-full" color="green">
        <CheckOutline class="mr-2" />
        Salvar
      </Button>
    </form>
  {/if}
</div>
