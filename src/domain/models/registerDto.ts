export interface RegisterCredentials {
  username: string;
  email: string;
  password: string;
}

export interface RegisterResponse {
  // Serializer errors
  error?: string;
  non_field_errors?: string[];
  username?: string[];
  email?: string[];
  password?: string[];
}
