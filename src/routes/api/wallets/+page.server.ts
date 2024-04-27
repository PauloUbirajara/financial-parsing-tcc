import type { Actions, PageServerLoad } from "./$types";
import { CurrencyRepository } from "$lib/repositories/currencyRepository";
import { WalletRepository } from "$lib/repositories/walletRepository";

export const load: PageServerLoad = async (event) => {
  try {
    const accessToken = event.cookies.get("accessToken");
    let search: string | null = event.url.searchParams.get("search");
    let page: number | null = Number(event.url.searchParams.get("page")) || 1;
    if (isNaN(page) || page <= 0) {
      page = null;
    }

    const walletResponse = await new WalletRepository({ accessToken }).getAll({
      page,
      search,
    });
    const currencyResponse = await new CurrencyRepository({
      accessToken,
    }).getAll({
      page: null,
      search: null,
    });

    return {
      walletResponse,
      currencyResponse,
    };
  } catch (e) {
    console.warn(e);
  }
};

export const actions: Actions = {
  create: async (event) => {
    const data = await event.request.json();
    const wallet = {
      name: data["name"],
      description: data["description"],
      currency: data["currency"],
    };

    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const walletRepository = new WalletRepository({ accessToken });
    const response = await walletRepository.create(wallet);
    return response;
  },

  "bulk-delete": async (event) => {
    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const walletRepository = new WalletRepository({ accessToken });

    const data = await event.request.json();
    await walletRepository.bulkDelete({ id: data });
  },
};
