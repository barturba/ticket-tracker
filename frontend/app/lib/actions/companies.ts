// Companies

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { State } from "@/app/lib/actions";
import { CompanyData } from "@/app/lib/definitions";
import { ITEMS_PER_PAGE } from "@/app/lib/constants";

export async function fetchCompanies(
  query: string,
  currentPage: number
): Promise<CompanyData> {
  try {
    const url = new URL(`http://localhost:8080/v1/companies`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "name");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const companyData: CompanyData = await data.json();
      if (companyData) {
        return {
          companies: companyData.companies,
          metadata: companyData.metadata,
        };
      } else {
        console.log(`fetchCompanies url: ${url.toString()}`);
        console.log(`fetchCompanies error: !companyData`);
        throw new Error("Failed to fetch companies data.");
      }
    } else {
      console.log(`fetchCompanies url: ${url.toString()}`);
      console.log(
        `fetchCompanies error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch companies data.");
    }
  } catch (error) {
    console.log(`fetchCompanies error: ${error}`);
    throw new Error(`Failed to fetch companies data: ${error}`);
  }
}
export async function deleteCompany(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/companies/${id}`);
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
      message: "Database Error: Failed to Delete Company.",
    };
  }
}

const FormSchemaCompany = z.object({
  id: z.string(),
  name: z.string({
    required_error: "Please enter a short description.",
  }),
});

const CreateCompany = FormSchemaCompany.omit({ id: true });
export async function createCompany(prevState: State, formData: FormData) {
  // Validate form fields using Zod
  const validatedFields = CreateCompany.safeParse({
    name: formData.get("name"),
  });

  // If form validation fails, return errors early. Otherwise, continue.
  if (!validatedFields.success) {
    console.log(
      `createCompany error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create Company.",
    };
  }

  // Prepare data for sending to the API.
  const { name } = validatedFields.data;
  try {
    const url = new URL(`http://localhost:8080/v1/companies`);
    await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        name: name,
      }),
    });
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`createCompany error: ${error}`);
    return {
      message: "Database Error: Failed to Create Company.",
    };
  }
  // Revalidate the cache for the companies page and redirect the user.
  revalidatePath("/dashboard/companies");
  redirect("/dashboard/companies");
}
const UpdateCompany = FormSchemaCompany.omit({ id: true });
export async function updateCompany(
  id: string,
  prevState: State,
  formData: FormData
) {
  const validatedFields = UpdateCompany.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    console.log(
      "actions.ts updateCompany !validatedFields.success: ",
      validatedFields.error.flatten().fieldErrors
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update Incident.",
    };
  }

  // Validate form fields using Zod
  const { name } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/companies/${id}`);
    await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        id: id,
        name: name,
      }),
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 2000));
  } catch (error) {
    return {
      message: "Database Error: Failed to Update Incident.",
    };
  }
  // Revalidate the cache for the companies page and redirect the user.
  revalidatePath("/dashboard/companies");
  redirect("/dashboard/companies");
}

export async function fetchCompanyById(id: string) {
  const url = new URL(`http://localhost:8080/v1/company_by_id`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
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
    console.log(`fetchCompanyById error: ${error}`);
    throw new Error("Failed to fetch company data.");
  }
}
