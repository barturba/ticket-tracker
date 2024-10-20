import type { Metadata } from "../metadata/metadata";

export type CI = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CIField = {
  id: string;
  name: string;
};

export type CIsData = {
  cis: CI[];
  metadata: Metadata;
};

export type CIData = {
  ci: CI;
  metadata: Metadata;
};

export type CIsField = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
};

export type CIForm = {
  id: string;
  name: string;
};
