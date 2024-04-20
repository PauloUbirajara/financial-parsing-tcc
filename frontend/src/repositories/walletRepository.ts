import type { Cookies } from "@sveltejs/kit";
import type {
  BulkDeleteModelRepositoryInput,
  BulkDeleteModelRepositoryResponse,
  CreateModelsRepositoryInput,
  CreateModelsRepositoryResponse,
  DeleteModelByIdRepositoryInput,
  DeleteModelByIdRepositoryResponse,
  GetAllModelsRepositoryInput,
  GetAllModelsRepositoryResponse,
  GetModelByIdRepositoryInput,
  GetModelByIdRepositoryResponse,
  UpdateModelByIdRepositoryInput,
  UpdateModelByIdRepositoryResponse,
} from "../domain/models/modelRepositoryDto";
import type { IModelRepository } from "../protocols/modelRepository";

export class WalletRepository implements IModelRepository {
  private headers: Record<string, any>;
  private url: string;

  constructor(cookies: Cookies) {
    const accessToken = cookies.get("accessToken");
    this.url = import.meta.env.VITE_API_WALLETS_ENDPOINT;

    this.headers = {
      "Content-Type": "application/json",
      Authorization: `Bearer ${accessToken}`,
    };
  }

  async getAll(
    input: GetAllModelsRepositoryInput,
  ): Promise<GetAllModelsRepositoryResponse> {
    let url = this.url;
    if (input.page !== null && input.page > 0) {
      url = `${this.url}?page=${input.page}`;
    }

    const response = await fetch(url, {
      headers: this.headers,
    });

    const data: GetAllModelsRepositoryResponse = await response.json();
    return data;
  }

  getById(
    input: GetModelByIdRepositoryInput,
  ): Promise<GetModelByIdRepositoryResponse> {
    throw new Error("Method not implemented.");
  }

  create(
    input: CreateModelsRepositoryInput,
  ): Promise<CreateModelsRepositoryResponse> {
    throw new Error("Method not implemented.");
  }

  updateById(
    input: UpdateModelByIdRepositoryInput,
  ): Promise<UpdateModelByIdRepositoryResponse> {
    throw new Error("Method not implemented.");
  }

  deleteById(
    input: DeleteModelByIdRepositoryInput,
  ): Promise<DeleteModelByIdRepositoryResponse> {
    throw new Error("Method not implemented.");
  }

  bulkDelete(
    input: BulkDeleteModelRepositoryInput,
  ): Promise<BulkDeleteModelRepositoryResponse> {
    throw new Error("Method not implemented.");
  }
}
