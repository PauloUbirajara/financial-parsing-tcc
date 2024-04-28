import type { ToastMessage } from "../domain/models/toastMessage";
import { toastStore } from "../stores/toast";

export function showToast(message: ToastMessage) {
  toastStore.set(message);
}
