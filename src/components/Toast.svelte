<script lang="ts">
  import { onMount } from "svelte";
  import type { ToastMessage } from "../domain/models/toastMessage";
  import { ToastType } from "../domain/models/toastMessage";
  import { toastStore } from "../stores/toast";
  import { Progressbar, Toast } from "flowbite-svelte";
  import { fly } from "svelte/transition";
  import {
    CheckCircleSolid,
    CloseCircleSolid,
    ExclamationCircleSolid,
  } from "flowbite-svelte-icons";

  let toast: ToastMessage | null = null;
  let showToast: boolean = true;
  let currentDuration = 0;
  let interval: NodeJS.Timeout;
  const MAX_TOAST_DURATION_IN_MS = 30000;

  function removeToast() {
    showToast = false;
    currentDuration = 0;
    toastStore.set(null);
  }

  setInterval(() => {
    if (toast === null) {
      clearInterval(interval);
      return;
    }
    currentDuration += 1000;
    if (currentDuration < MAX_TOAST_DURATION_IN_MS) {
      return;
    }
    currentDuration = 0;
    removeToast();
  }, 1000);

  onMount(() => {
    const unsubscribe = toastStore.subscribe((value) => {
      toast = value;
      showToast = true;
      currentDuration = 0;
    });

    return unsubscribe;
  });

  type ToastColors =
    | "green"
    | "red"
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

{#if toast !== null}
  <Toast
    transition={fly}
    params={{ x: 200 }}
    color={getColorFromToast(toast)}
    class="mb-4"
    bind:open={showToast}
    on:close={() => removeToast()}
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
    <span class="mb-1 text-lg font-semibold text-gray-900 dark:text-white">
      {toast.title}
    </span>
    <div class="mb-2 text-md font-normal">{toast.message}</div>
    <Progressbar
      color={getColorFromToast(toast)}
      animate={true}
      tweenDuration={1000}
      progress={(currentDuration / MAX_TOAST_DURATION_IN_MS) * 100.0}
      size="h-1"
      divClass="w-full bg-gray-200 rounded-full dark:bg-gray-700 relative bottom-0"
    />
  </Toast>
{/if}
