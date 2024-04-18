import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async (event) => {
  const access = event.cookies.get("accessToken");
  const refresh = event.cookies.get("refreshToken");

  const isLogged = access !== undefined && refresh !== undefined;
  return {
    isLogged,
  };
};
