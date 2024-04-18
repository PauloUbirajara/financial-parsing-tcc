import type { Cookies } from "@sveltejs/kit";
import type { Category } from "../domain/models/category";
import type { IModelService } from "../protocols/modelService";

export class CategoryService implements IModelService<Category> {
  private categoriesEndpoint: string;
  private headers: Record<string, any>;

  constructor(cookies: Cookies) {
    if (cookies.get("accessToken") === undefined) {
      throw new Error("Could not get access token");
    }
    this.headers = {
      Authorization: `Bearer ${cookies.get("accessToken")!}`,
    };

    this.categoriesEndpoint = import.meta.env.VITE_API_CATEGORIES_ENDPOINT;
    if (this.categoriesEndpoint === undefined) {
      throw new Error("Categories endpoint not set");
    }
  }

  async getAll(): Promise<Category[]> {
    const url = this.categoriesEndpoint;
    const response = await fetch(url, { headers: this.headers });
    const categories = await response.json();
    return categories;
  }

  getById(): Promise<Category> {
    throw new Error("Method not implemented.");
  }
  updateById(id: string, updated: Category): Promise<Category> {
    throw new Error("Method not implemented.");
  }
  deleteById(id: string): Promise<Category> {
    throw new Error("Method not implemented.");
  }
}
