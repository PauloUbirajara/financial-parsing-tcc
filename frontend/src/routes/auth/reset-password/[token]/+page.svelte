<script lang="ts">
  import { page } from "$app/stores";
  import ResetPasswordForm from "./ResetPasswordForm.svelte";
  import ResetPasswordSuccessForm from "./ResetPasswordSuccessForm.svelte";

  enum RegisterFormEnum {
    RESET_PASSWORD,
    RESET_PASSWORD_SUCCESS,
  }
  const token = $page.params.token;
  let step = RegisterFormEnum.RESET_PASSWORD;

  function onResetPasswordSuccess() {
    step = RegisterFormEnum.RESET_PASSWORD_SUCCESS;
  }
</script>

{#if step === RegisterFormEnum.RESET_PASSWORD}
  <ResetPasswordForm {onResetPasswordSuccess} {token} />
{/if}

{#if step === RegisterFormEnum.RESET_PASSWORD_SUCCESS}
  <ResetPasswordSuccessForm />
{/if}
