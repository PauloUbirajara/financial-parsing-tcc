<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  import ModelList from "../../../components/model/ModelList.svelte";
  import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";
  import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";
  import type { Wallet } from "../../../domain/models/wallet";
  import WalletAddDrawer from "./WalletAddDrawer.svelte";
  import walletModelListInfo from "../../../data/usecases/modelListInfo/wallet";
  import walletSerializer from "../../../data/usecases/modelSerializer/wallet";

  let hideAddDrawer = true;

  function onAdd() {
    hideAddDrawer = false;
  }

  async function onDelete() {
    if (selectedWallet === null) {
      console.warn("Cannot delete null wallet!");
      return;
    }

    const response = await fetch(`/api/wallets/${selectedWallet?.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });

    if (response.ok) {
      goto("?", { invalidateAll: true });
    }
  }

  // model list
  let wallets = $page.data.wallets.results;
  $: filteredItems = wallets.filter((w: Wallet) => w.name.includes(searchTerm));

  let serializer: IModelSerializer = walletSerializer;
  let searchTerm = "";
  let modelListInfo: IModelListInfo = walletModelListInfo;
  let selectedWallet: Wallet | null = null;
</script>

<WalletAddDrawer bind:hideAddDrawer />
<ModelList
  {onAdd}
  {onDelete}
  {serializer}
  {modelListInfo}
  bind:searchTerm
  bind:filteredItems
  bind:selectedModel={selectedWallet}
/>
