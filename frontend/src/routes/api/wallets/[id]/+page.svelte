<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumb from "../../../../components/Breadcrumb.svelte";
  import Sidebar from "../../../../components/Sidebar.svelte";
  import WalletDetail from "./WalletDetail.svelte";

  const wallet = $page.data.wallet;
  const breadcrumbs = [
    { label: "Carteiras", href: "/api/wallets" },
    { label: wallet.name, href: `/api/wallets/${wallet.id}` },
  ];
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>
<div class="container mx-auto grid gap-4">
  <WalletDetail {wallet} />
</div>
