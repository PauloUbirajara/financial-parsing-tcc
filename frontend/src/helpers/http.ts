import type { Cookies } from "@sveltejs/kit";

export function getHeaders(cookies: Cookies): Record<any, any> {
  if (cookies.get("accessToken") === undefined) {
    throw new Error("Could not get access token");
  }

  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${cookies.get("accessToken")!}`,
  };

  return headers;
}
