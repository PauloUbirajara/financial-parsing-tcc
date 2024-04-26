<script lang="ts">
  import type { Category } from "../../../domain/models/category";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import { Chart, Heading, Spinner } from "flowbite-svelte";
  import { navigating } from "$app/stores";

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
  const colors = [
    "#4169E1",
    "#FF8C00",
    "#006400",
    "#DC143C",
    "#483D8B",
    "#008B8B",
    "#9932CC",
    "#B8860B",
    "#8FBC8F",
    "#4B0082",
  ];

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
              label: "categorias",
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
      position: "right",
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
  <Heading tag="h4">10 categorias mais usadas</Heading>
  {#if $navigating}
    <div class="mx-auto">
      <Spinner />
    </div>
  {:else if transactionResponse.count}
    <Chart {options} />
  {:else}
    <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
      Sem transações
    </p>
  {/if}
</div>
