<script lang="ts">
  import { onMount } from "svelte";
  import type { ToastMessage } from "../domain/models/toastMessage";
  import { ToastType } from "../domain/models/toastMessage";
  import { toastStore } from "../stores/toast";
  import { Toast } from "flowbite-svelte";
  import { fly } from "svelte/transition";
  import {
    CheckCircleSolid,
    CloseCircleSolid,
    ExclamationCircleSolid,
  } from "flowbite-svelte-icons";

  let toasts: ToastMessage[] = [];
  const MAX_TOASTS_SHOWN = 2;

  onMount(() => {
    const unsubscribe = toastStore.subscribe((value) => {
      toasts = value;
    });

    return unsubscribe;
  });

  type ToastColors =
    | "green"
    | "red"
    | "orange"
    | "none"
    | "gray"
    | "yellow"
    | "indigo"
    | "purple"
    | "blue"
    | "primary"
    | undefined;

  function getColorFromToast(toast: ToastMessage): ToastColors {
    const supportedColors = {
      [ToastType.SUCCESS]: "green",
      [ToastType.ERROR]: "red",
      [ToastType.WARNING]: "orange",
    };
    return (supportedColors[toast.type] || "none") as ToastColors;
  }
</script>

{#each toasts.slice(0, MAX_TOASTS_SHOWN) as toast}
  <Toast
    transition={fly}
    params={{ x: 200 }}
    color={getColorFromToast(toast)}
    class="mb-4"
  >
    <svelte:fragment slot="icon">
      {#if toast.type === ToastType.SUCCESS}
        <CheckCircleSolid class="w-5 h-5" />
        <span class="sr-only">Check icon</span>
      {/if}
      {#if toast.type === ToastType.ERROR}
        <CloseCircleSolid class="w-5 h-5" />
        <span class="sr-only">Error icon</span>
      {/if}
      {#if toast.type === ToastType.WARNING}
        <ExclamationCircleSolid class="w-5 h-5" />
        <span class="sr-only">Warning icon</span>
      {/if}
    </svelte:fragment>
    <span class="mb-1 text-lg font-semibold text-gray-900 dark:text-white"
      >{toast.title}</span
    >
    <div class="mb-2 text-md font-normal">{toast.message}</div>
  </Toast>
{/each}
