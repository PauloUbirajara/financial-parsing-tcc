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

  async getById(
    input: GetModelByIdRepositoryInput,
  ): Promise<GetModelByIdRepositoryResponse> {
    let url = `${this.url}/${input.id}`;

    const response = await fetch(url, {
      headers: this.headers,
    });

    const data: GetModelByIdRepositoryResponse = await response.json();
    return data;
  }

  async create(
    input: CreateModelsRepositoryInput,
  ): Promise<CreateModelsRepositoryResponse> {
    let url = `${this.url}/`;

    const response = await fetch(url, {
      headers: this.headers,
      method: "POST",
      body: JSON.stringify(input),
    });

    const data: CreateModelsRepositoryResponse = await response.json();
    return data;
  }

  async updateById(
    input: UpdateModelByIdRepositoryInput,
  ): Promise<UpdateModelByIdRepositoryResponse> {
    let url = `${this.url}/${input.id}/`;

    const response = await fetch(url, {
      headers: this.headers,
      method: "PUT",
      body: JSON.stringify(input.updated),
    });

    const data: UpdateModelByIdRepositoryResponse = await response.json();
    return data;
  }

  async deleteById(
    input: DeleteModelByIdRepositoryInput,
  ): Promise<DeleteModelByIdRepositoryResponse> {
    let url = `${this.url}/${input.id}/`;

    const response = await fetch(url, {
      headers: this.headers,
      method: "DELETE",
    });

    if (!response.ok) {
      return Promise.reject(await response.json());
    }

    return;
  }

  bulkDelete(
    input: BulkDeleteModelRepositoryInput,
  ): Promise<BulkDeleteModelRepositoryResponse> {
    throw new Error("Method not implemented.");
  }
}
