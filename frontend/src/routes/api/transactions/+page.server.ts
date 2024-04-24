import type { Actions, PageServerLoad } from "./$types";
import { WalletRepository } from "$lib/repositories/walletRepository";
import { TransactionRepository } from "$lib/repositories/transactionRepository";
import { CategoryRepository } from "$lib/repositories/categoryRepository";

export const load: PageServerLoad = async (event) => {
  try {
    const accessToken = event.cookies.get("accessToken");
    let search: string | null = event.url.searchParams.get("search");
    let page: number | null = Number(event.url.searchParams.get("page")) || 1;
    if (isNaN(page) || page <= 0) {
      page = null;
    }

    const transactionResponse = await new TransactionRepository({
      accessToken,
    }).getAll({
      page,
      search,
    });
    const walletResponse = await new WalletRepository({
      accessToken,
    }).getAll({
      page: null,
      search: null,
    });
    const categoryResponse = await new CategoryRepository({
      accessToken,
    }).getAll({
      page: null,
      search: null,
    });

    return {
      transactionResponse,
      walletResponse,
      categoryResponse,
    };
  } catch (e) {
    console.warn(e);
  }
};

export const actions: Actions = {
  create: async (event) => {
    const data = await event.request.json();
    const transaction = {
      name: data["name"],
      description: data["description"],
      wallet: data["wallet"],
      categories: data["categories"],
      value: data["value"],
      transaction_date: data["transaction_date"],
    };

    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const transactionRepository = new TransactionRepository({ accessToken });
    const response = await transactionRepository.create(transaction);
    return response;
  },

  "bulk-delete": async (event) => {
    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const transactionRepository = new TransactionRepository({ accessToken });

    const data = await event.request.json();
    await transactionRepository.bulkDelete({ id: data });
  },
};
