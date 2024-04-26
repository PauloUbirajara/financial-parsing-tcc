<script lang="ts">
  import { deserialize } from "$app/forms";
  import { goto } from "$app/navigation";
  import { showToast } from "$lib/toast";
  import { Button, Dropdown, DropdownItem } from "flowbite-svelte";
  import { UserSolid } from "flowbite-svelte-icons";
  import { ToastType } from "../domain/models/toastMessage";

  async function logout() {
    const response = await fetch("/auth?/logout", {
      method: "POST",
      body: JSON.stringify({}),
    });
    if (!response.ok) {
      console.log(await response.text());
      return;
    }
    showToast({
      title: "Sair da conta",
      message: "Usuário saiu da conta com sucesso!",
      type: ToastType.SUCCESS,
    });
    goto("/auth/login", { invalidateAll: true });
  }
</script>

<Button class="!p-2 profile-btn" color="dark">
  <UserSolid />
</Button>
<Dropdown triggeredBy=".profile-btn">
  <DropdownItem on:click={logout}>Sair da conta</DropdownItem>
</Dropdown>
