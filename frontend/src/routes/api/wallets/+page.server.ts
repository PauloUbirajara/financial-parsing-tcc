import { CurrencyRepository } from "$lib/repositories/currencyRepository";
import { WalletRepository } from "$lib/repositories/walletRepository";
import { constants } from "http2";
import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";
import { showToast } from "$lib/toast";
import { ToastType } from "../../../domain/models/toastMessage";

export const load: PageServerLoad = async (event) => {
  const accessToken = event.cookies.get("accessToken");
  let search: string | null = event.url.searchParams.get("search");
  let page: number | null = Number(event.url.searchParams.get("page")) || 1;
  if (isNaN(page) || page <= 0) {
    page = null;
  }

  try {
    const walletResponse = await new WalletRepository({ accessToken }).getAll({
      page,
      search,
    });
    const currencyResponse = await new CurrencyRepository({
      accessToken,
    }).getAll({
      page: null,
      search: null,
    });

    return {
      walletResponse,
      currencyResponse,
    };
  } catch (e) {
    console.warn(e);
    showToast({ title: "", message: "", type: ToastType.WARNING });
    redirect(constants.HTTP_STATUS_TEMPORARY_REDIRECT, "/api/wallets");
  }
};

export const actions: Actions = {
  create: async (event) => {
    const formData = await event.request.formData();
    const data = Object.fromEntries(formData.entries());
    const wallet = {
      name: data["name"],
      description: data["description"],
      currency: data["currency"],
    };

    const accessToken: string =
      (event.cookies.get("accessToken") as string) || "";
    const walletRepository = new WalletRepository({ accessToken });
    const response = await walletRepository.create(wallet);
    return response;
  },
};
