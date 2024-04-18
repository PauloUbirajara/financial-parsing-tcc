export interface ResetPasswordCredentials {
  password: string;
  token: string;
}

export interface ResetPasswordResponse {
  error?: string;
  password?: string[];
  detail?: string;
  non_field_errors?: string[];
}
