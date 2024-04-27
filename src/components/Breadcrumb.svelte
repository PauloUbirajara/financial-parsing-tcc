<script lang="ts">
  import { Breadcrumb, BreadcrumbItem } from "flowbite-svelte";
  import type { BreadcrumbPath } from "../domain/models/breadcrumb";

  export let breadcrumbs: BreadcrumbPath[] = [];
</script>

<Breadcrumb>
  {#each breadcrumbs as breadcrumb, i}
    <BreadcrumbItem href={breadcrumb.href} home={i === 0}>
      {breadcrumb.label}
    </BreadcrumbItem>
  {/each}
</Breadcrumb>
