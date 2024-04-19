interface WalletCurrency {
  id: string;
  name: string;
  representation: string;
  created_at: string;
  updated_at: string;
}

export interface Wallet {
  id?: string;
  name: string;
  description: string;
  currency: WalletCurrency;
}
