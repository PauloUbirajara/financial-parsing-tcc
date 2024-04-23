<script lang="ts">
  import { Button, CloseButton, Drawer, Input, Label } from "flowbite-svelte";
  import { InfoCircleSolid, PlusOutline } from "flowbite-svelte-icons";
  import { sineIn } from "svelte/easing";

  export let hideAddDrawer: boolean = true;

  let transitionParams = {
    x: 320,
    duration: 200,
    easing: sineIn,
  };
</script>

<Drawer
  placement="right"
  transitionType="fly"
  {transitionParams}
  bind:hidden={hideAddDrawer}
  id="sidebar4"
>
  <div class="flex items-center">
    <h5
      id="drawer-label"
      class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
    >
      <InfoCircleSolid class="w-5 h-5 me-2.5" />
      Nova Categoria
    </h5>
    <CloseButton
      on:click={() => (hideAddDrawer = true)}
      class="mb-4 dark:text-white"
    />
  </div>
  <form method="POST" action="?/create" class="mb-6">
    <div class="mb-6">
      <Label for="name" class="block mb-2">Nome*</Label>
      <Input
        id="name"
        name="name"
        required
        placeholder="Digite o nome da categoria"
      />
    </div>
    <Button type="submit" class="w-full">
      <PlusOutline class="w-3.5 h-3.5 me-2.5" />
      Criar categoria
    </Button>
  </form>
</Drawer>
