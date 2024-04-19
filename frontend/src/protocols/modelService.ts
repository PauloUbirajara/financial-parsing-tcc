export interface IModelService<T> {
  getAll(): Promise<T[]>;
  getById(id: string): Promise<T | null>;
  create(model: Record<any, any>): Promise<T | null>;
  updateById(id: string, updated: Record<any, any>): Promise<T | null>;
  deleteById(id: string): Promise<void>;
}
