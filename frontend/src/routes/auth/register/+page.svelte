<script lang="ts">
  import { Input, Label, Button, P, Card, Alert } from "flowbite-svelte";
  import {
    EyeOutline,
    EyeSlashOutline,
    InfoCircleSolid,
  } from "flowbite-svelte-icons";

  let showPassword = false;
  let submitted = false;

  function onRegister() {
    submitted = true;
  }

  export let form;
</script>

<div class="flex flex-col items-start gap-5">
  {#if [true, false].includes(form?.success)}
    <Alert
      color={form?.success === true ? "green" : "red"}
      class="!items-start"
    >
      <span slot="icon">
        <InfoCircleSolid slot="icon" class="w-4 h-4" />
        <span class="sr-only">Info</span>
      </span>
      <P>{form?.title}: {form?.message}</P>
    </Alert>
  {/if}

  <Card size="md">
    <div class="flex gap-5 items-center justify-between mb-5">
      <P size="4xl" weight="bold">Crie sua conta</P>
      <Button color="alternative" href="/auth/login">Já possuo conta</Button>
    </div>

    <form method="POST" on:submit={onRegister}>
      <div class="mb-6">
        <Label for="email" class="mb-2">E-mail</Label>
        <Input
          type="email"
          id="email"
          name="email"
          size="lg"
          placeholder="john.doe@company.com"
          required
        />
      </div>
      <div class="mb-6">
        <Label for="password" class="mb-2">Senha</Label>
        <Input
          id="password"
          name="password"
          type={showPassword ? "text" : "password"}
          placeholder="•••••••••"
          size="lg"
          required
        >
          <Button
            slot="right"
            on:click={() => (showPassword = !showPassword)}
            class="pointer-events-auto"
            color="none"
            size="sm"
          >
            {#if showPassword}
              <EyeOutline class="w-6 h-6" />
            {:else}
              <EyeSlashOutline class="w-6 h-6" />
            {/if}
          </Button>
        </Input>
      </div>
      <div class="mb-6">
        <Label for="confirm-password" class="mb-2">Senha</Label>
        <Input
          id="confirm-password"
          name="confirm-password"
          type={showPassword ? "text" : "password"}
          placeholder="•••••••••"
          size="lg"
          required
        >
          <Button
            slot="right"
            on:click={() => (showPassword = !showPassword)}
            class="pointer-events-auto"
            color="none"
            size="sm"
          >
            {#if showPassword}
              <EyeOutline class="w-6 h-6" />
            {:else}
              <EyeSlashOutline class="w-6 h-6" />
            {/if}
          </Button>
        </Input>
      </div>
      <Button type="submit" disabled={submitted}>Acessar</Button>
    </form>
  </Card>
</div>
