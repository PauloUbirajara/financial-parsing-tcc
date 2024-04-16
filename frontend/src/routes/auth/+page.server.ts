import AuthManager from "$lib/auth/AuthManager";
import type {
  ForgotPasswordCredentials,
  LoginCredentials,
  RegisterCredentials,
} from "../../domain/models/auth";
import { constants } from "http2";
import { fail, type Actions } from "@sveltejs/kit";

export const actions: Actions = {
  login: async (event) => {
    const credentials: LoginCredentials = await event.request.json();
    const response = await AuthManager.login(credentials);

    if (response.username || response.password) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        success: false,
        errors: {
          username: response.username,
          password: response.password,
        },
      });
    }

    if (response.detail) {
      console.warn(response.detail);
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        success: false,
        errors: {
          detail: "Erro ao realizar login.",
        },
      });
    }

    if (!(response.access && response.refresh)) {
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        success: false,
        errors: {
          detail: "Erro ao realizar login.",
        },
      });
    }

    event.cookies.set("accessToken", response.access, { path: "/" });
    event.cookies.set("refreshToken", response.refresh, { path: "/" });
    return { success: true, errors: {} };
  },

  register: async (event) => {
    const credentials: RegisterCredentials = await event.request.json();
    const response = await AuthManager.register(credentials);

    if (response.non_field_errors) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        success: false,
        errors: {
          detail: response.non_field_errors,
        },
      });
    }

    if (response.username || response.password || response.confirmPassword) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        errors: {
          username: response.username,
          password: response.password,
          confirmPassword: response.confirmPassword,
        },
        success: false,
      });
    }

    if (response.error || response.non_field_errors) {
      console.warn(response.error, response.non_field_errors);
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        errors: {
          detail: "Erro ao realizar cadastro.",
        },
        success: false,
      });
    }

    return {
      success: true,
      errors: {},
    };
  },

  forgotPassword: async (event) => {
    const credentials: ForgotPasswordCredentials = await event.request.json();
    const response = await AuthManager.resetPassword(credentials);
  },
};
