"use server";

import { AuthError } from "next-auth";
import { z } from "zod";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { CompanyData, Metadata } from "@/app/lib/definitions";

export async function authenticate(
  prevState: string | undefined,
  formData: FormData
) {
  try {
    await signIn("credentials", formData);
  } catch (error) {
    if (error instanceof AuthError) {
      switch (error.type) {
        case "CredentialsSignin":
          return "Invalid credentials.";
        default:
          return "Something went wrong.";
      }
    }
    throw error;
  }
}

// Incidents

const ITEMS_PER_PAGE = 6;
export async function fetchIncidents(query: string, currentPage: number) {
  try {
    const url = new URL(`http://localhost:8080/v1/incidents`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", currentPage.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const incidents = await data.json();
      if (incidents) {
        // await new Promise((resolve) => setTimeout(resolve, 60000));
        return incidents;
      } else {
        return [];
      }
    } else {
      return [];
    }
  } catch (error) {
    console.log(`fetchIncidents error: ${error}`);
    throw new Error("Failed to fetch incidents.");
  }
}

export async function fetchIncidentsPages(query: string) {
  try {
    const url = new URL(`http://localhost:8080/v1/incidents_count`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const count = await data.json();
      if (count > 0) {
        const totalPages = Math.ceil(Number(count) / ITEMS_PER_PAGE);
        return totalPages;
      } else {
        return 0;
      }
    }
  } catch (error) {
    console.log(`fetchIncidentsPages error: ${error}`);
    throw new Error("Failed to fetch incidents pages.");
  }
}

export async function fetchUsers() {
  try {
    const url = new URL(`http://localhost:8080/v1/users`);

    const searchParams = url.searchParams;
    searchParams.set("sort", "last_name");
    searchParams.set("page", "1");
    searchParams.set("page_size", "100");
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const users = await data.json();
      if (users) {
        return users;
      } else {
        return [];
      }
    }
  } catch (error) {
    console.log(`fetchUsers error: ${error}`);
    throw new Error("Failed to fetch users data.");
  }
}

export async function fetchCIs() {
  try {
    const url = new URL(`http://localhost:8080/v1/cis`);
    const searchParams = url.searchParams;
    searchParams.set("sort", "name");
    searchParams.set("page", "1");
    searchParams.set("page_size", "100");
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const configurationItems = await data.json();
      if (configurationItems) {
        console.log(
          `fetchCIs configurationItems: ${JSON.stringify(
            configurationItems,
            null,
            2
          )}`
        );
        return configurationItems;
      } else {
        return [];
      }
    }
  } catch (error) {
    console.log(`fetchCIs error: ${error}`);
    throw new Error("Failed to fetch configuration items data.");
  }
}

export async function fetchIncidentById(id: string) {
  const url = new URL(`http://localhost:8080/v1/incidents/${id}`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    if (data.ok) {
      const incident = await data.json();
      if (incident) {
        return incident;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchIncidentById error: ${error}`);
    throw new Error("Failed to fetch incident data.");
  }
}

export async function deleteIncident(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/incidents/${id}`);
    await fetch(url.toString(), {
      method: "DELETE",
    });
    // Revalidate the cache for the incident page
    revalidatePath("/dashboard/incidents");
    return { message: "Deleted Incident." };
  } catch (error) {
    console.log(`deleteIncident error: ${error}`);
    // If a database error occurs, return a more specific error.
    return {
      message: "Database Error: Failed to Update Incident.",
    };
  }
}

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

// Companies

export async function fetchCompanies(
  query: string,
  currentPage: number
): Promise<CompanyData> {
  try {
    const url = new URL(`http://localhost:8080/v1/companies`);
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
        return { companies: [], metadata: {} as Metadata };
      }
    } else {
      return { companies: [], metadata: {} as Metadata };
    }
  } catch (error) {
    console.log(`fetchCompanies error: ${error}`);
    throw new Error("Failed to fetch incidents pages.");
  }
}
export async function fetchCompaniesPages(query: string) {
  try {
    const url = new URL(`http://localhost:8080/v1/filtered_companies_count`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const count = await data.json();
      if (count > 0) {
        const totalPages = Math.ceil(Number(count) / ITEMS_PER_PAGE);
        return totalPages;
      } else {
        return 0;
      }
    }
  } catch (error) {
    console.log(`fetchCompaniesPages error: ${error}`);
    throw new Error("Failed to fetch companies pages.");
  }
}
// export async function fetchFilteredCompanies(
//   query: string,
//   currentPage: number
// ) {
//   const offset = (currentPage - 1) * ITEMS_PER_PAGE;
//   try {
//     const url = new URL(`http://localhost:8080/v1/companies`);
//     const searchParams = url.searchParams;
//     searchParams.set("query", query);
//     searchParams.set("limit", ITEMS_PER_PAGE.toString());
//     searchParams.set("offset", offset.toString());

//     const data = await fetch(url.toString(), {
//       method: "GET",
//     });
//     // Simulate slow load
//     // await new Promise((resolve) => setTimeout(resolve, 1000));
//     if (data.ok) {
//       const companies = await data.json();
//       if (companies) {
//         console.log(
//           `fetchFilteredCompanies data received: ${JSON.stringify(
//             companies.length,
//             null,
//             2
//           )}`
//         );
//         return companies;
//       } else {
//         console.log(`fetchFilteredCompanies data not received`);
//         return [];
//       }
//     } else {
//       console.log(`fetchFilteredCompanies data not ok`);
//       return [];
//     }
//   } catch (error) {
//     console.log(`fetchFilteredCompanies error: ${error}`);
//     throw new Error("Failed to fetch companies.");
//   }
// }
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

const FormSchemaIncident = z.object({
  id: z.string(),
  shortDescription: z.string({
    required_error: "Please enter a short description.",
  }),
  description: z.string({
    required_error: "Please enter a description.",
  }),
  companyId: z.string({
    invalid_type_error: "Please select a company.",
  }),
  assignedToId: z.string({
    invalid_type_error: "Please select a user to assign to.",
  }),
  configurationItemId: z.string({
    invalid_type_error: "Please select a configuration item to assign to.",
  }),
  state: z.enum(["New", "Assigned", "In Progress", "On Hold", "Resolved"], {
    invalid_type_error: "Please select an incident state.",
  }),
});

const CreateIncident = FormSchemaIncident.omit({ id: true });
export async function createIncident(prevState: State, formData: FormData) {
  // Validate form fields using Zod
  const validatedFields = CreateIncident.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    companyId: formData.get("company_id"),
    assignedToId: formData.get("assigned_to_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  // If form validation fails, return errors early. Otherwise, continue.
  if (!validatedFields.success) {
    console.log(
      `createIncident error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create Incident.",
    };
  }

  // Prepare data for sending to the API.
  const {
    shortDescription,
    description,
    companyId,
    assignedToId,
    configurationItemId,
    state,
  } = validatedFields.data;
  try {
    const url = new URL(`http://localhost:8080/v1/incidents`);
    await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        short_description: shortDescription,
        description: description,
        company_id: companyId,
        assigned_to_id: assignedToId,
        configuration_item_id: configurationItemId,
        state: state,
      }),
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    // if (data.ok) {
    //   console.log("got ok message");
    //   const incident = await data.json();
    //   if (incident) {
    //     return incident;
    //   }
    // }
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`createIncident error: ${error}`);
    return {
      message: "Database Error: Failed to Create Incident.",
    };
  }
  // Revalidate the cache for the incidents page and redirect the user.
  revalidatePath("/dashboard/incidents");
  redirect("/dashboard/incidents");
}

const UpdateIncident = FormSchemaIncident.omit({ id: true });
export async function updateIncident(
  id: string,
  prevState: State,
  formData: FormData
) {
  const validatedFields = UpdateIncident.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    companyId: formData.get("company_id"),
    assignedToId: formData.get("assigned_to_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  if (!validatedFields.success) {
    console.log(
      `updateIncident error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update Incident.",
    };
  }

  // Validate form fields using Zod
  const {
    shortDescription,
    description,
    companyId,
    assignedToId,
    configurationItemId,
    state,
  } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/incidents/${id}`);
    await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        id: id,
        short_description: shortDescription,
        description: description,
        company_id: companyId,
        assigned_to_id: assignedToId,
        configuration_item_id: configurationItemId,
        state: state,
      }),
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 2000));
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`updateIncident error: ${error}`);
    return {
      message: "Database Error: Failed to Update Incident.",
    };
  }
  // Revalidate the cache for the incidents page and redirect the user.
  revalidatePath("/dashboard/incidents");
  redirect("/dashboard/incidents");
}
export async function fetchLatestIncidents() {
  try {
    console.log(`calling fetchLatestIncidents()`);
    const url = new URL(`http://localhost:8080/v1/incidents_latest`);
    const searchParams = url.searchParams;
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", "1");

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const incidents = await data.json();
      if (incidents) {
        return incidents;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchFilteredIncidents error: ${error}`);
    throw new Error("Failed to fetch incidents.");
  }
}
const FormSchemaCompany = z.object({
  id: z.string(),
  name: z.string({
    required_error: "Please enter a short description.",
  }),
});

// Companies

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
  console.log("actions.ts updateCompany");
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
  console.log("actions.ts updateCompany validated fields");

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
