<script lang="ts">
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import TransactionSummary from "./TransactionSummary.svelte";
  import { page } from "$app/stores";
  import CategorySummary from "./CategorySummary.svelte";
  import UserNavbar from "../../../components/UserNavbar.svelte";

  const breadcrumbs = [{ label: "Dashboard", href: "/app/dashboard" }];

  const transactionResponse: GetAllModelsRepositoryResponse =
    $page.data.transactionResponse;
</script>

<UserNavbar {breadcrumbs} />

<div id="dashboard-chart-cards" class="grid gap-4 container mx-auto">
  <TransactionSummary {transactionResponse} />
  <CategorySummary {transactionResponse} />
</div>
