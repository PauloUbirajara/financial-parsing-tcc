import { CategoryService } from "../../../services/categoryService";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {
  return {
    categories: await new CategoryService(event.cookies).getAll(),
  };
};
