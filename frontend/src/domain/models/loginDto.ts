export interface LoginCredentials {
  username: string;
  password: string;
}

export interface LoginResponse {
  access?: string;
  refresh?: string;

  detail?: string;
  username?: string[];
  password?: string[];
}
