import type { IModelRepository } from "../../protocols/modelRepository";
import { BaseRepository } from "./baseRepository";

type CategoryRepositoryDTO = {
  accessToken: string | undefined;
};

export class CategoryRepository
  extends BaseRepository
  implements IModelRepository
{
  constructor(input: CategoryRepositoryDTO) {
    super({
      accessToken: input.accessToken,
      url: process.env.VITE_API_CATEGORIES_ENDPOINT,
    });
  }
}
