<script lang="ts">
  import { Modal, Button } from "flowbite-svelte";
  import { ExclamationCircleOutline } from "flowbite-svelte-icons";

  export let showDeleteModal: boolean;
  export let title: string;
  export let onDelete: Function;
</script>

<Modal open={showDeleteModal} size="xs" autoclose>
  <div class="text-center">
    <ExclamationCircleOutline
      class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
    />
    <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
      {title}
    </h3>
    <Button color="red" class="me-2" on:click={() => onDelete()}
      >Sim, apagar</Button
    >
    <Button color="alternative" on:click={() => (showDeleteModal = false)}
      >Não, cancelar</Button
    >
  </div>
</Modal>
