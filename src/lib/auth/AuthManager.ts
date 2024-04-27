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
import backendJwtAuth from "../../utils/jwtAuth/backendJwtAuth";

class AuthManager {
  private jwtAuth: IJWTAuth;

  constructor(jwtAuth: IJWTAuth) {
    this.jwtAuth = jwtAuth;
  }

  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    return this.jwtAuth.login(credentials);
  }

  async register(credentials: RegisterCredentials): Promise<RegisterResponse> {
    return this.jwtAuth.register(credentials);
  }

  async sendResetPassword(
    credentials: SendResetPasswordCredentials,
  ): Promise<boolean> {
    try {
      await this.jwtAuth.sendResetPassword(credentials);
      return true;
    } catch (e) {
      console.warn("Error when requesting password reset", e);
    }
    return false;
  }

  async resetPassword(
    credentials: ResetPasswordCredentials,
  ): Promise<ResetPasswordResponse> {
    return this.jwtAuth.resetPassword(credentials);
  }

  async validate(accessToken: string): Promise<boolean> {
    try {
      return await this.jwtAuth.validate(accessToken);
    } catch (e) {
      console.warn("Error when validating user", e);
    }
    return false;
  }

  async refresh(refreshToken: string): Promise<string | null> {
    try {
      return await this.jwtAuth.refresh(refreshToken);
    } catch (e) {
      console.warn("Error when refreshing user", e);
    }
    return null;
  }
}

export default new AuthManager(backendJwtAuth);
