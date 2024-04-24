import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";

class TransactionModelListInfo implements IModelListInfo {
  getDetailUrl(model: any): string {
    return `/api/transactions/${model.id}`;
  }

  getCreateUrl(model: any): string {
    return `/api/transactions/${model.id}/create`;
  }

  getEditUrl(model: any): string {
    return `/api/transactions/${model.id}/edit`;
  }

  getDeleteUrl(model: any): string {
    return `/api/transactions/${model.id}/delete`;
  }

  getDeleteModalTitle(model: any): string {
    return `Deseja remover a transação "${model.name}"?`;
  }
}

export default new TransactionModelListInfo();
