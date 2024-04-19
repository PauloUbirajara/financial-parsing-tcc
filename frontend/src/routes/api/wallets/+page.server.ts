import { CurrencyService } from "../../../services/currencyService";
import { WalletService } from "../../../services/walletService";

import type { PageServerLoad } from "./$types";
import type { Actions } from "@sveltejs/kit";

export const load: PageServerLoad = async (event) => {
  const walletService = new WalletService(event.cookies);
  const currencyService = new CurrencyService(event.cookies);

  return {
    currencies: await currencyService.getAll(),
    wallets: await walletService.getAll(),
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
    const response = await new WalletService(event.cookies).create(wallet);
    return response;
  },
};
