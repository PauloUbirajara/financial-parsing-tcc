import type {
  LoginCredentials,
  LoginResponse,
  RegisterCredentials,
  RegisterResponse,
  ResetPasswordCredentials,
  ResetPasswordResponse,
  SendResetPasswordCredentials,
} from "../domain/models/auth";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<LoginResponse>;
  register(credentials: RegisterCredentials): Promise<RegisterResponse>;
  sendResetPassword(credentials: SendResetPasswordCredentials): Promise<void>;
  resetPassword(
    credentials: ResetPasswordCredentials,
  ): Promise<ResetPasswordResponse>;
  validate(accessToken: string): Promise<void>;
}
