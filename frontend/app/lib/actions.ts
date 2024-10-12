"use server";

export type State = {
  errors?: {
    shortDescription?: string[];
    description?: string[];
    companyId?: string[];
    assignedToId?: string[];
    configurationItemId?: string[];
    state?: string[];
  };
  message?: string | null;
};
