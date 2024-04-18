<script lang="ts">
  import {
    Alert,
    Button,
    Heading,
    Helper,
    Input,
    Label,
    Spinner,
  } from "flowbite-svelte";
  import { ArrowLeftOutline, InfoCircleSolid } from "flowbite-svelte-icons";
  import { fly } from "svelte/transition";

  let credentials = {
    email: "",
  };

  type ResetPasswordForm = {
    success: string;
    errors: {
      [field: string]: any;
    };
  };
  let formMessages: ResetPasswordForm = {
    success: "",
    errors: {},
  };
  $: isValid = credentials.email;
  let loading: boolean = false;

  async function handleSubmit() {
    formMessages = {
      success: "",
      errors: {},
    };

    try {
      loading = true;
      const response = await fetch("/auth?/reset-password", {
        method: "POST",
        body: JSON.stringify(credentials),
      });

      if (response.ok) {
        formMessages.success =
          "Caso o e-mail esteja cadastrado, um link será enviado com as instruções para redefinição da senha.";
        return;
      }

      const data = await response.json();
      formMessages.errors = data.errors;
    } catch (e) {
      console.warn(e);
      formMessages.errors.detail = "Erro ao solicitar redefinição de senha.";
    } finally {
      loading = false;
    }
  }
</script>

{#if formMessages.success}
  <Alert color="green" dismissable>
    <InfoCircleSolid slot="icon" class="w-4 h-4" />
    {formMessages.success}
  </Alert>
{/if}

<form
  in:fly={{ x: "-20" }}
  method="POST"
  class="grid gap-4 p-4"
  on:submit|preventDefault={handleSubmit}
>
  <div class="flex gap-4 items-center">
    <Button
      outline={true}
      class="!p-2"
      size="lg"
      color="primary"
      href="/auth/login"
    >
      <ArrowLeftOutline class="w-6" />
    </Button>
    <Heading tag="h3" class="text-custom-foreground"
      >Redefinição de senha</Heading
    >
  </div>

  <div>
    <Label for="email">E-mail</Label>
    <Input
      type="text"
      id="email"
      name="email"
      placeholder="Digite seu e-mail (ex.: johndoe@email.com)"
      class="mt-2"
      bind:value={credentials.email}
      required
    />
  </div>

  <Button type="submit" disabled={loading || !isValid}>
    {#if loading}
      <Spinner class="me-3" size="4" color="white" />
    {:else}
      Enviar
    {/if}
  </Button>
</form>

<style>
  form {
    height: 100%;
    grid-template-rows: auto 1fr auto;
    overflow-y: auto;
  }
</style>
