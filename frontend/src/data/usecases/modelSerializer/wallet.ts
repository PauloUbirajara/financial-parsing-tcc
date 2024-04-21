import type { IModelSerializer } from "../../../domain/usecases/modelSerializer";

enum WalletFieldEnum {
  NAME = "Nome",
  DESCRIPTION = "Descrição",
  CURRENCY = "Moeda",
}

class WalletModelSerializer implements IModelSerializer {
  getFields(): string[] {
    return [
      WalletFieldEnum.NAME,
      WalletFieldEnum.DESCRIPTION,
      WalletFieldEnum.CURRENCY,
    ];
  }

  serialize(model: Record<any, any>): Record<string, string> {
    return {
      [WalletFieldEnum.NAME]: model.name,
      [WalletFieldEnum.DESCRIPTION]: model.description || "-",
      [WalletFieldEnum.CURRENCY]: `${model.currency.name} (${model.currency.representation})`,
    };
  }
}

export default new WalletModelSerializer();
