<script lang="ts">
  import { Chart, Spinner } from "flowbite-svelte";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import { navigating } from "$app/stores";
  import type { Category } from "../../../domain/models/category";

  export let transactionResponse: GetAllModelsRepositoryResponse;

  let categoryNames = transactionResponse.results.reduce((total, current) => {
    const categories = current.categories || [];
    categories.forEach((c: Category) => {
      total[c.id!] = c.name;
    });
    return total;
  }, {});

  let transactionsPerCategory = Object.entries(
    transactionResponse.results.reduce((total, current) => {
      const categories = current.categories || [];
      categories.forEach((c: Category) => {
        total[c.id!] = (total[c.id!] || 0) + 1;
      });
      return total;
    }, {}),
  );

  // Let only top 10
  transactionsPerCategory.sort(
    (a: [string, number], b: [string, number]) => b[1] - a[1],
  );
  transactionsPerCategory = transactionsPerCategory.slice(0, 10);

  const labels = transactionsPerCategory.map((c) => categoryNames[c[0]]);
  const series = transactionsPerCategory.map((c) => c[1]);
  const colors = ["#25AF9B", "#D16C2C", "#48211C", "#A8152B", "#1A5BF1"];

  const options = {
    labels,
    series,
    colors,
    chart: {
      height: 320,
      width: "100%",
      type: "donut",
      toolbar: {
        show: true,
      },
    },
    stroke: {
      colors: ["transparent"],
      lineCap: "",
    },
    plotOptions: {
      pie: {
        donut: {
          labels: {
            show: true,
            name: {
              show: true,
              offsetY: 20,
            },
            total: {
              showAlways: true,
              show: true,
            },
            value: {
              show: true,
              offsetY: -20,
            },
          },
          size: "70%",
        },
      },
    },
    dataLabels: {
      enabled: false,
    },
    legend: {
      position: "bottom",
    },
    xaxis: {
      axisTicks: {
        show: false,
      },
      axisBorder: {
        show: false,
      },
    },
  };
</script>

<div class="content flex flex-col gap-4 p-4 bg-white rounded">
  <h5
    class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
  >
    10 categorias mais usadas
  </h5>
  {#if $navigating}
    <div class="mx-auto">
      <Spinner />
    </div>
  {:else if transactionResponse}
    <Chart {options} />
  {:else}
    <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
      Sem transações
    </p>
  {/if}
</div>
