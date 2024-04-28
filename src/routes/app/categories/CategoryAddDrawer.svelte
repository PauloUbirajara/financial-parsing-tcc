<script lang="ts">
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";
  import { Button, CloseButton, Drawer, Input, Label } from "flowbite-svelte";
  import { InfoCircleSolid, PlusOutline } from "flowbite-svelte-icons";
  import { sineIn } from "svelte/easing";
  import { ToastType } from "../../../domain/models/toastMessage";

  export let hideAddDrawer: boolean = true;

  let transitionParams = {
    x: 320,
    duration: 200,
    easing: sineIn,
  };

  let category = {
    name: "",
  };

  async function onAdd() {
    const response = await fetch("/app/categories?/create", {
      method: "POST",
      body: JSON.stringify(category),
    });
    if (response.ok) {
      showToast({
        title: "Adicionar categoria",
        message: `Categoria "${category.name}" adicionada com sucesso.`,
        type: ToastType.SUCCESS,
      });
      goto("/app/categories", { invalidateAll: true });
      return;
    }
    showToast({
      title: "Adicionar categoria",
      message: `Houve um erro ao adicionar a categoria "${category.name}".`,
      type: ToastType.ERROR,
    });
  }
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
  <form class="mb-6" on:submit|preventDefault={onAdd}>
    <div class="mb-6">
      <Label for="name" class="block mb-2">Nome*</Label>
      <Input
        id="name"
        name="name"
        required
        bind:value={category.name}
        placeholder="Digite o nome da categoria"
      />
    </div>
    <Button type="submit" class="w-full">
      Salvar categoria
    </Button>
  </form>
</Drawer>
