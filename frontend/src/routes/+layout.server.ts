import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { constants } from "http2";
import AuthManager from "$lib/auth/AuthManager";

export const load: LayoutServerLoad = async (event) => {
  let isAuthorized = false;

  const access = event.cookies.get("accessToken");
  if (access !== undefined) {
    isAuthorized = await AuthManager.validate(access);
  }

  if (!isAuthorized) {
    event.cookies.delete("accessToken", { path: "/" });
    event.cookies.delete("refreshToken", { path: "/" });

    if (event.url.pathname.startsWith("/api")) {
      redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/auth/login");
    }
  }

  if (isAuthorized && event.url.pathname.startsWith("/auth")) {
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/dashboard");
  }
};