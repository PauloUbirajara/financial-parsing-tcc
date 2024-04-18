import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { constants } from "http2";

export const load: LayoutServerLoad = async (event) => {
  const access = event.cookies.get("accessToken");
  const refresh = event.cookies.get("refreshToken");

  if (access && refresh && event.url.pathname.startsWith("/auth")) {
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/dashboard");
  }
};
