"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { ALL_ITEMS_LIMIT, ITEMS_PER_PAGE } from "@/app/api/constants/constants";
import type { CompaniesData } from "@/app/api/companies/companies.d";

export type CompanyState = {
  message?: string;
  errors?: {
    name?: string[];
  };
};

const FormSchemaCompany = z.object({
  id: z.string(),
  name: z
    .string({
      required_error: "Please enter a short description.",
    })
    .min(1, { message: "Short description must be at least 1 character." }),
});

// GET
export async function getCompanies(
  query: string,
  currentPage: number
): Promise<CompaniesData> {
  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/companies`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const CompaniesData: CompaniesData = await data.json();
      if (CompaniesData) {
        return {
          companies: CompaniesData.companies,
          metadata: CompaniesData.metadata,
        };
      } else {
        throw new Error("Failed to fetch companies data: !CompaniesData");
      }
    } else {
      console.log(`getCompanies url: ${url.toString()}`);
      console.log(
        `getCompanies error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch companies data: !data.ok");
    }
  } catch (error) {
    console.log(`getCompanies error: ${error}`);
    throw new Error(`Failed to fetch companies data: ${error}`);
  }
}

export async function getCompaniesAll(
  query: string,
  currentPage: number
): Promise<CompaniesData> {
  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/companies_all`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ALL_ITEMS_LIMIT.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const CompaniesData: CompaniesData = await data.json();
      if (CompaniesData) {
        return {
          companies: CompaniesData.companies,
          metadata: CompaniesData.metadata,
        };
      } else {
        throw new Error("Failed to fetch companies data: !CompaniesData");
      }
    } else {
      console.log(`getCompanies url: ${url.toString()}`);
      console.log(
        `getCompanies error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch companies data: !data.ok");
    }
  } catch (error) {
    console.log(`getCompanies error: ${error}`);
    throw new Error(`Failed to fetch companies data: ${error}`);
  }
}

export async function fetchLatestCompanies() {
  try {
    const url = new URL(
      `http://${process.env.BACKEND}:8080/v1/companies_latest`
    );
    const searchParams = url.searchParams;
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", "1");

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const companies = await data.json();
      if (companies) {
        return companies;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchFilteredCompanies error: ${error}`);
    throw new Error("Failed to fetch companies.");
  }
}

export async function getCompany(id: string) {
  const url = new URL(`http://${process.env.BACKEND}:8080/v1/companies/${id}`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    if (data.ok) {
      const company = await data.json();
      if (company) {
        return company;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`getCompany error: ${error}`);
    throw new Error("Failed to fetch company data.");
  }
}

// POST

const CreateCompany = FormSchemaCompany.omit({ id: true });
export async function createCompany(
  prevState: CompanyState,
  formData: FormData
): Promise<CompanyState> {
  const validatedFields = CreateCompany.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create Company.",
    };
  }

  // Validate form fields using Zod
  const { name } = validatedFields.data;

  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/companies`);
    const data = await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        name: name,
      }),
    });
    if (data.ok) {
      const company = await data.json();
      if (company) {
        console.log(`createCompany success`);
      } else {
        console.log(`createCompany error: !company`);
        return {
          message: "Database Error: Failed to Create Company.",
        };
      }
    } else {
      console.log(
        `createCompany error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Create Company.",
      };
    }
  } catch (error) {
    console.log(`createCompany error: ${error}`);
    return {
      message: "Database Error: Failed to Create Company.",
    };
  }
  // Revalidate the cache for the companies page and redirect the user.
  revalidatePath("/dashboard/companies");
  redirect("/dashboard/companies");
}

// PUT

const UpdateCompany = FormSchemaCompany.omit({ id: true });
export async function updateCompany(
  id: string,
  prevState: CompanyState,
  formData: FormData
): Promise<CompanyState> {
  // Parse the form data using Zod
  const validatedFields = UpdateCompany.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update Company.",
    };
  }

  // Validate form fields using Zod
  const { name } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(
      `http://${process.env.BACKEND}:8080/v1/companies/${id}`
    );
    console.log(`updateCompany PUT`);
    const data = await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        name: name,
      }),
    });
    if (data.ok) {
      const company = await data.json();
      if (company) {
        console.log(`update success`);
      } else {
        console.log(`update error: !company`);
        return {
          message: "Database Error: Failed to Update Company.",
        };
      }
    } else {
      console.log(
        `update error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Update Company.",
      };
    }
  } catch (error) {
    console.log(`createCompany error: ${error}`);
    return {
      message: "Database Error: Failed to Update Company.",
    };
  }
  // Revalidate the cache for the companies page and redirect the user.
  revalidatePath(`/dashboard/companies/${id}/edit`);
  return {
    message: "Update Successful",
  };
}

// DELETE

export async function deleteCompany(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(
      `http://${process.env.BACKEND}:8080/v1/companies/${id}`
    );
    await fetch(url.toString(), {
      method: "DELETE",
    });
    // Revalidate the cache for the company page
    revalidatePath("/dashboard/companies");
    return { message: "Deleted Company." };
  } catch (error) {
    console.log(`deleteCompany error: ${error}`);
    // If a database error occurs, return a more specific error.
    return {
      message: "Database Error: Failed to Update Company.",
    };
  }
}
