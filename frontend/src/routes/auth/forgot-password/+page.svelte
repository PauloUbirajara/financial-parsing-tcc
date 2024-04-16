<script lang="ts">
  import { Heading, Input, Label, Button, Helper } from "flowbite-svelte";
</script>

<form class="p-4">
  <div
    class="flex align-center flex-col sm:flex-row"
    id="register-form-heading"
  >
    <Heading tag="h3" class="text-custom-foreground"
      >Esqueceu sua senha?</Heading
    >
  </div>
  <div class="flex flex-col gap-5">
    <div>
      <Label for="username">Apelido</Label>
      <Input
        type="text"
        id="username"
        placeholder="Digite seu apelido (ex.: JohnDoe)"
        class="mt-2"
        required
      />
      <Helper class="text-sm mt-2">
        Caso o apelido pertença a um usuário cadastrado, enviaremos uma nova
        senha via e-mail para que acesse a nossa plataforma.
      </Helper>
    </div>
  </div>
  <Button type="submit">Enviar</Button>
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
