import type {
  ResetPasswordCredentials,
  ResetPasswordResponse,
} from "../../domain/models/resetPasswordDto";
import type {
  LoginCredentials,
  LoginResponse,
} from "../../domain/models/loginDto";
import type {
  RegisterCredentials,
  RegisterResponse,
} from "../../domain/models/registerDto";
import type { SendResetPasswordCredentials } from "../../domain/models/sendResetPasswordDto";
import type { IJWTAuth } from "../../protocols/jwtAuth";
import { constants } from "http2";

class BackendJWTAuth implements IJWTAuth {
  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    const url = process.env.VITE_AUTH_LOGIN_URL;

    if (url === undefined) {
      return Promise.reject("URL para login de usuários não definida.");
    }

    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data: LoginResponse = await response.json();
    return data;
  }

  async register(credentials: RegisterCredentials): Promise<RegisterResponse> {
    const url = process.env.VITE_AUTH_REGISTER_URL;

    if (url === undefined) {
      return Promise.reject("URL para cadastro de usuários não definida.");
    }

    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const errorData: RegisterResponse = await response.json();
      return errorData;
    }

    return {};
  }

  async sendResetPassword(
    credentials: SendResetPasswordCredentials,
  ): Promise<void> {
    const url = process.env.VITE_AUTH_SEND_RESET_PASSWORD_URL;

    if (url === undefined) {
      return Promise.reject(
        "URL para enviar redefinição de senha não definida.",
      );
    }

    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    const { detail } = await response.json();

    if (!response.ok) {
      console.warn(detail);
      return Promise.reject("Houve um erro ao solicitar redefinição de senha.");
    }

    return;
  }

  async resetPassword(
    credentials: ResetPasswordCredentials,
  ): Promise<ResetPasswordResponse> {
    const url = process.env.VITE_AUTH_RESET_PASSWORD_URL;

    if (url === undefined) {
      return Promise.reject("URL para redefinição de senha não definida.");
    }

    const response = await fetch(`${url}/${credentials.token}`, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      console.warn("Houve um erro ao redefinir senha.");
      const errorData = await response.json();
      return errorData;
    }

    return {};
  }

  async validate(accessToken: string): Promise<boolean> {
    const url = process.env.VITE_AUTH_VALIDATION_URL;

    if (url === undefined) {
      return Promise.reject("URL para validação de usuário não definida.");
    }

    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ token: accessToken }),
    });

    if (response.status === constants.HTTP_STATUS_UNAUTHORIZED) {
      return false;
    }

    return response.status === constants.HTTP_STATUS_OK;
  }

  async refresh(refreshToken: string): Promise<string | null> {
    const url = process.env.VITE_AUTH_REFRESH_TOKEN_URL;

    if (url === undefined) {
      return Promise.reject("URL para validação de usuário não definida.");
    }

    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refresh: refreshToken }),
    });

    const json = await response.json();
    const newAccess = json["access"] || null;

    return newAccess;
  }
}

export default new BackendJWTAuth();
