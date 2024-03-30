<script lang="ts">
  import { Input, Label, Button, P } from "flowbite-svelte";
  import { EyeOutline, EyeSlashOutline } from "flowbite-svelte-icons";
  let showPassword = false;
</script>

<div class="flex gap-5 items-center justify-between mb-5">
  <P size="4xl" weight="bold">Acesse sua conta</P>
  <Button color="alternative" href="/auth/register">Não possuo conta</Button>
</div>

<form method="POST">
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
  <Button type="submit">Acessar</Button>
</form>
