import { Metadata } from "@/app/lib/definitions";

export type Company = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CompanyField = {
  id: string;
  name: string;
};

export type CompanyData = {
  companies: Company[];
  metadata: Metadata;
};

export type CompaniesField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CompanyForm = {
  id: string;
  name: string;
};
