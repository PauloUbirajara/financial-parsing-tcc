<script lang="ts">
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";

  import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";
  import type { Wallet } from "../../../domain/models/wallet";
  import ModelList from "../../../components/model/ModelList.svelte";
  import walletModelListInfo from "../../../data/usecases/modelListInfo/wallet";

  export let onAdd: Function;

  async function onDelete() {
    if (selectedWallet === null) {
      console.warn("Cannot delete null wallet");
      return;
    }

    const response = await fetch(`/api/wallets/${selectedWallet?.id}?/delete`, {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (response.ok) {
      goto("/api/wallets", { invalidateAll: true });
    }
  }

  // $: wallets = $page.data.wallets.results;
  let wallets = $page.data.wallets.results;
  let selectedWallet: Wallet | null;
  let searchTerm: string = "";
  $: filteredItems = wallets.filter(
    (item: Wallet) =>
      item.name.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1,
  );
  let fields = {
    Nome: (w: Wallet) => w.name,
    Descrição: (w: Wallet) => w.description,
    Moeda: (w: Wallet) => w.currency.representation,
  };
  let modelListInfo: IModelListInfo = walletModelListInfo;
</script>

<ModelList
  {onAdd}
  {onDelete}
  {searchTerm}
  {fields}
  {filteredItems}
  {modelListInfo}
  selectedModel={selectedWallet}
/>
