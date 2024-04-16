<script lang="ts">
  import LoginForm from "./LoginForm.svelte";
  import ForgotPasswordForm from "./ForgotPasswordForm.svelte";

  enum LoginFormEnum {
    LOGIN,
    FORGOT_PASSWORD,
  }
  let step = LoginFormEnum.LOGIN;

  function goToLogin() {
    step = LoginFormEnum.LOGIN;
  }

  function onForgotPassword() {
    step = LoginFormEnum.FORGOT_PASSWORD;
  }
</script>

{#if step === LoginFormEnum.LOGIN}
  <LoginForm {onForgotPassword} />
{/if}

{#if step === LoginFormEnum.FORGOT_PASSWORD}
  <ForgotPasswordForm {goToLogin} />
{/if}
