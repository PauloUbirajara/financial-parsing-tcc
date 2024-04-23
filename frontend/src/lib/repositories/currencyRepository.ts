import type {
  BulkDeleteModelRepositoryInput,
  CreateModelsRepositoryInput,
  CreateModelsRepositoryResponse,
  DeleteModelByIdRepositoryInput,
  GetAllModelsRepositoryInput,
  GetAllModelsRepositoryResponse,
  GetModelByIdRepositoryInput,
  GetModelByIdRepositoryResponse,
  UpdateModelByIdRepositoryInput,
  UpdateModelByIdRepositoryResponse,
} from "../../domain/models/modelRepositoryDto";
import { getFilteredUrlSearchParams } from "../../helpers/url";
import type { IModelRepository } from "../../protocols/modelRepository";

type CurrencyRepositoryDTO = {
  accessToken: string | undefined;
};

export class CurrencyRepository implements IModelRepository {
  private headers: Record<string, any>;
  private url: string;

  constructor(input: CurrencyRepositoryDTO) {
    if (input.accessToken === undefined) {
      throw new Error(
        "Could not setup currency repository due to undefined access token",
      );
    }

    this.url = import.meta.env.VITE_API_CURRENCIES_ENDPOINT;

    this.headers = {
      "Content-Type": "application/json",
      Authorization: `Bearer ${input.accessToken}`,
    };
  }

  async getAll(
    input: GetAllModelsRepositoryInput,
  ): Promise<GetAllModelsRepositoryResponse> {
    let url = this.url;

    // Filtering
    let queryUrl = getFilteredUrlSearchParams({
      page: input.page,
      search: input.search,
    });
    if (queryUrl.toString()) {
      url = `${this.url}?${queryUrl.toString()}`;
    }

    const response = await fetch(url, {
      headers: this.headers,
    });

    if (!response.ok) {
      console.warn("Could not get currencies using url search params");
      return {
        links: { next: "", previous: "" },
        num_pages: 0,
        count: 0,
        results: [],
      };
    }

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
  deleteById(input: DeleteModelByIdRepositoryInput): Promise<void> {
    throw new Error("Method not implemented.");
  }
  bulkDelete(input: BulkDeleteModelRepositoryInput): Promise<void> {
    throw new Error("Method not implemented.");
  }
}
