import type { Metadata } from "../metadata/metadata";

export type User = {
  id: string;
  created_at: string;
  updated_at: string;
  first_name: string;
  last_name: string;
  email: string;
};

export type UserField = {
  id: string;
  name: string;
};

export type UsersData = {
  users: User[];
  metadata: Metadata;
};

export type UserData = {
  user: User;
  metadata: Metadata;
};

export type UsersField = {
  id: string;
  created_at: string;
  updated_at: string;
  first_name: string;
  last_name: string;
  email: string;
};

export type UserForm = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
};
