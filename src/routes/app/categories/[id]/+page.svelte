<script lang="ts">
  import { page } from "$app/stores";
  import UserNavbar from "../../../../components/UserNavbar.svelte";
  import CategoryDetail from "./CategoryDetail.svelte";

  const category = $page.data.category;
  const breadcrumbs = [
    { label: "Categorias", href: "/app/categories" },
    { label: category.name, href: `/app/categories/${category.id}` },
  ];
</script>

<UserNavbar {breadcrumbs} />
<div class="container mx-auto grid gap-8 h-min">
  <CategoryDetail {category} />
</div>
