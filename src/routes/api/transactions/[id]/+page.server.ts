import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { constants } from "http2";
import { fail, redirect } from "@sveltejs/kit";
import { TransactionRepository } from "$lib/repositories/transactionRepository";
import type { GetModelByIdRepositoryResponse } from "../../../../domain/models/modelRepositoryDto";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");
  const transactionRepository = new TransactionRepository({ accessToken });

  let transaction: GetModelByIdRepositoryResponse | null = null;
  try {
    transaction = await transactionRepository.getById({ id: event.params.id });
  } catch (e) {
    console.warn("loading", e);
  }

  if (transaction === null) {
    console.warn("Could not get transaction by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/transactions");
  }

  return { transaction: transaction };
};

export const actions: Actions = {
  delete: async (event) => {
    const accessToken = event.cookies.get("accessToken");
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, { error: "ID inválido." });
    }

    try {
      await new TransactionRepository({ accessToken }).deleteById({ id });
    } catch (e) {
      console.warn(e);
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        error: "Erro ao remover transação",
      });
    }
  },
};
