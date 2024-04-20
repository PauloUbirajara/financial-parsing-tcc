import type { Currency } from "../../../domain/models/currency";
import type { Wallet } from "../../../domain/models/wallet";
import { CurrencyRepository } from "../../../repositories/currencyRepository";
import { WalletRepository } from "../../../repositories/walletRepository";

import type { PageServerLoad } from "./$types";
import type { Actions } from "@sveltejs/kit";

export const load: PageServerLoad = async (event) => {
  const walletRepository = new WalletRepository(event.cookies);
  const currencyRepository = new CurrencyRepository(event.cookies);

  let page: number = Number(event.url.searchParams.get("page") || 1);
  if (isNaN(page)) {
    page = 1;
  }

  const currencies = await currencyRepository.getAll({ page: null });
  const wallets = await walletRepository.getAll({ page });

  return {
    currencies: currencies,
    wallets: wallets,
  };
};

export const actions: Actions = {
  create: async (event) => {
    const formData = await event.request.formData();
    const data = Object.fromEntries(formData.entries());
    const wallet = {
      name: data["name"],
      description: data["description"],
      currency: data["currency"],
    };
    const walletRepository = new WalletRepository(event.cookies);
    const response = await walletRepository.create(wallet);
    return response;
  },
};
