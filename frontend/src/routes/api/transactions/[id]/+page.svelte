<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumb from "../../../../components/Breadcrumb.svelte";
  import Sidebar from "../../../../components/Sidebar.svelte";
  import TransactionDetail from "./TransactionDetail.svelte";

  const transaction = $page.data.transaction;
  const breadcrumbs = [
    { label: "Transações", href: "/api/transactions" },
    { label: transaction.name, href: `/api/transactions/${transaction.id}` },
  ];
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>
<div class="container mx-auto grid gap-4">
  <TransactionDetail {transaction} />
</div>
