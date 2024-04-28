import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";

class TransactionModelListInfo implements IModelListInfo {
  getListUrl(model: any): string {
    return `/app/transactions`;
  }

  getDetailUrl(model: any): string {
    return `/app/transactions/${model.id}`;
  }

  getCreateUrl(model: any): string {
    return `/app/transactions/${model.id}/create`;
  }

  getEditUrl(model: any): string {
    return `/app/transactions/${model.id}/edit`;
  }

  getDeleteUrl(model: any): string {
    return `/app/transactions/${model.id}/delete`;
  }

  getDeleteModalTitle(model: any): string {
    return `Deseja remover a transação "${model.name}"?`;
  }
}

export default new TransactionModelListInfo();
