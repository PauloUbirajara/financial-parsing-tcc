import { writable } from "svelte/store";
import type { Writable } from "svelte/store";

export const toastStore: Writable<any[]> = writable([]);
