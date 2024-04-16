<script lang="ts">
  import ConfirmAccountForm from "./ConfirmAccountForm.svelte";
  import RegisterForm from "./RegisterForm.svelte";

  enum RegisterFormEnum {
    REGISTER,
    CONFIRM_ACCOUNT,
  }
  let step = RegisterFormEnum.REGISTER;

  function onRegisterSuccess() {
    step = RegisterFormEnum.CONFIRM_ACCOUNT;
  }
</script>

{#if step === RegisterFormEnum.REGISTER}
  <RegisterForm onSubmitSuccess={onRegisterSuccess} />
{/if}

{#if step === RegisterFormEnum.CONFIRM_ACCOUNT}
  <ConfirmAccountForm />
{/if}
