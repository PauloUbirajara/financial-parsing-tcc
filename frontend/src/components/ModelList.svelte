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
  } from "flowbite-svelte";
  import {
    DotsHorizontalOutline,
    PlusOutline,
    TrashBinSolid,
  } from "flowbite-svelte-icons";

  // List actions
  export let onAdd: Function;
  export let onDelete: Function;

  // List input
  export let selectedModel: Record<any, any> | null;
  export let searchTerm: string;
  export let fields: Record<string, Function>;
  export let filteredItems: any[];
</script>

{#if $navigating}
  <TextPlaceholder
    divClass="space-y-2.5 animate-pulse w-full mx-auto container"
  />
{:else}
  <div class="container mx-auto flex flex-col gap-4">
    <div class="flex items-center justify-between">
      <Input
        class="w-80"
        placeholder="Buscar carteira por nome"
        bind:value={searchTerm}
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
  </div>
{/if}
