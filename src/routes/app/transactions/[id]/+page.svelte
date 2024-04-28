<script lang="ts">
  import { page } from "$app/stores";
  import UserNavbar from "../../../../components/UserNavbar.svelte";
  import TransactionDetail from "./TransactionDetail.svelte";

  const transaction = $page.data.transaction;
  const breadcrumbs = [
    { label: "Transações", href: "/app/transactions" },
    { label: transaction.name, href: `/app/transactions/${transaction.id}` },
  ];
</script>

<UserNavbar {breadcrumbs} />
<div class="container mx-auto grid gap-8 h-min">
  <TransactionDetail {transaction} />
</div>
