<script lang="ts">
  import Breadcrumb from "./Breadcrumb.svelte";
  import type { BreadcrumbPath } from "../domain/models/breadcrumb";
  import ProfileDropdown from "./ProfileDropdown.svelte";

  import Sidebar from "./Sidebar.svelte";

  export let breadcrumbs: BreadcrumbPath[];
</script>

<div class="flex justify-between">
  <div class="flex items-center gap-4">
    <Sidebar />
    <Breadcrumb {breadcrumbs} />
  </div>
  <ProfileDropdown />
</div>
