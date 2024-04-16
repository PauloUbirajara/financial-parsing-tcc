import type {
  LoginCredentials,
  LoginResponse,
  RegisterCredentials,
  RegisterResponse,
} from "../../domain/models/auth";
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

  async validate(accessToken: string): Promise<boolean> {
    try {
      await this.jwtAuth.validate(accessToken);
      return true;
    } catch (e) {
      console.warn("Error when validating user in AuthManager", e);
    }
    return false;
  }
}

export default new AuthManager(backendJwtAuth);
