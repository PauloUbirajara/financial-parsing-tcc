<script lang="ts">
  import { navigating, page } from "$app/stores";
  import {
    Button,
    Span,
    ButtonGroup,
    Dropdown,
    DropdownItem,
    Input,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Modal,
    TextPlaceholder,
  } from "flowbite-svelte";
  import type { Wallet } from "../../../domain/models/wallet";
  import {
    DotsHorizontalOutline,
    TrashBinSolid,
    PlusOutline,
    ExclamationCircleOutline,
  } from "flowbite-svelte-icons";
  import { deserialize } from "$app/forms";

  export let onAdd: Function;
  export let onDelete: Function;

  let showDeleteWalletModal = false;

  function selectWallet(w: Wallet) {
    selectedWallet = selectedWallet === w ? null : w;
  }

  async function deleteWallet() {
    if (selectedWallet === null) {
      console.warn("Cannot delete null wallet");
      return;
    }

    const response = await fetch(`/api/wallets/${selectedWallet?.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    const data = deserialize(await response.text());
    console.log({ response, data });
  }

  let wallets = $page.data.wallets as Wallet[];
  let selectedWallet: Wallet | null;
  let searchTerm: string = "";
  $: filteredItems = wallets.filter(
    (item) => item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1,
  );
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
        <TableHeadCell>Nome</TableHeadCell>
        <TableHeadCell>Descrição</TableHeadCell>
        <TableHeadCell>Moeda</TableHeadCell>
        <TableHeadCell>Ações</TableHeadCell>
      </TableHead>
      <TableBody>
        {#each filteredItems as item}
          <TableBodyRow>
            <TableBodyCell>{item.name}</TableBodyCell>
            <TableBodyCell>{item.description}</TableBodyCell>
            <TableBodyCell>{item.currency.representation}</TableBodyCell>
            <TableBodyCell>
              <Button
                class="!p-2 dots-menu"
                color="alternative"
                on:click={() => selectWallet(item)}
              >
                <DotsHorizontalOutline />
              </Button>
            </TableBodyCell>
          </TableBodyRow>
        {/each}
      </TableBody>
    </Table>
  </div>

  <Dropdown triggeredBy=".dots-menu">
    <DropdownItem href={`/api/wallets/${selectedWallet?.id}`}>Ver</DropdownItem>
    <DropdownItem href={`/api/wallets/${selectedWallet?.id}/edit`}>
      Editar
    </DropdownItem>
    <DropdownItem slot="footer" on:click={() => (showDeleteWalletModal = true)}
      >Apagar</DropdownItem
    >
  </Dropdown>

  <Modal open={showDeleteWalletModal} size="xs" autoclose>
    <div class="text-center">
      <ExclamationCircleOutline
        class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
      />
      <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
        Deseja apagar a carteira "<Span>{selectedWallet?.name}</Span>"?
      </h3>
      <Button color="red" class="me-2" on:click={() => deleteWallet()}
        >Sim, apagar</Button
      >
      <Button color="alternative">Não, cancelar</Button>
    </div>
  </Modal>
{/if}
