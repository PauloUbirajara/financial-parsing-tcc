<script lang="ts">
  import { navigating } from "$app/stores";
  import {
    TextPlaceholder,
    ButtonGroup,
    Button,
    Table,
    TableHead,
    TableHeadCell,
    TableBody,
    Input,
    TableBodyCell,
    TableBodyRow,
    Dropdown,
    DropdownItem,
    Pagination,
  } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    ArrowRightOutline,
    DotsHorizontalOutline,
    PlusOutline,
    TrashBinSolid,
  } from "flowbite-svelte-icons";
  import DeleteModelModal from "./DeleteModelModal.svelte";
  import type { IModelListInfo } from "../../domain/usecases/modelListInfo";

  // List actions
  export let onAdd: Function;
  export let onDelete: Function;

  // List input
  export let selectedModel: Record<any, any> | null;
  export let searchTerm: string;
  export let fields: Record<string, Function>;
  export let filteredItems: any[];

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
    <div class="flex items-center justify-between">
      <Input class="w-80" placeholder="Buscar por nome" bind:value={searchTerm}
      ></Input>
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
    <Table divClass="relative overflow-x-auto rounded dark" hoverable={true}>
      <TableHead>
        {#each Object.keys(fields) as key}
          <TableHeadCell>{key}</TableHeadCell>
        {/each}
        <TableHeadCell>Ações</TableHeadCell>
      </TableHead>
      <TableBody>
        {#each filteredItems as item}
          <TableBodyRow>
            {#each Object.values(fields) as keyFn}
              <TableBodyCell>
                {keyFn(item)}
              </TableBodyCell>
            {/each}
            <TableBodyCell>
              <Button
                class="!p-2 dots-menu"
                color="alternative"
                on:click={() => (selectedModel = item)}
              >
                <DotsHorizontalOutline />
              </Button>
            </TableBodyCell>
          </TableBodyRow>
        {/each}
      </TableBody>
    </Table>

    <div class="flex flex-col items-center justify-center gap-2">
      <div class="text-sm text-gray-700 dark:text-gray-400">
        <span class="font-semibold text-gray-900 dark:text-white">{0}</span>
        ~
        <span class="font-semibold text-gray-900 dark:text-white">{10}</span>
        de
        <span class="font-semibold text-gray-900 dark:text-white">{9999}</span>
        resultados.
      </div>

      <Pagination table large>
        <span slot="prev" class="flex items-center">
          <ArrowLeftOutline class="me-2 w-5 h-5" />
          Anterior
        </span>
        <span slot="next" class="flex items-center">
          Próximo
          <ArrowRightOutline class="ms-2 w-5 h-5" />
        </span>
      </Pagination>
    </div>
  </div>

  <Dropdown triggeredBy=".dots-menu">
    <DropdownItem href={modelListInfo.getDetailUrl(selectedModel)}
      >Ver</DropdownItem
    >
    <DropdownItem href={modelListInfo.getEditUrl(selectedModel)}
      >Editar</DropdownItem
    >
    <DropdownItem slot="footer" on:click={() => (showDeleteModelModal = true)}>
      Apagar
    </DropdownItem>
  </Dropdown>

  {#if selectedModel !== null}
    <DeleteModelModal
      bind:showDeleteModal={showDeleteModelModal}
      title={modelListInfo.getDeleteModalTitle(selectedModel)}
      {onDelete}
    />
  {/if}
{/if}
