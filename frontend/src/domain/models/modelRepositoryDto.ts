// Create
export type CreateModelsRepositoryInput = Record<string, any>;

export type CreateModelsRepositoryResponse = Record<string, any>;

// Get All
export type GetAllModelsRepositoryInput = {
  page: number | null;
  search: string | null;
};

export type GetAllModelsRepositoryResponse = {
  links: {
    next: string;
    previous: string;
  };
  num_pages: number;
  count: number;
  results: Record<string, any>[];
};

// Get By ID
export type GetModelByIdRepositoryInput = {
  id: string;
};

export type GetModelByIdRepositoryResponse = Record<string, any> | null;

// Update By ID
export type UpdateModelByIdRepositoryInput = {
  id: string;
  updated: Record<string, any>;
};

export type UpdateModelByIdRepositoryResponse = Record<string, any> | null;

// Delete By ID
export type DeleteModelByIdRepositoryInput = {
  id: string;
};

export type DeleteModelByIdRepositoryResponse = void;

// Bulk Delete
export type BulkDeleteModelRepositoryInput = {
  id: string[];
};

export type BulkDeleteModelRepositoryResponse = void;
