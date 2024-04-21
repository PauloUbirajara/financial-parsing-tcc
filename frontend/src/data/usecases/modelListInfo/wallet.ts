import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";

class WallelModelListInfo implements IModelListInfo {
  getDetailUrl(model: any): string {
    if (model === undefined) {
      return "";
    }
    return `/api/wallets/${model.id}`;
  }

  getCreateUrl(model: any): string {
    if (model === undefined) {
      return "";
    }
    return `/api/wallets/${model.id}/create`;
  }

  getEditUrl(model: any): string {
    if (model === undefined) {
      return "";
    }
    return `/api/wallets/${model.id}/edit`;
  }

  getDeleteUrl(model: any): string {
    if (model === undefined) {
      return "";
    }
    return `/api/wallets/${model.id}/delete`;
  }

  getDeleteModalTitle(model: any): string {
    if (model === undefined) {
      return "";
    }
    return `Deseja remover a carteira "${model.name}"?`;
  }
}

export default new WallelModelListInfo();
