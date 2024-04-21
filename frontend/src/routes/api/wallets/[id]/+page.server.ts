import type { Wallet } from "../../../../domain/models/wallet";
import { WalletRepository } from "../../../../repositories/walletRepository";
import { CurrencyRepository } from "../../../../repositories/currencyRepository";

import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { constants } from "http2";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async (event) => {
  const walletRepository = new WalletRepository(event.cookies);

  let wallet: any;
  try {
    wallet = await walletRepository.getById({ id: event.params.id });
    console.log({ wallet });
  } catch (e) {
    console.warn("loading", wallet, e);
  }

  if (wallet === null) {
    console.warn("Could not get wallet by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/wallets");
  }

  return { wallet: wallet };
};

export const actions: Actions = {
  delete: async (event) => {
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, { error: "ID inválido." });
    }

    try {
      await new WalletRepository(event.cookies).deleteById({ id });
    } catch (e) {
      console.warn(e);
      return fail(constants.HTTP_STATUS_INTERNAL_SERVER_ERROR, {
        error: "Erro ao remover carteira",
      });
    }
  },
};
