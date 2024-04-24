import type {
  BulkDeleteModelRepositoryInput,
  CreateModelsRepositoryInput,
  CreateModelsRepositoryResponse,
  DeleteModelByIdRepositoryInput,
  UpdateModelByIdRepositoryInput,
  UpdateModelByIdRepositoryResponse,
} from "../../domain/models/modelRepositoryDto";
import type { IModelRepository } from "../../protocols/modelRepository";
import { BaseRepository } from "./baseRepository";

type CurrencyRepositoryDTO = {
  accessToken: string | undefined;
};

export class CurrencyRepository
  extends BaseRepository
  implements IModelRepository
{
  constructor(input: CurrencyRepositoryDTO) {
    super({
      accessToken: input.accessToken,
      url: import.meta.env.VITE_API_CURRENCIES_ENDPOINT,
    });
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
