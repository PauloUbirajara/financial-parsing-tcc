<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { Pagination, Button, type LinkType } from "flowbite-svelte";
  import { ArrowLeftOutline, ArrowRightOutline } from "flowbite-svelte-icons";
  import type { GetAllModelsRepositoryResponse } from "../../domain/models/modelRepositoryDto";
  import { getFilteredUrlSearchParams } from "../../helpers/url";

  export let response: GetAllModelsRepositoryResponse;
  export let baseUrl: string;
  let urlQuery = {
    page: parseInt($page.url.searchParams.get("page") || "1"),
    search: $page.url.searchParams.get("search"),
  };
  if (urlQuery.page < 1 || isNaN(urlQuery.page)) {
    urlQuery.page = 1;
  }
  let filteredUrlQuery = getFilteredUrlSearchParams(urlQuery);
  let currentPage: number = parseInt(filteredUrlQuery.get("page") || "1");
  let hasPreviousLink = Boolean(response.links.previous);
  let hasNextLink = Boolean(response.links.next);

  let pages: LinkType[];
  let previousLink: string;
  let nextLink: string;

  $: {
    pages = Array.from(Array(response.num_pages)).map((_, i): LinkType => {
      let urlQueryClone = new URLSearchParams(
        Object.fromEntries(filteredUrlQuery.entries()),
      );
      urlQueryClone.set("page", (i + 1).toString());
      return {
        name: (i + 1).toString(),
        href: `${baseUrl}?${urlQueryClone.toString()}`,
        active: i + 1 === currentPage,
      };
    });
  }

  previousLink = `${baseUrl}?page=${currentPage - 1}`;
  nextLink = `${baseUrl}?page=${currentPage + 1}`;
  function goToPreviousPage() {
    if (!hasPreviousLink) return;
    if (currentPage === 1) return;
    goto(previousLink);
  }
  function goToNextPage() {
    if (!hasNextLink) return;
    goto(nextLink);
  }
</script>

<div class="flex flex-col items-center justify-center gap-2">
  <div class="text-sm text-gray-700 dark:text-gray-400">
    Página {currentPage}
  </div>

  <Pagination
    {pages}
    table
    large
    on:previous={goToPreviousPage}
    on:next={goToNextPage}
  >
    <svelte:fragment slot="prev">
      <Button color="none" tag="span" disabled={!hasPreviousLink} class="!p-2">
        <ArrowLeftOutline class="me-2 w-5 h-5" />
      </Button>
    </svelte:fragment>
    <svelte:fragment slot="next">
      <Button color="none" tag="span" disabled={!hasNextLink} class="!p-2">
        <ArrowRightOutline class="ms-2 w-5 h-5" />
      </Button>
    </svelte:fragment>
  </Pagination>
</div>
