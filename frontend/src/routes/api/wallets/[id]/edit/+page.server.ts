import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type {
  GetAllModelsRepositoryResponse,
  GetModelByIdRepositoryResponse,
} from "../../../../../domain/models/modelRepositoryDto";
import { constants } from "http2";
import { CurrencyRepository } from "$lib/repositories/currencyRepository";
import { fail, redirect } from "@sveltejs/kit";
import { WalletRepository } from "$lib/repositories/walletRepository";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");

  let walletResponse: GetModelByIdRepositoryResponse | null = null;
  let currencyResponse: GetAllModelsRepositoryResponse | null = null;

  try {
    walletResponse = await new WalletRepository({ accessToken }).getById({
      id: event.params.id,
    });
    currencyResponse = await new CurrencyRepository({ accessToken }).getAll({
      page: null,
      search: null,
    });
  } catch (e) {
    console.warn("loading", e);
  }

  if (walletResponse === null) {
    console.warn("Could not edit wallet by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/wallets");
  }

  if (currencyResponse === null) {
    console.warn(
      "Could not get currencies for editing wallet by id",
      event.params.id,
    );
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/wallets");
  }

  return { walletResponse, currencyResponse };
};

export const actions: Actions = {
  default: async (event) => {
    const id = event.params.id;
    const accessToken = event.cookies.get("accessToken");

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        error: "Não foi possível editar carteira com ID inválido.",
      });
    }

    const data = await event.request.json();
    const updated = {
      name: data["name"],
      description: data["description"],
      currency: data["currency"],
    };
    await new WalletRepository({ accessToken }).updateById({ id, updated });
  },
};
