import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { constants } from "http2";
import { fail, redirect } from "@sveltejs/kit";
import { WalletRepository } from "$lib/repositories/walletRepository";
import type { GetModelByIdRepositoryResponse } from "../../../../domain/models/modelRepositoryDto";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");
  const walletRepository = new WalletRepository({ accessToken });

  let wallet: GetModelByIdRepositoryResponse | null = null;
  try {
    wallet = await walletRepository.getById({ id: event.params.id });
  } catch (e) {
    console.warn("loading", e);
  }

  if (wallet === null) {
    console.warn("Could not get wallet by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/app/wallets");
  }

  return { wallet: wallet };
};

export const actions: Actions = {
  delete: async (event) => {
    const accessToken = event.cookies.get("accessToken");
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, { error: "ID inválido." });
    }

    try {
      await new WalletRepository({ accessToken }).deleteById({ id });
    } catch (e) {
      console.warn("failed when removing wallet", e);
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        error: "Erro ao remover carteira",
      });
    }
  },
  export: async (event) => {
    const accessToken = event.cookies.get("accessToken");
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, { error: "ID inválido." });
    }

    try {
      const data = await event.request.json();
      const format = data["format"];
      return await new WalletRepository({ accessToken }).export({ id, format });
    } catch (e) {
      console.warn("Could not get credentials for export fetch call", e);
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        error: "Erro ao exportar carteira",
      });
    }
  },
};
