import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { constants } from "http2";
import AuthManager from "$lib/auth/AuthManager";

export const load: LayoutServerLoad = async (event) => {
  const access = event.cookies.get("accessToken");

  if (!access) {
    event.cookies.delete("accessToken", { path: "/" });
    event.cookies.delete("refreshToken", { path: "/" });
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/auth/login");
  }

  const isAuthorized = await AuthManager.validate(access);
  if (!isAuthorized && event.url.pathname.startsWith("/auth")) {
    event.cookies.delete("accessToken", { path: "/" });
    event.cookies.delete("refreshToken", { path: "/" });
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/auth/login");
  }
};
