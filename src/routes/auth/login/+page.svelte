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
  import { goto } from "$app/navigation";
  import { fly } from "svelte/transition";

  type FormMessages = {
    success: string;
    errors: {
      detail?: string;
      username?: string[];
      password?: string[];
    };
  };
  let formMessages: FormMessages = { success: "", errors: {} };
  let credentials = {
    username: "",
    password: "",
  };

  $: isValid = credentials.username && credentials.password;
  let loading: boolean = false;
  let showPassword: boolean = false;

  async function handleSubmit() {
    loading = true;
    formMessages = {
      success: "",
      errors: {},
    };

    try {
      const response = await fetch("/auth?/login", {
        body: JSON.stringify(credentials),
        method: "POST",
      });

      let loginResponse = deserialize(await response.text());
      formMessages.errors = loginResponse.data.errors;

      if (loginResponse.data.success) {
        formMessages.success = "Login realizado com sucesso.";
      }
    } catch (e: any) {
      console.warn(e);
      formMessages.errors.detail = "Houve um erro ao tentar realizar login.";
    } finally {
      loading = false;
    }

    if (formMessages.success) {
      setTimeout(() => {
        goto("/app/dashboard");
      }, 500);
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
  <div class="flex align-center flex-col sm:flex-row" id="login-form-heading">
    <Heading tag="h3" class="text-custom-foreground">Acesse Sua Conta</Heading>
    <A href="/auth/register" class="text-custom-primary font-bold"
      >Não possuo conta</A
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
        color={formMessages.username ? "red" : "base"}
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
    <A href="/auth/reset-password" class="text-custom-primary font-bold text-sm"
      >Esqueceu sua senha?</A
    >
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
    display: grid;
    grid-template-rows: auto 1fr auto;
    height: 100%;
    gap: 50px;
    overflow-y: auto;
  }

  #login-form-heading {
    text-wrap: nowrap;
  }
</style>
