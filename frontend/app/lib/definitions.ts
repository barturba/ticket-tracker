import {
  BriefcaseIcon,
  CpuChipIcon,
  DocumentDuplicateIcon,
  HomeIcon,
  UserGroupIcon,
} from "@heroicons/react/24/outline";

export const Links = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: HomeIcon,
  },
  {
    name: "Incidents",
    href: "/dashboard/incidents",
    icon: DocumentDuplicateIcon,
  },
  { name: "Companies", href: "/dashboard/companies", icon: BriefcaseIcon },
  { name: "Users", href: "/dashboard/users", icon: UserGroupIcon },
  {
    name: "Configuration Items",
    href: "/dashboard/configuration-items",
    icon: CpuChipIcon,
  },
];
export type User = {
  id: string;
  email: string;
  password: string;
};

export type Incident = {
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

export type Company = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CompaniesField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type UsersField = {
  id: string;
  created_at: string;
  updated_at: string;
  first_name: string;
  last_name: string;
};

export type ConfigurationItemsField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

// Forms

export type IncidentForm = {
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

export type CompanyForm = {
  id: string;
  name: string;
};
