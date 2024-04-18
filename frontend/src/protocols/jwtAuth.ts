import type {
  ResetPasswordCredentials,
  ResetPasswordResponse,
  SendResetPasswordCredentials,
} from "../domain/models/auth";
import type {
  LoginCredentials,
  LoginResponse,
} from "../domain/models/loginDto";
import type {
  RegisterCredentials,
  RegisterResponse,
} from "../domain/models/registerDto";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<LoginResponse>;
  register(credentials: RegisterCredentials): Promise<RegisterResponse>;
  sendResetPassword(credentials: SendResetPasswordCredentials): Promise<void>;
  resetPassword(
    credentials: ResetPasswordCredentials,
  ): Promise<ResetPasswordResponse>;
  validate(accessToken: string): Promise<void>;
}
