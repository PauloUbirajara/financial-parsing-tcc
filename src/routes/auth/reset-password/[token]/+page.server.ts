import AuthManager from "$lib/auth/AuthManager";
import { constants } from "http2";
import type { ResetPasswordCredentials } from "../../../../domain/models/resetPasswordDto";
import { fail } from "@sveltejs/kit";
import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
  default: async (event) => {
    const credentials: ResetPasswordCredentials = await event.request.json();
    const response = await AuthManager.resetPassword(credentials);

    if (response.error) {
      return fail(constants.HTTP_STATUS_NOT_FOUND, {
        success: false,
        errors: {
          error: response.error,
        },
      });
    }

    if (response.password || response.non_field_errors) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        success: false,
        errors: {
          password: response.password,
          non_field_errors: response.non_field_errors,
        },
      });
    }

    return { success: true, errors: {} };
  },
};
