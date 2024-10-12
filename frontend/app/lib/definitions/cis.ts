import { Metadata } from "@/app/lib/definitions";

export type CI = {
  id: string;
  created_at: string;
  updated_at: string;
  short_description: string;
  description: string;
  configuration_item_id: string;
  configuration_item_id_name: string;
  company_id: string;
  state: string;
  assigned_to: string;
  assigned_to_name: string;
};

export type CIData = {
  cis: CI[];
  metadata: Metadata;
};

export type CIField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CIForm = {
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
