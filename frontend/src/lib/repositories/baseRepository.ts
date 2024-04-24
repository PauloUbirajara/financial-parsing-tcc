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
} from "../../domain/models/modelRepositoryDto";
import { getFilteredUrlSearchParams } from "../../helpers/url";
import type { IModelRepository } from "../../protocols/modelRepository";

type BaseRepositoryDTO = {
  accessToken: string | undefined;
  url: string;
};

export abstract class BaseRepository implements IModelRepository {
  private headers: Record<string, any>;
  private url: string;

  constructor(input: BaseRepositoryDTO) {
    if (input.accessToken === undefined) {
      throw new Error(
        "Could not setup repository due to undefined access token",
      );
    }

    this.url = input.url;

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
      console.warn("Could not get models using url search params");
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

  async bulkDelete(
    input: BulkDeleteModelRepositoryInput,
  ): Promise<BulkDeleteModelRepositoryResponse> {
    let url = `${this.url}/bulk_delete/`;

    const response = await fetch(url, {
      headers: this.headers,
      method: "POST",
      body: JSON.stringify({ ids: input.id }),
    });

    if (!response.ok) {
      return Promise.reject(await response.json());
    }

    return;
  }
}
