<script lang="ts">
  import { navigating, page } from "$app/stores";
  import {
    TextPlaceholder,
    ButtonGroup,
    Button,
    Table,
    TableHead,
    TableHeadCell,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    Dropdown,
    DropdownItem,
    Checkbox,
    P,
    Search,
  } from "flowbite-svelte";
  import { getFilteredUrlSearchParams } from "../../helpers/url";
  import {
    DotsHorizontalOutline,
    PlusOutline,
    SearchOutline,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import DeleteModelModal from "./DeleteModelModal.svelte";
  import type { IModelListInfo } from "../../domain/usecases/modelListInfo";
  import type { IModelSerializer } from "../../domain/usecases/modelSerializer";
  import type { GetAllModelsRepositoryResponse } from "../../domain/models/modelRepositoryDto";
  import { goto } from "$app/navigation";
  import ModelListPagination from "./ModelListPagination.svelte";

  // List actions
  export let onAdd: Function;
  export let onDelete: Function;

  // List input
  export let serializer: IModelSerializer;
  export let response: GetAllModelsRepositoryResponse;
  export let selectedModel: Record<any, any> | null;
  let searchTerm = "";

  $: urlQuery = getFilteredUrlSearchParams({
    page: $page.url.searchParams.get("page") || null,
    search: searchTerm,
  });

  // List
  async function updateResults() {
    let url = "/api/wallets";

    url = `/api/wallets?${urlQuery.toString()}`;
    goto(url, { invalidateAll: true });
  }

  let currentQuery = $page.url.searchParams.get("search");

  // Actions
  export let modelListInfo: IModelListInfo;

  // Deletion modal
  let showDeleteModelModal: boolean = false;
</script>

{#if $navigating}
  <TextPlaceholder
    divClass="space-y-2.5 animate-pulse w-full mx-auto container"
  />
{:else}
  <div class="container mx-auto flex flex-col gap-4">
    {#if currentQuery}
      <P class="w-full">
        Você buscou por: {currentQuery}
      </P>
    {/if}
    <div class="flex items-center justify-between">
      <form
        class="flex gap-2 items-center min-w-80"
        on:submit|preventDefault={updateResults}
      >
        <Search
          size="md"
          class="w-full"
          placeholder="Buscar carteira por nome"
          bind:value={searchTerm}
        />
        <Button type="submit" class="!p-2">
          <SearchOutline class="w-6 h-6" />
        </Button>
      </form>
      <ButtonGroup>
        <Button color="primary" on:click={() => onAdd()}>
          <PlusOutline class="w-4 h-4 me-2" />
          Adicionar
        </Button>
        <Button disabled color="red" class="!p-2" on:click={() => onDelete()}>
          <TrashBinSolid class="w-4 h-4 me-2" />
          Deletar
        </Button>
      </ButtonGroup>
    </div>
    {#if !response.count}
      <P align="center">Sem resultados.</P>
    {:else}
      <Table divClass="relative overflow-x-auto rounded dark" hoverable={true}>
        <TableHead>
          <TableHeadCell class="!p-4">
            <Checkbox />
          </TableHeadCell>
          {#each serializer.getFields() as key}
            <TableHeadCell>{key}</TableHeadCell>
          {/each}
          <TableHeadCell>Ações</TableHeadCell>
        </TableHead>
        <TableBody>
          {#each response.results as item}
            <TableBodyRow>
              <TableBodyCell class="!p-4">
                <Checkbox />
              </TableBodyCell>
              {#each Object.values(serializer.serialize(item)) as val}
                <TableBodyCell>
                  {val}
                </TableBodyCell>
              {/each}
              <TableBodyCell>
                <Button
                  class="!p-2 dots-menu"
                  color="alternative"
                  on:click={() => {
                    selectedModel = item;
                  }}
                >
                  <DotsHorizontalOutline />
                </Button>
              </TableBodyCell>
            </TableBodyRow>
          {/each}
        </TableBody>
      </Table>
    {/if}

    <ModelListPagination {response} />
  </div>

  <Dropdown triggeredBy=".dots-menu">
    {#if selectedModel !== null}
      <DropdownItem href={modelListInfo.getDetailUrl(selectedModel)}>
        Ver
      </DropdownItem>
      <DropdownItem href={modelListInfo.getEditUrl(selectedModel)}>
        Editar
      </DropdownItem>
      <DropdownItem
        slot="footer"
        on:click={() => (showDeleteModelModal = true)}
      >
        Apagar
      </DropdownItem>
    {/if}
  </Dropdown>

  {#if selectedModel !== null}
    <DeleteModelModal
      bind:showDeleteModal={showDeleteModelModal}
      title={modelListInfo.getDeleteModalTitle(selectedModel)}
      {onDelete}
    />
  {/if}
{/if}
