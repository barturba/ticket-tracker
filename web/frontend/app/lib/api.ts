import { ApiResponse } from "@/types/users/api";

export class ApiError extends Error {
  constructor(message: string, public status?: number, public code?: string) {
    super(message);
    this.name = "ApiError";
  }
}

export async function fetchApi<T>(
  url: string,
  options: RequestInit
): Promise<ApiResponse<T>> {
  try {
    const response = await fetch(url, {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...options.headers,
      },
    });
    const data = await response.json();
    if (!response.ok) {
      throw new ApiError(
        data.message || "Failed to fetch",
        response.status,
        data.code
      );
    }
    return {
      data,
      success: true,
    };
  } catch (error) {
    if (error instanceof ApiError) {
      throw error;
    }
    throw new ApiError(
      error instanceof Error ? error.message : "An error occurred"
    );
  }
}
