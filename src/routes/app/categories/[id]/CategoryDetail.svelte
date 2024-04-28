<script lang="ts">
  import { navigating } from "$app/stores";
  import { Button, P, TextPlaceholder } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    EditSolid,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../../components/model/DeleteModelModal.svelte";
  import categorySerializer from "../../../../data/usecases/modelSerializer/category";
  import type { IModelSerializer } from "../../../../domain/usecases/modelSerializer";
  import type { Category } from "../../../../domain/models/category";

  function onEdit() {
    goto(`/app/categories/${category.id}/edit`);
  }

  export let category: Category;
  let showDeleteCategoryModal: boolean = false;

  let serializer: IModelSerializer = categorySerializer;

  async function onDelete() {
    const response = await fetch(`/app/categories/${category.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (response.ok) {
      goto("/app/categories", { invalidateAll: true });
    }
  }
</script>

{#if $navigating}
  <TextPlaceholder divClass="space-y-2.5 animate-pulse mx-auto w-full" />
{:else}
  <div class="actions flex justify-between">
    <Button
      outline={true}
      class="!p-2"
      size="lg"
      color="primary"
      href="/app/categories"
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <div class="flex gap-4 items-center">
      <Button on:click={() => onEdit()}>
        <EditSolid />
        Editar
      </Button>
      <Button color="red" on:click={() => (showDeleteCategoryModal = true)}>
        <TrashBinSolid />
        Apagar
      </Button>
    </div>
  </div>
  {#each Object.entries(serializer.serialize(category)) as [key, value]}
    <div class="content">
      <P>{key}</P>
      <P size="3xl" weight="semibold">{value}</P>
    </div>
  {/each}

  <DeleteModelModal
    bind:showDeleteModal={showDeleteCategoryModal}
    title={`Deseja apagar a categoria "${category.name}"?`}
    {onDelete}
  />
{/if}
