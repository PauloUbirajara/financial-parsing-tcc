import type { LoginCredentials, User } from "../models/auth";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<void>;
  logout(): Promise<void>;
  isAuthenticated(): Promise<User>;
}
