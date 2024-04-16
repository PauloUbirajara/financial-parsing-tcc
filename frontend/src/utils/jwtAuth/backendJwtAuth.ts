import type {
  LoginCredentials,
  LoginResponse,
  RegisterCredentials,
  RegisterResponse,
  User,
} from "../../domain/models/auth";
import type { IJWTAuth } from "../../protocols/jwtAuth";

class BackendJWTAuth implements IJWTAuth {
  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    const AUTH_LOGIN_URL = import.meta.env.VITE_AUTH_LOGIN_URL;

    if (AUTH_LOGIN_URL === undefined) {
      return Promise.reject("Login URL not set");
    }

    const response = await fetch(AUTH_LOGIN_URL, {
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
    const AUTH_REGISTER_URL = import.meta.env.VITE_AUTH_REGISTER_URL;

    if (AUTH_REGISTER_URL === undefined) {
      return Promise.reject("Register URL not set");
    }

    const response = await fetch(AUTH_REGISTER_URL, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.ok) {
      return {};
    }

    const errorData: RegisterResponse = await response.json();
    return errorData;
  }

  async validate(accessToken: string | undefined): Promise<void> {
    const AUTH_VALIDATION_URL = import.meta.env.VITE_AUTH_VALIDATION_URL;

    if (AUTH_VALIDATION_URL === undefined) {
      return Promise.reject("Validation URL not set");
    }

    if (accessToken === undefined) {
      return Promise.reject("Access token not provided");
    }

    const response = await fetch(AUTH_VALIDATION_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      return Promise.reject("User not authorized");
    }

    return;
  }
}

export default new BackendJWTAuth();
