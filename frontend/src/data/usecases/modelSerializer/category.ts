import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";

enum CategoryFieldEnum {
  NAME = "Nome",
}

class CategoryModelSerializer implements IModelSerializer {
  getFields(): string[] {
    return [CategoryFieldEnum.NAME];
  }

  serialize(model: Record<any, any>): Record<string, string> {
    return {
      [CategoryFieldEnum.NAME]: model.name,
    };
  }
}

export default new CategoryModelSerializer();
