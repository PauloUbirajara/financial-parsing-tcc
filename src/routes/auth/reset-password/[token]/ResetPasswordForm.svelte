<script lang="ts">
  import {
    Input,
    Button,
    Heading,
    Label,
    Spinner,
    Helper,
    Alert,
  } from "flowbite-svelte";
  import { fly } from "svelte/transition";
  import {
    EyeSlashOutline,
    EyeOutline,
    InfoCircleSolid,
  } from "flowbite-svelte-icons";
  import { deserialize } from "$app/forms";

  export let token: string;
  export let onResetPasswordSuccess: Function;

  type FormMessages = {
    errors: {
      error?: string;
      non_field_errors?: string;
      password?: string;
    };
    success: string;
  };
  let formMessages: FormMessages = { success: "", errors: {} };

  let credentials = {
    password: "",
    confirmPassword: "",
  };

  let showPassword: boolean = false;
  let loading: boolean = false;

  $: hasConfirmedPassword =
    credentials.password === credentials.confirmPassword;
  $: isValid = credentials.password && hasConfirmedPassword;

  async function handleSubmit() {
    formMessages = {
      success: "",
      errors: {},
    };

    try {
      loading = true;
      const response = await fetch("?", {
        method: "POST",
        body: JSON.stringify({ password: credentials.password, token }),
      });

      const resetPasswordResponse = deserialize(await response.text());

      if (resetPasswordResponse.type === "success") {
        formMessages.success = "Redefinição de senha realizada com sucesso.";
        onResetPasswordSuccess();
      }
      formMessages.errors = resetPasswordResponse.data.errors;
    } catch (e) {
      console.warn("Error when resetting password", e);
      formMessages.errors.detail = "Erro ao redefinir senha.";
    } finally {
      loading = false;
    }
  }
</script>

{#if formMessages.errors.error}
  <Alert dismissable>
    <InfoCircleSolid slot="icon" class="w-4 h-4" />
    {formMessages.errors.error}
  </Alert>
{/if}

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
  <div
    class="flex align-center flex-col sm:flex-row"
    id="reset-password-form-heading"
  >
    <Heading tag="h3" class="text-custom-foreground"
      >Redefinição de senha</Heading
    >
  </div>
  <div class="flex flex-col gap-5">
    <div>
      <Label for="password">Nova senha</Label>
      <Input
        type={showPassword ? "text" : "password"}
        id="password"
        name="password"
        placeholder="Digite sua nova senha"
        class="mt-2"
        color={formMessages.errors.password ||
        formMessages.errors.non_field_errors
          ? "red"
          : "base"}
        bind:value={credentials.password}
        required
      >
        <button
          type="button"
          slot="right"
          on:click={() => (showPassword = !showPassword)}
        >
          {#if showPassword}
            <EyeSlashOutline />
          {:else}
            <EyeOutline />
          {/if}
        </button>
      </Input>
      {#if formMessages.errors.password}
        <Helper class="mt-2" color="red">{formMessages.errors.password}</Helper>
      {/if}
      {#if formMessages.errors.non_field_errors}
        <Helper class="mt-2" color="red"
          >{formMessages.errors.non_field_errors.join(", ")}</Helper
        >
      {/if}
      {#if credentials.password.length && !hasConfirmedPassword}
        <Helper class="mt-2" color="red">Senhas não coincidem</Helper>
      {/if}
    </div>
    <div>
      <Label for="confirmPassword">Confirme a nova senha</Label>
      <Input
        type={showPassword ? "text" : "password"}
        id="confirmPassword"
        name="confirmPassword"
        placeholder="Digite sua nova senha novamente"
        class="mt-2"
        bind:value={credentials.confirmPassword}
        required
      >
        <button
          type="button"
          slot="right"
          on:click={() => (showPassword = !showPassword)}
        >
          {#if showPassword}
            <EyeSlashOutline />
          {:else}
            <EyeOutline />
          {/if}
        </button>
      </Input>
      {#if credentials.confirmPassword.length && !hasConfirmedPassword}
        <Helper class="mt-2" color="red">Senhas não coincidem</Helper>
      {/if}
    </div>
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
    gap: 50px;
    overflow-y: auto;
  }
</style>
