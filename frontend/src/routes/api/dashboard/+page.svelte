<script lang="ts">
  import Sidebar from "../../../components/Sidebar.svelte";
  import Breadcrumb from "../../../components/Breadcrumb.svelte";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import TransactionSummary from "./TransactionSummary.svelte";
  import { page } from "$app/stores";
  import CategorySummary from "./CategorySummary.svelte";

  const breadcrumbs = [{ label: "Dashboard", href: "/api/dashboard" }];

  const transactionResponse: GetAllModelsRepositoryResponse =
    $page.data.transactionResponse;
  const currencyResponse: GetAllModelsRepositoryResponse =
    $page.data.currencyResponse;
  const categoryResponse: GetAllModelsRepositoryResponse =
    $page.data.categoryResponse;
  const walletResponse: GetAllModelsRepositoryResponse =
    $page.data.walletResponse;
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>

<div
  id="dashboard-chart-cards"
  class="flex justify-center items-start gap-4 container mx-auto"
>
  <TransactionSummary {transactionResponse} />
  <CategorySummary {transactionResponse} />
</div>

<style>
  #dashboard-chart-cards {
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  }
</style>
