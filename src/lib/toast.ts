import type { ToastMessage } from "../domain/models/toastMessage";
import { toastStore } from "../stores/toast";

const MAX_TOASTS_SHOWN = 3;

export function showToast(message: ToastMessage) {
  toastStore.update((toasts: ToastMessage[]) => [
    message,
    ...toasts.slice(0, MAX_TOASTS_SHOWN),
  ]);
}
