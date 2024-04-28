<script lang="ts">
  import { page } from "$app/stores";
  import UserNavbar from "../../../../components/UserNavbar.svelte";
  import WalletDetail from "./WalletDetail.svelte";

  const wallet = $page.data.wallet;
  const breadcrumbs = [
    { label: "Carteiras", href: "/app/wallets" },
    { label: wallet.name, href: `/app/wallets/${wallet.id}` },
  ];
</script>

<UserNavbar {breadcrumbs} />
<div class="container mx-auto grid gap-8 h-min">
  <WalletDetail {wallet} />
</div>
