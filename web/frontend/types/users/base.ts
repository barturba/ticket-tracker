import type { Metadata, Timestamp } from "../common";

export type UserState =
  | "New"
  | "Assigned"
  | "In Progress"
  | "On Hold"
  | "Resolved";

export type UserProperties = {
  first_name: string;
  last_name: string;
  email: string;
};

export type User = {
  id: string;
  first_name: string;
  last_name: string;
  email?: string;
  state?: UserState;
  created_at: Timestamp;
  updated_at: Timestamp;
};

export type UserCreateInput = UserProperties;
export type UserUpdateInput = Partial<UserProperties>;

export type PaginatedResponse<T> = {
  data: T[];
  metadata: Metadata;
};

export type SingleResponse<T> = {
  data: T;
  metadata: Metadata;
};

export type UsersResponse = {
  users: User[];
  metadata: Metadata;
};
export type UserResponse = {
  user: User;
  metadata: Metadata;
};

export type UserFormState = {
  message?: string;
  errors?: {
    [K in keyof UserProperties]?: string[];
  };
};

export type UserListItem = Pick<
  User,
  "id" | "first_name" | "last_name" | "updated_at"
>;
export type UserSummary = Pick<User, "id" | "first_name" | "last_name">;

export type GetUserParams = {
  query?: string;
  page?: number;
  limit?: number;
};

export interface UserAPI {
  getUsers: (params: GetUserParams) => Promise<UsersResponse>;
  getUser: (id: string) => Promise<UserResponse>;
  createUser: (data: UserCreateInput) => Promise<SingleResponse<User>>;
  updateUser: (
    id: string,
    data: UserUpdateInput
  ) => Promise<SingleResponse<User>>;
  deleteUser: (id: string) => Promise<void>;
}
