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
    Badge,
    Input,
    Modal,
  } from "flowbite-svelte";
  import {
    DotsHorizontalOutline,
    ExclamationCircleOutline,
    PlusOutline,
    SearchOutline,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import DeleteModelModal from "./DeleteModelModal.svelte";
  import type { IModelListInfo } from "../../domain/usecases/modelListInfo";
  import type { IModelSerializer } from "../../domain/usecases/modelSerializer";
  import type { GetAllModelsRepositoryResponse } from "../../domain/models/modelRepositoryDto";
  import ModelListPagination from "./ModelListPagination.svelte";

  // List actions
  export let onAdd: Function;
  export let onDelete: Function;

  // List input
  export let serializer: IModelSerializer;
  export let response: GetAllModelsRepositoryResponse;
  export let selectedModel: Record<any, any> | null;

  // Bulk delete
  let showBulkDeleteModal: boolean = false;
  let checked: Record<string, boolean> = {};
  let idsToDelete: string[] = [];
  $: {
    idsToDelete = Object.entries(checked)
      .filter((c) => c[1])
      .map((c) => c[0]);
  }

  // Search
  let searchTerm = "";

  // List
  export let updateResults: (search: string) => Promise<void>;

  export let removeQuery: () => void;

  function onBulkCheck() {
    if (idsToDelete.length === response.results.length) {
      checked = {};
      return;
    }
    checked = Object.fromEntries(response.results.map((r) => [r.id, true]));
  }

  export let onBulkDelete: (ids: string[]) => Promise<any>;

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
      <Badge dismissable large color="dark" on:close={() => removeQuery()}>
        Filtrando por "{currentQuery}"
      </Badge>
    {/if}
    <div class="flex items-center justify-between gap-4">
      <form
        class="flex gap-2 items-center min-w-80"
        on:submit|preventDefault={() => updateResults(searchTerm)}
      >
        <Input
          placeholder="Buscar por nome"
          type="text"
          required
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
        <Button
          disabled={idsToDelete.length === 0}
          color="red"
          class="!p-2"
          on:click={() => (showBulkDeleteModal = true)}
        >
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
          <TableHeadCell class="!p-4 w-5">
            <Checkbox
              on:click={() => onBulkCheck()}
              indeterminate={![0, response.results.length].includes(
                idsToDelete.length,
              )}
              checked={idsToDelete.length === response.results.length}
            />
          </TableHeadCell>
          {#each serializer.getFields() as key}
            <TableHeadCell>{key}</TableHeadCell>
          {/each}
          <TableHeadCell class="float-right">Ações</TableHeadCell>
        </TableHead>
        <TableBody>
          {#each response.results as item}
            <TableBodyRow>
              <TableBodyCell class="!p-4">
                <Checkbox bind:checked={checked[item.id]} />
              </TableBodyCell>
              {#each Object.values(serializer.serialize(item)) as val}
                <TableBodyCell class="max-w-20 overflow-x-clip text-ellipsis">
                  {val}
                </TableBodyCell>
              {/each}
              <TableBodyCell class="float-right">
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

    <ModelListPagination
      {response}
      baseUrl={modelListInfo.getListUrl(response.results[0])}
    />
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

  <Modal open={showBulkDeleteModal} size="xs" autoclose dismissable={false}>
    <div class="text-center">
      <ExclamationCircleOutline
        class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
      />
      <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
        Deseja apagar os itens selecionados?
      </h3>
      <Button
        color="red"
        class="me-2"
        on:click={() => onBulkDelete(idsToDelete)}>Sim, apagar</Button
      >
      <Button
        color="alternative"
        on:click={() => (showBulkDeleteModal = false)}
      >
        Não, cancelar
      </Button>
    </div>
  </Modal>

  {#if selectedModel !== null}
    <DeleteModelModal
      bind:showDeleteModal={showDeleteModelModal}
      title={modelListInfo.getDeleteModalTitle(selectedModel)}
      {onDelete}
    />
  {/if}
{/if}
