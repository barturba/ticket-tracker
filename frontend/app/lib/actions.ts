"use server";

import { z } from "zod";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { CompanyData, Metadata } from "@/app/lib/definitions";

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
