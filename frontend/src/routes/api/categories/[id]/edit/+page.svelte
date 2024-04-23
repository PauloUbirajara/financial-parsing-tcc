<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { TextPlaceholder, Button, Input, Label } from "flowbite-svelte";
  import { ArrowLeftOutline, CheckOutline } from "flowbite-svelte-icons";
  import type { Category } from "../../../../../domain/models/category";
  import Sidebar from "../../../../../components/Sidebar.svelte";
  import Breadcrumb from "../../../../../components/Breadcrumb.svelte";
  import type { GetModelByIdRepositoryResponse } from "../../../../../domain/models/modelRepositoryDto";

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

    <form method="POST" action="?">
      <div class="mb-6">
        <Label for="name" class="block mb-2">Nome*</Label>
        <Input
          id="name"
          name="name"
          value={updated.name}
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
