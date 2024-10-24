import type { Metadata } from "../metadata/metadata";

export type User = {
  id: string;
  created_at: string;
  updated_at: string;
  first_name: string;
  last_name: string;
  email: string;
};

export type UsersField = {
  id: string;
  created_at: string;
  updated_at: string;
  first_name: string;
  last_name: string;
  email: string;
};

// Derive other types
export type UserCreateInput = Omit<User, "id" | "created_at" | "updated_at">;
export type UserUpdateInput = Partial<User>;

// Response types
export type UsersResponse = {
  users: User[];
  metadata: Metadata;
};

export type UserResponse = {
  user: User;
  metadata: Metadata;
};

export type UserState = {
  message?: string;
  errors?: {
    first_name?: string[];
    last_name?: string[];
    email?: string[];
  };
};
