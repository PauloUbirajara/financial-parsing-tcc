import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");
  return { accessToken };
};
