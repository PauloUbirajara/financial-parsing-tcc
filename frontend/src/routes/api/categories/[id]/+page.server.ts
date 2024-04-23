import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { constants } from "http2";
import { fail, redirect } from "@sveltejs/kit";
import { CategoryRepository } from "$lib/repositories/categoryRepository";
import type { GetModelByIdRepositoryResponse } from "../../../../domain/models/modelRepositoryDto";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");
  const categoryRepository = new CategoryRepository({ accessToken });

  let category: GetModelByIdRepositoryResponse | null = null;
  try {
    category = await categoryRepository.getById({ id: event.params.id });
  } catch (e) {
    console.warn("loading", e);
  }

  if (category === null) {
    console.warn("Could not get category by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/categories");
  }

  return { category: category };
};

export const actions: Actions = {
  delete: async (event) => {
    const accessToken = event.cookies.get("accessToken");
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, { error: "ID inválido." });
    }

    try {
      await new CategoryRepository({ accessToken }).deleteById({ id });
    } catch (e) {
      console.warn(e);
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        error: "Erro ao remover categoria",
      });
    }
  },
};
