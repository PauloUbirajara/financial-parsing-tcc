<script lang="ts">
  import {
    Button,
    Heading,
    Helper,
    Input,
    Label,
    Spinner,
  } from "flowbite-svelte";
  import { ArrowLeftOutline } from "flowbite-svelte-icons";
  import { fly } from "svelte/transition";

  export let goToLogin: Function;

  let credentials = {
    username: "",
  };
  $: isValid = credentials.username;

  let loading: boolean = false;
</script>

<form
  in:fly={{ x: "-20" }}
  action="/auth/forgotPassword"
  class="grid gap-4 p-4"
>
  <div class="flex gap-4 items-center">
    <Button
      outline={true}
      class="!p-2"
      size="lg"
      color="primary"
      on:click={() => goToLogin()}
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <Heading tag="h3" class="text-custom-foreground"
      >Redefinição de senha</Heading
    >
  </div>

  <div>
    <Label for="username">Apelido</Label>
    <Input
      type="text"
      id="username"
      name="username"
      placeholder="Digite seu apelido (ex.: JohnDoe)"
      class="mt-2"
      bind:value={credentials.username}
      required
    />
    <Helper class="text-sm mt-2">
      Caso o apelido pertença a um usuário cadastrado, enviaremos uma nova senha
      via e-mail para que acesse a nossa plataforma.
    </Helper>
  </div>

  <Button type="submit" disabled={loading || !isValid}>
    {#if loading}
      <Spinner class="me-3" size="4" color="white" />
    {:else}
      Acessar
    {/if}
  </Button>
</form>

<style>
  form {
    height: 100%;
    grid-template-rows: auto 1fr auto;
  }
</style>
