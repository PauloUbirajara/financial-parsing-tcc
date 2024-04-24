interface TransactionWallet {
  id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

interface TransactionCategory {
  id: string;
  name: string;
}

export interface Transaction {
  id?: string;
  name: string;
  description: string;
  value: string;
  transaction_date: string;
  wallet: TransactionWallet;
  categories: TransactionCategory[];
}
