import type { Metadata } from "../metadata/metadata";
export type Incident = {
  id: string;
  created_at: string;
  updated_at: string;
  short_description: string;
  description: string;
  configuration_item_id: string;
  configuration_item_id_name: string;
  company_id: string;
  state: "New" | "Assigned" | "In Progress" | "On Hold" | "Resolved";
  assigned_to: string;
  assigned_to_name: string;
};

export type IncidentField = {
  id: string;
  name: string;
};

export type IncidentsData = {
  incidents: Incident[];
  metadata: Metadata;
};

export type IncidentData = {
  incident: Incident;
  metadata: Metadata;
};

export type IncidentsField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type IncidentForm = {
  id: string;
  short_description: string;
  description: {
    String: string;
    Valid: boolean;
  };
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
