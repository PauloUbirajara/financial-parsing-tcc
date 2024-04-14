import type {
  LoginCredentials,
  LoginResponse,
  User,
} from "../../../../domain/models/auth";
import type { IJWTAuth } from "../../../../domain/usecases/jwtAuth";
import { AUTH_VALIDATE_URL, AUTH_LOGIN_URL } from "./constants";

export class SessionJWTAuth implements IJWTAuth {
  async login(credentials: LoginCredentials): Promise<void> {
    if (AUTH_LOGIN_URL === undefined) {
      return Promise.reject("Login URL not set");
    }

    const response = await fetch(AUTH_LOGIN_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      throw new Error("Login failed");
    }

    const data: LoginResponse = await response.json();
    sessionStorage.setItem("accessToken", data.accessToken);
    sessionStorage.setItem("refreshToken", data.refreshToken);
    return;
  }

  async logout(): Promise<void> {
    sessionStorage.removeItem("accessToken");
    sessionStorage.removeItem("refreshToken");
  }

  async isAuthenticated(): Promise<User> {
    if (AUTH_VALIDATE_URL === undefined) {
      return Promise.reject("Validation URL not set");
    }

    const accessToken = sessionStorage.getItem("accessToken");

    if (accessToken === null) {
      return Promise.reject("No access token");
    }

    const credentials = { token: accessToken };
    const response = await fetch(AUTH_VALIDATE_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      return Promise.reject("Validation failed");
    }

    const user: User = await response.json();
    return user;
  }
}
