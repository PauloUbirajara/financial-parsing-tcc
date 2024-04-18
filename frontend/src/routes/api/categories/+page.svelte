<script lang="ts">
  import { P } from "flowbite-svelte";
  import { page } from "$app/stores";
</script>

<P>Categories</P>

{JSON.stringify($page.data.categories)}
