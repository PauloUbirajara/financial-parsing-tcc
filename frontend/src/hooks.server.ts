import { redirect } from "@sveltejs/kit";
import type { Handle, RequestEvent } from "@sveltejs/kit";
import AuthManager from "$lib/auth/AuthManager";
import { constants } from "http2";

async function handleJwtAuthorization(
  event: RequestEvent<Partial<Record<string, string>>, string | null>,
) {
  let isAuthorized = false;

  let access = event.cookies.get("accessToken");
  if (access !== undefined) {
    isAuthorized = await AuthManager.validate(access);
  }

  const refresh = event.cookies.get("refreshToken");
  if (refresh !== undefined && !isAuthorized) {
    const newAccess = await AuthManager.refresh(refresh);

    if (newAccess !== null) {
      access = newAccess;
      event.cookies.set("accessToken", access, { path: "/" });
      isAuthorized = true;
    }
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
}

export const handle: Handle = async ({ event, resolve }) => {
  await handleJwtAuthorization(event);

  const response = await resolve(event);
  return response;
};
