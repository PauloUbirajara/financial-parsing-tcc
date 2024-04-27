export enum ToastType {
  SUCCESS,
  ERROR,
  WARNING,
}

export type ToastMessage = {
  title: string;
  message: string;
  type: ToastType;
};
