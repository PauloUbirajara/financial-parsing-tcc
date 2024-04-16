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

export interface RegisterCredentials {
  username: string;
  email: string;
  password: string;
  confirm_password: string;
}

export interface RegisterResponse {
  // Serializer errors
  error?: string;
  non_field_errors?: string[];
  username?: string[];
  email?: string[];
  password?: string[];
  confirmPassword?: string[];
}

export interface ForgotPasswordCredentials {
  username: string;
}

export interface User {
  username: string;
  email: string;
}
