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

export interface IModelRepository {
  getAll(
    input: GetAllModelsRepositoryInput,
  ): Promise<GetAllModelsRepositoryResponse>;
  create(
    input: CreateModelsRepositoryInput,
  ): Promise<CreateModelsRepositoryResponse>;
  getById(
    input: GetModelByIdRepositoryInput,
  ): Promise<GetModelByIdRepositoryResponse>;
  updateById(
    input: UpdateModelByIdRepositoryInput,
  ): Promise<UpdateModelByIdRepositoryResponse>;
  deleteById(
    input: DeleteModelByIdRepositoryInput,
  ): Promise<DeleteModelByIdRepositoryResponse>;
  bulkDelete(
    input: BulkDeleteModelRepositoryInput,
  ): Promise<BulkDeleteModelRepositoryResponse>;
}
