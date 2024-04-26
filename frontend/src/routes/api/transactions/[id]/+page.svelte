<script lang="ts">
  import { page } from "$app/stores";
  import UserNavbar from "../../../../components/UserNavbar.svelte";
  import TransactionDetail from "./TransactionDetail.svelte";

  const transaction = $page.data.transaction;
  const breadcrumbs = [
    { label: "Transações", href: "/api/transactions" },
    { label: transaction.name, href: `/api/transactions/${transaction.id}` },
  ];
</script>

<UserNavbar {breadcrumbs} />
<div class="container mx-auto grid gap-4">
  <TransactionDetail {transaction} />
</div>
