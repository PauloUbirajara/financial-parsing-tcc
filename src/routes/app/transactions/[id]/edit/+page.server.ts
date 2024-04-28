import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type {
  GetAllModelsRepositoryResponse,
  GetModelByIdRepositoryResponse,
} from "../../../../../domain/models/modelRepositoryDto";
import { constants } from "http2";
import { WalletRepository } from "$lib/repositories/walletRepository";
import { fail, redirect } from "@sveltejs/kit";
import { TransactionRepository } from "$lib/repositories/transactionRepository";
import { CategoryRepository } from "$lib/repositories/categoryRepository";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");

  let transactionResponse: GetModelByIdRepositoryResponse | null = null;
  let walletResponse: GetAllModelsRepositoryResponse | null = null;
  let categoryResponse: GetAllModelsRepositoryResponse | null = null;

  try {
    transactionResponse = await new TransactionRepository({
      accessToken,
    }).getById({
      id: event.params.id,
    });
    walletResponse = await new WalletRepository({ accessToken }).getAll({
      page: null,
      search: null,
    });
    categoryResponse = await new CategoryRepository({ accessToken }).getAll({
      page: null,
      search: null,
    });
  } catch (e) {
    console.warn("loading", e);
  }

  if (transactionResponse === null) {
    console.warn("Could not edit transaction by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/app/transactions");
  }

  if (walletResponse === null) {
    console.warn(
      "Could not get wallets for editing transaction by id",
      event.params.id,
    );
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/app/transactions");
  }

  if (categoryResponse === null) {
    console.warn(
      "Could not get categories for editing transaction by id",
      event.params.id,
    );
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/app/transactions");
  }

  return { transactionResponse, walletResponse, categoryResponse };
};

export const actions: Actions = {
  default: async (event) => {
    const id = event.params.id;
    const accessToken = event.cookies.get("accessToken");

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        error: "Não foi possível editar transação com ID inválido.",
      });
    }

    const data = await event.request.json();
    const updated = {
      name: data["name"],
      description: data["description"],
      wallet: data["wallet"],
      categories: data["categories"],
      value: data["value"],
      transaction_date: data["transaction_date"],
    };
    await new TransactionRepository({ accessToken }).updateById({
      id,
      updated,
    });
  },
};
