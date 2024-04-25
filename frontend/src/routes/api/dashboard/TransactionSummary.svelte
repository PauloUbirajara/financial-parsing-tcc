<script lang="ts">
  import { Card, Chart, Spinner } from "flowbite-svelte";
  import type { GetAllModelsRepositoryResponse } from "../../../domain/models/modelRepositoryDto";
  import { navigating } from "$app/stores";

  export let transactionResponse: GetAllModelsRepositoryResponse;

  let wallets = transactionResponse.results.reduce((total, current) => {
    const key = current.wallet.id;
    const name = `${current.wallet.name} (${current.wallet.currency.representation})`;
    total[key] = name;
    return total;
  }, {});

  let transactionsPerCurrency = Object.entries(
    transactionResponse.results.reduce((total, current) => {
      const key = current.wallet.id;
      if (total[key] === undefined) {
        total[key] = 0.0;
      }
      total[key] += Number(current.value);
      return total;
    }, {}),
  );
  transactionsPerCurrency.sort((a, b) => a[1] - b[1]);

  const options = {
    series: [
      {
        name: "Valor total",
        data: transactionsPerCurrency.map((c) => c[1]),
      },
    ],
    chart: {
      sparkline: {
        enabled: false,
      },
      type: "bar",
      width: "100%",
      height: 400,
      toolbar: {
        show: false,
      },
    },
    fill: {
      opacity: 1,
    },
    plotOptions: {
      bar: {
        horizontal: false,
        columnWidth: "100%",
        borderRadiusApplication: "end",
        borderRadius: 6,
        dataLabels: {
          position: "top",
        },
      },
    },
    legend: {
      show: true,
      position: "bottom",
    },
    dataLabels: {
      enabled: false,
    },
    tooltip: {
      shared: true,
      intersect: false,
      formatter: function (value: string) {
        return "$" + value;
      },
    },
    xaxis: {
      labels: {
        show: true,
        style: {
          cssClass: "text-xs font-normal fill-gray-500 dark:fill-gray-400",
        },
        formatter: (value: string) => {
          return value;
        },
      },
      categories: Object.values(wallets),
      axisTicks: {
        show: false,
      },
      axisBorder: {
        show: false,
      },
    },
    yaxis: {
      labels: {
        show: true,
        style: {
          fontFamily: "Inter, sans-serif",
          cssClass: "text-xs font-normal fill-gray-500 dark:fill-gray-400",
        },
      },
    },
    grid: {
      show: true,
      strokeDashArray: 4,
    },
  };
</script>

<Card size="lg">
  <div class="content flex flex-col gap-4">
    <h5
      class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
    >
      Valor total de transações (por carteira)
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
</Card>
