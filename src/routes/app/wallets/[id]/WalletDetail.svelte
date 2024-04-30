<script lang="ts">
  import { navigating } from "$app/stores";
  import {
    Button,
    Dropdown,
    Spinner,
    DropdownItem,
    P,
    A,
    TextPlaceholder,
  } from "flowbite-svelte";
  import {
    ArrowLeftOutline,
    EditSolid,
    FileExportSolid,
    TrashBinSolid,
    DownloadSolid,
  } from "flowbite-svelte-icons";
  import { goto } from "$app/navigation";
  import DeleteModelModal from "../../../../components/model/DeleteModelModal.svelte";
  import walletSerializer from "../../../../data/usecases/modelSerializer/wallet";
  import type { IModelSerializer } from "../../../../domain/usecases/modelSerializer";
  import type { Wallet } from "../../../../domain/models/wallet";
  import { showToast } from "$lib/toast";
  import { ToastType } from "../../../../domain/models/toastMessage";
  import { deserialize } from "$app/forms";
  import { fly } from "svelte/transition";

  function onEdit() {
    goto(`/app/wallets/${wallet.id}/edit`);
  }

  export let wallet: Wallet;
  let showDeleteWalletModal: boolean = false;

  let serializer: IModelSerializer = walletSerializer;
  let lastExportBlob: { blob: Blob | null; loading: boolean } = {
    blob: null,
    loading: false,
  };

  async function onDelete() {
    const response = await fetch(`/app/wallets/${wallet.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (response.ok) {
      goto("/app/wallets", { invalidateAll: true });
    }
  }

  async function onExport(format: string) {
    lastExportBlob.blob = null;
    lastExportBlob.loading = false;

    try {
      const response = await fetch(`/app/wallets/${wallet.id}?/export`, {
        method: "POST",
        body: JSON.stringify({ format }),
      });
      if (!response.ok) {
        throw new Error("Erro ao solicitar dados para exportar");
      }
      const serverResponse = deserialize(await response.text());
      if (
        serverResponse.type !== "success" ||
        serverResponse.data === undefined
      )
        throw new Error(
          "Resultado inesperado ao solicitar dados para exportar",
        );

      const url = serverResponse.data["url"] as string;
      const body = serverResponse.data["body"] as string;
      const headers = serverResponse.data["headers"] as Record<any, any>;

      lastExportBlob.loading = true;
      const fileResponse = await fetch(url, { headers, body, method: "POST" });
      if (!fileResponse.ok) {
        throw new Error("Erro ao usar dados solicitados para exportar");
      }
      const blob = await fileResponse.blob();
      lastExportBlob.blob = blob;

      showToast({
        title: "Exportar carteira",
        message: `Exportação da carteira para o formato "${format}" realizada com sucesso!`,
        type: ToastType.SUCCESS,
      });
    } catch (e) {
      console.warn("Could not export", e);
      showToast({
        title: "Exportar carteira",
        message: `Não foi possível abrir resultado da exportação para o formato "${format}"`,
        type: ToastType.WARNING,
      });
      return;
    } finally {
      lastExportBlob.loading = false;
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
      href="/app/wallets"
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <div class="flex gap-4 items-center">
      {#if lastExportBlob.blob && !lastExportBlob.loading}
        <div transition:fly>
          <A
            color="alternative"
            href={URL.createObjectURL(lastExportBlob.blob)}
            target="_blank"
          >
            <DownloadSolid class="me-2" />
            Abrir exportação
          </A>
        </div>
      {/if}
      <Button
        color="alternative"
        class="export-btn"
        disabled={lastExportBlob.loading}
      >
        {#if lastExportBlob.loading}
          <Spinner size="4" />
        {:else}
          <FileExportSolid class="me-2" />
          Exportar
        {/if}
      </Button>
      <Dropdown triggeredBy=".export-btn">
        <DropdownItem on:click={() => onExport("pdf")}>PDF</DropdownItem>
        <DropdownItem on:click={() => onExport("csv")}>CSV</DropdownItem>
        <DropdownItem on:click={() => onExport("html")}>HTML</DropdownItem>
      </Dropdown>
      <Button on:click={() => onEdit()}>
        <EditSolid class="me-2" />
        Editar
      </Button>
      <Button color="red" on:click={() => (showDeleteWalletModal = true)}>
        <TrashBinSolid class="me-2" />
        Apagar
      </Button>
    </div>
  </div>
  {#each Object.entries(serializer.serialize(wallet)) as [key, value]}
    <div class="content">
      <P>{key}</P>
      <P size="3xl" weight="semibold">{value}</P>
    </div>
  {/each}

  <DeleteModelModal
    bind:showDeleteModal={showDeleteWalletModal}
    title={`Deseja apagar a carteira "${wallet.name}"?`}
    {onDelete}
  />
{/if}
