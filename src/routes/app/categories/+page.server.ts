import type { Actions, PageServerLoad } from "./$types";
import { CategoryRepository } from "$lib/repositories/categoryRepository";

export const load: PageServerLoad = async (event) => {
  try {
    const accessToken = event.cookies.get("accessToken");
    let search: string | null = event.url.searchParams.get("search");
    let page: number | null = Number(event.url.searchParams.get("page")) || 1;
    if (isNaN(page) || page <= 0) {
      page = null;
    }

    const categoryResponse = await new CategoryRepository({
      accessToken,
    }).getAll({
      page,
      search,
    });

    return {
      categoryResponse,
    };
  } catch (e) {
    console.warn("failed when loading categories", e);
  }
};

export const actions: Actions = {
  create: async (event) => {
    const data = await event.request.json();
    const category = {
      name: data["name"],
    };

    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const categoryRepository = new CategoryRepository({ accessToken });
    const response = await categoryRepository.create(category);
    return response;
  },

  "bulk-delete": async (event) => {
    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const categoryRepository = new CategoryRepository({ accessToken });

    const data = await event.request.json();
    await categoryRepository.bulkDelete({ id: data });
  },
};
