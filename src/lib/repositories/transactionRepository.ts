import type { IModelRepository } from "../../protocols/modelRepository";
import { BaseRepository } from "./baseRepository";

type TransactionRepositoryDTO = {
  accessToken: string | undefined;
};

export class TransactionRepository
  extends BaseRepository
  implements IModelRepository
{
  constructor(input: TransactionRepositoryDTO) {
    super({
      accessToken: input.accessToken,
      url: process.env.VITE_API_TRANSACTIONS_ENDPOINT,
    });
  }
}
