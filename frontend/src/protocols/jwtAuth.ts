import type {
  LoginCredentials,
  LoginResponse,
  RegisterCredentials,
  RegisterResponse,
  ForgotPasswordCredentials,
} from "../domain/models/auth";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<LoginResponse>;
  register(credentials: RegisterCredentials): Promise<RegisterResponse>;
  resetPassword(credentials: ForgotPasswordCredentials): Promise<void>;
  validate(accessToken: string): Promise<void>;
}
