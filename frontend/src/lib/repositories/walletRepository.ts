import type { IModelRepository } from "../../protocols/modelRepository";
import { BaseRepository } from "./baseRepository";

type WalletRepositoryDTO = {
  accessToken: string | undefined;
};

export class WalletRepository
  extends BaseRepository
  implements IModelRepository
{
  constructor(input: WalletRepositoryDTO) {
    super({
      accessToken: input.accessToken,
      url: import.meta.env.VITE_API_WALLETS_ENDPOINT,
    });
  }
}
