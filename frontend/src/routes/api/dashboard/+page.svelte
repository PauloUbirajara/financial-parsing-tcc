<script lang="ts">
  import Sidebar from "../../../components/Sidebar.svelte";
  import Breadcrumb from "../../../components/Breadcrumb.svelte";

  const breadcrumbs = [{ label: "Dashboard", href: "/api/dashboard" }];
</script>

<div class="flex items-center gap-4">
  <Sidebar />
  <Breadcrumb {breadcrumbs} />
</div>
