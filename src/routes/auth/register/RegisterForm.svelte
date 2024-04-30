<script lang="ts">
  import {
    Spinner,
    Heading,
    A,
    Input,
    Label,
    Button,
    Helper,
    Alert,
  } from "flowbite-svelte";
  import {
    EyeOutline,
    EyeSlashOutline,
    InfoCircleSolid,
  } from "flowbite-svelte-icons";
  import { deserialize } from "$app/forms";
  import { fly } from "svelte/transition";

  type FormMessages = {
    errors: {
      error?: string;
      detail?: string;
      email?: string[];
      username?: string[];
      password?: string[];
    };
    success: string;
  };
  let formMessages: FormMessages = { success: "", errors: {} };
  let credentials = {
    username: "",
    password: "",
    email: "",
    confirmPassword: "",
  };

  let loading: boolean = false;
  let showPassword: boolean = false;
  export let onSubmitSuccess: Function;

  $: hasConfirmedPassword =
    credentials.password === credentials.confirmPassword;
  $: isValid =
    credentials.username && credentials.email && hasConfirmedPassword;

  async function handleSubmit() {
    loading = true;
    formMessages = {
      success: "",
      errors: {},
    };

    try {
      const response = await fetch("/auth?/register", {
        body: JSON.stringify(credentials),
        method: "POST",
      });

      let registerResponse = deserialize(await response.text());
      formMessages.errors = registerResponse.data.errors;

      if (response.ok && registerResponse.data.success) {
        formMessages.success = "Cadastro realizado com sucesso.";
      }
    } catch (e: any) {
      console.warn("Failed to register", e);
      formMessages.errors.detail = "Houve um erro ao tentar realizar cadastro.";
    } finally {
      loading = false;
    }

    if (formMessages.success) {
      formMessages.success = "Cadastro realizado com sucesso.";
      onSubmitSuccess();
    }
  }
</script>

{#if formMessages.errors.detail}
  <Alert dismissable>
    <InfoCircleSolid slot="icon" class="w-4 h-4" />
    {formMessages.errors.detail}
  </Alert>
{/if}

{#if formMessages.success}
  <Alert color="green" dismissable>
    <InfoCircleSolid slot="icon" class="w-4 h-4" />
    {formMessages.success}
  </Alert>
{/if}

<form class="p-4" on:submit|preventDefault={handleSubmit} in:fly={{ x: "-20" }}>
  <div
    class="flex align-center flex-col sm:flex-row"
    id="register-form-heading"
  >
    <Heading tag="h3" class="text-custom-foreground">Crie Sua Conta</Heading>
    <A href="/auth/login" class="text-custom-primary font-bold"
      >Já sou usuário</A
    >
  </div>
  <div class="flex flex-col gap-5">
    <div>
      <Label for="username">Apelido</Label>
      <Input
        type="text"
        id="username"
        name="username"
        placeholder="Digite seu apelido (ex.: JohnDoe)"
        class="mt-2"
        color={formMessages.errors.username ? "red" : "base"}
        bind:value={credentials.username}
        required
      />
      {#if formMessages.errors.username}
        <Helper class="mt-2" color="red"
          >{formMessages.errors.username.join(", ")}</Helper
        >
      {/if}
    </div>
    <div>
      <Label for="email">E-mail</Label>
      <Input
        type="text"
        id="email"
        name="email"
        placeholder="Digite seu e-mail (ex.: johndoe@email.com)"
        class="mt-2"
        color={formMessages.errors.email ? "red" : "base"}
        bind:value={credentials.email}
        required
      />
      {#if formMessages.errors.email}
        <Helper class="mt-2" color="red"
          >{formMessages.errors.email.join(", ")}</Helper
        >
      {/if}
    </div>
    <div>
      <Label for="password">Senha</Label>
      <Input
        type={showPassword ? "text" : "password"}
        id="password"
        name="password"
        placeholder="Digite sua senha"
        class="mt-2"
        color={formMessages.errors.password ? "red" : "base"}
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
        <Helper class="mt-2" color="red"
          >{formMessages.errors.password.join(", ")}</Helper
        >
      {/if}
    </div>
    <div>
      <Label for="confirmPassword">Confirmar Senha</Label>
      <Input
        type={showPassword ? "text" : "password"}
        id="confirmPassword"
        name="confirmPassword"
        placeholder="Digite sua senha novamente"
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
      {#if credentials.password.length && !hasConfirmedPassword}
        <Helper class="mt-2" color="red">Senhas não coincidem</Helper>
      {/if}
    </div>
  </div>
  <Button type="submit" disabled={loading || !isValid}>
    {#if loading}
      <Spinner class="me-3" size="4" color="white" />
    {:else}
      Criar
    {/if}
  </Button>
</form>

<style>
  form {
    display: grid;
    grid-template-rows: auto 1fr auto;
    height: 100%;
    gap: 50px;
    overflow-y: auto;
  }

  #register-form-heading {
    text-wrap: nowrap;
  }
</style>
