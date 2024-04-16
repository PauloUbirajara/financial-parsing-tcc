import AuthManager from "$lib/auth/AuthManager";
import type {
  LoginCredentials,
  RegisterCredentials,
} from "../../domain/models/auth";
import { constants } from "http2";
import { fail, type Actions } from "@sveltejs/kit";

export const actions: Actions = {
  login: async (event) => {
    const credentials: LoginCredentials = await event.request.json();
    const response = await AuthManager.login(credentials);

    if (response.username || response.password || response.detail) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        username: response.username,
        password: response.password,
        detail: response.detail,
      });
    }

    if (!(response.access && response.refresh)) {
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        detail: "Could not login",
      });
    }

    event.cookies.set("accessToken", response.access, { path: "/" });
    event.cookies.set("refreshToken", response.refresh, { path: "/" });
    return { success: true, errors: {} };
  },

  register: async (event) => {
    const credentials: RegisterCredentials = await event.request.json();
    const response = await AuthManager.register(credentials);

    if (
      response.username ||
      response.password ||
      response.confirmPassword ||
      response.error
    ) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        username: response.username,
        password: response.password,
        confirmPassword: response.confirmPassword,
        error: response.error,
      });
    }

    return { success: true, errors: {} };
  },
};
