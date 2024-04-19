import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {
  const isLogged = event.cookies.get("accessToken");
  return { isLogged };
};
