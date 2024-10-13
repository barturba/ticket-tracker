import { Metadata } from "@/app/lib/definitions";

export type User = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
};

export type UserField = {
  id: string;
  first_name: string;
  last_name: string;
};

export type UserData = {
  users: User[];
  metadata: Metadata;
};

export type UserForm = {
  id: string;
  short_description: string;
  description: string;
  company_id: string;
  assigned_to_id: string;
  configuration_item_id: string;
  state:
    | "New"
    | "Assigned"
    | "In Progress"
    | "Pending"
    | "On Hold"
    | "Resolved";
};
