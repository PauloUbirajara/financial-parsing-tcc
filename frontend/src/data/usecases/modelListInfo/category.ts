import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";

class CategoryModelListInfo implements IModelListInfo {
  getDetailUrl(model: any): string {
    return `/api/categories/${model.id}`;
  }

  getCreateUrl(model: any): string {
    return `/api/categories/${model.id}/create`;
  }

  getEditUrl(model: any): string {
    return `/api/categories/${model.id}/edit`;
  }

  getDeleteUrl(model: any): string {
    return `/api/categories/${model.id}/delete`;
  }

  getDeleteModalTitle(model: any): string {
    return `Deseja remover a categoria "${model.name}"?`;
  }
}

export default new CategoryModelListInfo();
