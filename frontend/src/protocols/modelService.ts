export interface IModelService<T> {
  getAll(): Promise<T[]>;
  getById(): Promise<T>;
  updateById(id: string, updated: T): Promise<T>;
  deleteById(id: string): Promise<T>;
}
