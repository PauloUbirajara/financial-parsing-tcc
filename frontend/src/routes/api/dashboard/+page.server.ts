import { CategoryRepository } from "$lib/repositories/categoryRepository";
import { CurrencyRepository } from "$lib/repositories/currencyRepository";
import { TransactionRepository } from "$lib/repositories/transactionRepository";
import { WalletRepository } from "$lib/repositories/walletRepository";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");

  const transactionResponse = await new TransactionRepository({
    accessToken,
  }).getAll({ page: null, search: null });
  const walletResponse = await new WalletRepository({ accessToken }).getAll({
    page: null,
    search: null,
  });
  const currencyResponse = await new CurrencyRepository({ accessToken }).getAll(
    { page: null, search: null },
  );
  const categoryResponse = await new CategoryRepository({ accessToken }).getAll(
    { page: null, search: null },
  );

  return {
    transactionResponse,
    currencyResponse,
    categoryResponse,
    walletResponse,
  };
};
