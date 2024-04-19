<script lang="ts">
  import { navigating, page } from "$app/stores";
  import { Button, Heading, P, TextPlaceholder } from "flowbite-svelte";
  import { ArrowLeftOutline } from "flowbite-svelte-icons";

  export let onDetail: Function;

  let wallet = $page.data.wallet;
</script>

{#if $navigating}
  <TextPlaceholder divClass="space-y-2.5 animate-pulse mx-auto w-full" />
{:else}
  <div class="flex gap-4 items-center">
    <Button
      outline={true}
      class="!p-2"
      size="lg"
      color="primary"
      on:click={() => onDetail()}
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <Heading tag="h3" class="text-custom-foreground">
      Editar {wallet.name}
    </Heading>
  </div>
  <P>EDITAR!</P>
{/if}
