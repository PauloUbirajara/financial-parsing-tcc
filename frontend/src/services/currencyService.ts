import type { Cookies } from "@sveltejs/kit";
import type { Currency } from "../domain/models/currency";
import type { IModelService } from "../protocols/modelService";

export class CurrencyService implements IModelService<Currency> {
  private currenciesEndpoint: string;
  private headers: Record<string, any>;

  constructor(cookies: Cookies) {
    if (cookies.get("accessToken") === undefined) {
      throw new Error("Could not get access token");
    }
    this.headers = {
      Authorization: `Bearer ${cookies.get("accessToken")!}`,
    };

    this.currenciesEndpoint = import.meta.env.VITE_API_CURRENCIES_ENDPOINT;
    if (this.currenciesEndpoint === undefined) {
      throw new Error("Currencies endpoint not set");
    }
  }

  async getAll(): Promise<Currency[]> {
    const url = this.currenciesEndpoint;
    try {
      const response = await fetch(url, { headers: this.headers });
      const currencies = await response.json();
      return currencies;
    } catch (e) {
      console.warn("Error when getting currencies", e);
    }
    return [];
  }

  async getById(id: string): Promise<Currency | null> {
    const url = `${this.currenciesEndpoint}/${id}`;
    try {
      const response = await fetch(url, { headers: this.headers });
      const currency = await response.json();
      return currency;
    } catch (e) {
      console.warn("Error when getting currency by id", e);
    }
    return null;
  }

  updateById(id: string, updated: Currency): Promise<Currency> {
    throw new Error("Method not implemented.");
  }

  async deleteById(id: string): Promise<void> {
    throw new Error("Method not implemented.");
  }

  async create(model: Currency): Promise<Currency | null> {
    throw new Error("Method not implemented.");
  }
}
