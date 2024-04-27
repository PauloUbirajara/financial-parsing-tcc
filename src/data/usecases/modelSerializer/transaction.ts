import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";

enum TransactionFieldEnum {
  NAME = "Nome",
  DESCRIPTION = "Descrição",
  WALLET = "Carteira",
  TRANSACTION_DATE = "Data da transação",
  CATEGORIES = "Categorias",
  VALUE = "Valor",
}

class TransactionModelSerializer implements IModelSerializer {
  getFields(): string[] {
    return [
      TransactionFieldEnum.NAME,
      TransactionFieldEnum.DESCRIPTION,
      TransactionFieldEnum.WALLET,
      TransactionFieldEnum.TRANSACTION_DATE,
      TransactionFieldEnum.CATEGORIES,
      TransactionFieldEnum.VALUE,
    ];
  }

  serialize(model: Record<any, any>): Record<string, string> {
    return {
      [TransactionFieldEnum.NAME]: model.name,
      [TransactionFieldEnum.DESCRIPTION]: model.description || "-",
      [TransactionFieldEnum.WALLET]: `${model.wallet.name}`,
      [TransactionFieldEnum.TRANSACTION_DATE]: `${model.transaction_date}`,
      [TransactionFieldEnum.CATEGORIES]: `${model.categories.map((c: { id: string; name: string }) => c["name"]).join(", ")}`,
      [TransactionFieldEnum.VALUE]: `${model.value}`,
    };
  }
}

export default new TransactionModelSerializer();
