import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type {
  GetAllModelsRepositoryResponse,
  GetModelByIdRepositoryResponse,
} from "../../../../../domain/models/modelRepositoryDto";
import { constants } from "http2";
import { CurrencyRepository } from "$lib/repositories/currencyRepository";
import { fail, redirect } from "@sveltejs/kit";
import { CategoryRepository } from "$lib/repositories/categoryRepository";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");

  let categoryResponse: GetModelByIdRepositoryResponse | null = null;

  try {
    categoryResponse = await new CategoryRepository({ accessToken }).getById({
      id: event.params.id,
    });
  } catch (e) {
    console.warn("loading", e);
  }

  if (categoryResponse === null) {
    console.warn("Could not edit category by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/categories");
  }

  return { categoryResponse };
};

export const actions: Actions = {
  default: async (event) => {
    const id = event.params.id;
    const accessToken = event.cookies.get("accessToken");

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        error: "Não foi possível editar categoria com ID inválido.",
      });
    }

    const formData = await event.request.formData();
    const data = Object.fromEntries(formData.entries());
    const updated = {
      name: data["name"],
    };
    await new CategoryRepository({ accessToken }).updateById({ id, updated });
  },
};
