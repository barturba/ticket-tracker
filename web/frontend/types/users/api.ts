export type ApiResponse<T> = {
  data: T;
  success: boolean;
  error?: string;
};