<script lang="ts">
  import { page } from "$app/stores";
  import { Alert, Heading, P } from "flowbite-svelte";
</script>

<Alert class="!items-start">
  <Heading size="2xl">Houve algum erro...</Heading>
  <P size="2xl">{$page.status}: {$page.error?.message}</P>
</Alert>
