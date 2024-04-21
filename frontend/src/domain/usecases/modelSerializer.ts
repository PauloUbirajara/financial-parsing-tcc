export interface IModelSerializer {
  getFields(): string[];
  serialize(model: Record<any, any>): Record<string, string>;
}
