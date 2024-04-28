import type { IModelRepository } from "../../protocols/modelRepository";
import { BaseRepository } from "./baseRepository";

type WalletRepositoryDTO = {
  accessToken: string | undefined;
};

type WalletRepositoryExportInput = {
  id: string;
  format: string;
};

type WalletRepositoryExportResponse = {
  url: string;
  headers: Record<string, any>;
  body: string;
};

export class WalletRepository
  extends BaseRepository
  implements IModelRepository
{
  constructor(input: WalletRepositoryDTO) {
    super({
      accessToken: input.accessToken,
      url: process.env.VITE_API_WALLETS_ENDPOINT,
    });
  }

  async export(
    input: WalletRepositoryExportInput,
  ): Promise<WalletRepositoryExportResponse> {
    let url = `${this.url}/${input.id}/export/`;

    return {
      url,
      headers: this.headers,
      body: JSON.stringify({ format: input.format }),
    };
  }
}
