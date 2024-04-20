import type { Currency } from "../../../../../domain/models/currency";
import type { Wallet } from "../../../../../domain/models/wallet";
import { CurrencyService } from "../../../../../services/currencyService";
import { WalletService } from "../../../../../services/walletService";

import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { constants } from "http2";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async (event) => {
  let wallet: Wallet | null = null;
  let currencies: Currency[] | null = null;

  try {
    wallet = await new WalletService(event.cookies).getById(event.params.id);
    currencies = await new CurrencyService(event.cookies).getAll();
  } catch (e) {
    console.warn("loading", wallet, e);
  }

  if (wallet === null) {
    console.warn("Could not edit wallet by id", event.params.id);
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/wallets");
  }

  return { wallet, currencies };
};

export const actions: Actions = {
  default: async (event) => {
    const id = event.params.id;

    if (id === undefined) {
      return fail(constants.HTTP_STATUS_BAD_REQUEST, {
        error: "Não foi possível editar carteira com ID inválido.",
      });
    }

    const formData = await event.request.formData();
    const data = Object.fromEntries(formData.entries());
    const updated = {
      name: data["name"],
      description: data["description"],
      currency: data["currency"],
    };
    await new WalletService(event.cookies).updateById(id, updated);
  },
};
