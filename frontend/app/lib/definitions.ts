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
  company_id: string;
  state: string;
  assigned_to: string;
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
  name: string;
};

export type ConfigurationItemsField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};
