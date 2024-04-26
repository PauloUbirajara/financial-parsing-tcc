import { TransactionRepository } from "$lib/repositories/transactionRepository";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");

  const transactionResponse = await new TransactionRepository({
    accessToken,
  }).getAll({ page: null, search: null });

  return {
    transactionResponse,
  };
};
