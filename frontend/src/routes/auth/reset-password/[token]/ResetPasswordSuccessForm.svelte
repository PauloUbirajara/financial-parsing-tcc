<script lang="ts">
  import { Button, Heading, P } from "flowbite-svelte";
  import { fly } from "svelte/transition";
</script>

<div
  in:fly={{ x: "-20" }}
  class="p-4 grid gap-4"
  id="reset-password-success-container"
>
  <Heading tag="h3" class="text-custom-foreground">Redefinição de senha</Heading
  >
  <div class="flex flex-col gap-4">
    <P>Sua senha foi redefinida com sucesso!</P>
  </div>
  <Button href="/auth/login">Ir à tela de login</Button>
</div>

<style>
  #reset-password-success-container {
    height: 100%;
    gap: 50px;
    grid-template-rows: auto 1fr auto;
  }
</style>
