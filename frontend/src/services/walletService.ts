import type { Cookies } from "@sveltejs/kit";
import type { Wallet } from "../domain/models/wallet";
import type { IModelService } from "../protocols/modelService";

export class WalletService implements IModelService<Wallet> {
  private walletsEndpoint: string;
  private headers: Record<string, any>;

  constructor(cookies: Cookies) {
    if (cookies.get("accessToken") === undefined) {
      throw new Error("Could not get access token");
    }
    this.headers = {
      "Content-Type": "application/json",
      Authorization: `Bearer ${cookies.get("accessToken")!}`,
    };

    this.walletsEndpoint = import.meta.env.VITE_API_WALLETS_ENDPOINT;
    if (this.walletsEndpoint === undefined) {
      throw new Error("Wallets endpoint not set");
    }
  }

  async getAll(): Promise<Wallet[]> {
    const url = this.walletsEndpoint;
    try {
      const response = await fetch(url, { headers: this.headers });
      const wallets = await response.json();
      return wallets;
    } catch (e) {
      console.warn("Error when getting wallets", e);
    }
    return [];
  }

  async getById(id: string): Promise<Wallet | null> {
    const url = `${this.walletsEndpoint}/${id}`;
    try {
      const response = await fetch(url, { headers: this.headers });
      const wallet = await response.json();
      return wallet;
    } catch (e) {
      console.warn("Error when getting wallet by id", e);
    }
    return null;
  }

  updateById(id: string, updated: Record<any, any>): Promise<Wallet> {
    throw new Error("Method not implemented.");
  }

  async deleteById(id: string): Promise<void> {
    const url = `${this.walletsEndpoint}/${id}/delete`;

    try {
      await fetch(url, {
        headers: this.headers,
        method: "DELETE",
      });
      return;
    } catch (e) {
      console.warn("Error when deleting wallet by id", e);
    }

    return Promise.reject("Não foi possível remover carteira.");
  }

  async create(model: Record<any, any>): Promise<Wallet | null> {
    const url = `${this.walletsEndpoint}/`;

    const wallet = {
      name: model["name"],
      description: model["description"],
      currency: model["currency"],
    };

    try {
      const response = await fetch(url, {
        headers: this.headers,
        method: "POST",
        body: JSON.stringify(wallet),
      });
      const data = await response.json();
      return data;
    } catch (e) {
      console.warn("Error when creating wallet", e);
    }
    return null;
  }
}
