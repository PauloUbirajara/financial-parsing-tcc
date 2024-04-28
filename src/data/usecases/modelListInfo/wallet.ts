import type { IModelListInfo } from "../../../domain/usecases/modelListInfo";

class WalletModelListInfo implements IModelListInfo {
  getListUrl(model: any): string {
    return `/app/wallets`;
  }

  getDetailUrl(model: any): string {
    return `/app/wallets/${model.id}`;
  }

  getCreateUrl(model: any): string {
    return `/app/wallets/${model.id}/create`;
  }

  getEditUrl(model: any): string {
    return `/app/wallets/${model.id}/edit`;
  }

  getDeleteUrl(model: any): string {
    return `/app/wallets/${model.id}/delete`;
  }

  getDeleteModalTitle(model: any): string {
    return `Deseja remover a carteira "${model.name}"?`;
  }
}

export default new WalletModelListInfo();
