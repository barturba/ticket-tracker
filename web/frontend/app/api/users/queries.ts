import { auth } from "@/auth";
import { ITEMS_PER_PAGE } from "../constants/constants";
import { JWT_SECRET } from "./constants";
import jwt from "jsonwebtoken";
import { GetUserParams, UserResponse, UsersResponse } from "@/types/users/base";
import { ApiError } from "@/app/lib/api";

// Utility function to generate JWT token
async function generateAuthToken() {
  const session = await auth();

  if (!session?.userId) {
    throw new ApiError("Unauthorized: No valid session", 401);
  }
  if (!JWT_SECRET) {
    throw new ApiError("Server configuration error", 500);
  }

  return jwt.sign({ userId: session.userId }, JWT_SECRET, {
    algorithm: "HS256",
    audience: "api",
    expiresIn: "1h",
  });
}

async function authenticatedFetch<T>(
  url: URL,
  options: RequestInit = {}
): Promise<T> {
  const token = await generateAuthToken();

  const response = await fetch(url.toString(), {
    ...options,
    headers: {
      ...options.headers,
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    const errorText = await response.text();

    console.error(`API Error: ${response.status} ${errorText}`);
    throw new ApiError(
      `API request failed: ${response.status} ${errorText}`,
      response.status
    );
  }

  const data = await response.json();
  if (!data) {
    throw new ApiError("Invalid response from server", 500);
  }
  return data;
}

// Query functions
export async function getUsers(
  query: string,
  currentPage: number
): Promise<UsersResponse> {
  const url = new URL(`${process.env.BACKEND}/v1/users`);

  url.searchParams.set("query", query);
  url.searchParams.set("sort", "-updated_at");
  url.searchParams.set("page", currentPage.toString());
  url.searchParams.set("page_size", ITEMS_PER_PAGE.toString());

  return authenticatedFetch<UsersResponse>(url);
}

export async function getUsersAll(
  params: GetUserParams
): Promise<UsersResponse> {
  const url = new URL(`${process.env.BACKEND}/v1/users_all`);
  const { query = "", page = 1 } = params;

  url.searchParams.set("query", query);
  url.searchParams.set("sort", "-updated_at");
  url.searchParams.set("page", page.toString());
  url.searchParams.set("page_size", ITEMS_PER_PAGE.toString());

  return authenticatedFetch<UsersResponse>(url);
}

export async function fetchLatestUsers() {
  const url = new URL(`${process.env.BACKEND}/v1/users_latest`);

  url.searchParams.set("page_size", ITEMS_PER_PAGE.toString());
  url.searchParams.set("page", "1");

  return authenticatedFetch<UsersResponse>(url);
}

export async function getUser(id: string): Promise<UserResponse> {
  if (!id) {
    throw new ApiError("User ID is required", 400);
  }

  const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);
  url.searchParams.set("id", id);

  return authenticatedFetch<UserResponse>(url);
}

export type ApiErrorResponse = {
  status: number;
  message: string;
  code?: string;
};

export function isApiErrorResponse(obj: unknown): obj is ApiErrorResponse {
  return (
    typeof obj === "object" &&
    obj !== null &&
    "message" in obj &&
    "status" in obj
  );
}
