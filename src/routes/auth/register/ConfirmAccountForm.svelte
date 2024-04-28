<script lang="ts">
  import { Button, Heading, P } from "flowbite-svelte";
  import { fly } from "svelte/transition";
</script>

<div
  in:fly={{ x: "-20" }}
  class="p-4 grid gap-4"
  id="confirm-account-container"
>
  <Heading tag="h3" class="text-custom-foreground">Confirme sua conta</Heading>
  <div class="flex flex-col gap-4">
    <P>Enviamos um link de confirmação para o e-mail fornecido.</P>
    <P>
      Após confirmá-lo, sua conta será ativada e você poderá acessar à
      plataforma.</P
    >
  </div>
  <Button href="/auth/login">Ir à tela de login</Button>
</div>

<style>
  #confirm-account-container {
    height: 100%;
    grid-template-rows: auto 1fr auto;
  }
</style>
