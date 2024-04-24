// Used when getting strings to be shown in the svelte list component
export interface IModelListInfo {
  getCreateUrl(model: any): string;
  getListUrl(model: any): string;
  getDetailUrl(model: any): string;
  getEditUrl(model: any): string;
  getDeleteUrl(model: any): string;
  getDeleteModalTitle(model: any): string;
}
