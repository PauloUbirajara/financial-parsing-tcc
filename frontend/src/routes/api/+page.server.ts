import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
  logout: async (event) => {
    event.cookies.delete("accessToken", { path: "/" });
    event.cookies.delete("refreshToken", { path: "/" });
  },
};
