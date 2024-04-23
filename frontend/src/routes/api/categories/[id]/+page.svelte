<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumb from "../../../../components/Breadcrumb.svelte";
  import Sidebar from "../../../../components/Sidebar.svelte";
  import CategoryDetail from "./CategoryDetail.svelte";

  const category = $page.data.category;
  const breadcrumbs = [
    { label: "Categorias", href: "/api/categories" },
    { label: category.name, href: `/api/categories/${category.id}` },
  ];
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>
<div class="container mx-auto grid gap-4">
  <CategoryDetail {category} />
</div>
