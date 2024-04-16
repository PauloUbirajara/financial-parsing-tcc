import type {
  LoginCredentials,
  LoginResponse,
  RegisterCredentials,
  RegisterResponse,
} from "../domain/models/auth";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<LoginResponse>;
  register(credentials: RegisterCredentials): Promise<RegisterResponse>;
  validate(accessToken: string): Promise<void>;
}
