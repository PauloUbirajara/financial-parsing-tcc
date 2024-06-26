import type {
  ResetPasswordCredentials,
  ResetPasswordResponse,
} from "../domain/models/resetPasswordDto";
import type {
  LoginCredentials,
  LoginResponse,
} from "../domain/models/loginDto";
import type {
  RegisterCredentials,
  RegisterResponse,
} from "../domain/models/registerDto";
import type { SendResetPasswordCredentials } from "../domain/models/sendResetPasswordDto";

export interface IJWTAuth {
  login(credentials: LoginCredentials): Promise<LoginResponse>;
  register(credentials: RegisterCredentials): Promise<RegisterResponse>;
  sendResetPassword(credentials: SendResetPasswordCredentials): Promise<void>;
  resetPassword(
    credentials: ResetPasswordCredentials,
  ): Promise<ResetPasswordResponse>;
  validate(accessToken: string): Promise<boolean>;
  refresh(refreshToken: string): Promise<string | null>;
}
