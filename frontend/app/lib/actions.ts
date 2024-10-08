"use server";

import { signIn } from "@/auth";
import { AuthError } from "next-auth";
import { z } from "zod";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

const FormSchema = z.object({
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
  state: z.enum(
    ["New", "Assigned", "In Progress", "Pending", "On Hold", "Resolved"],
    {
      invalid_type_error: "Please select an incident state.",
    }
  ),
});

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

export async function fetchIncidents() {
  try {
    const data = await fetch(`http://localhost:8080/v1/incidents`, {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));

    if (data.ok) {
      const incidents = await data.json();
      if (incidents) {
        return incidents;
      }
    }
  } catch (error) {
    throw new Error(`Failed to fetch incidents. ${error}`);
  }
}

const ITEMS_PER_PAGE = 6;
export async function fetchFilteredIncidents(
  query: string,
  currentPage: number
) {
  const offset = (currentPage - 1) * ITEMS_PER_PAGE;
  try {
    const url = new URL(`http://localhost:8080/v1/filtered_incidents`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("limit", ITEMS_PER_PAGE.toString());
    searchParams.set("offset", offset.toString());

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
        console.log(`fetchFilteredIncidents data not received`);
        return [];
      }
    } else {
      console.log(`fetchFilteredIncidents data not ok`);
      return [];
    }
  } catch (error) {
    console.log(`fetchFilteredIncidents error: ${error}`);
    throw new Error("Failed to fetch incidents.");
  }
}
export async function fetchIncidentsPages(query: string) {
  try {
    const url = new URL(`http://localhost:8080/v1/filtered_incidents_count`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
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

export async function fetchCompanies() {
  try {
    const url = new URL(`http://localhost:8080/v1/companies`);
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const companies = await data.json();
      if (companies) {
        return companies;
      } else {
        return [];
      }
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
export async function fetchFilteredCompanies(
  query: string,
  currentPage: number
) {
  const offset = (currentPage - 1) * ITEMS_PER_PAGE;
  try {
    const url = new URL(`http://localhost:8080/v1/filtered_companies`);
    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("limit", ITEMS_PER_PAGE.toString());
    searchParams.set("offset", offset.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const companies = await data.json();
      if (companies) {
        return companies;
      } else {
        console.log(`fetchFilteredCompanies data not received`);
        return [];
      }
    } else {
      console.log(`fetchFilteredCompanies data not ok`);
      return [];
    }
  } catch (error) {
    console.log(`fetchFilteredCompanies error: ${error}`);
    throw new Error("Failed to fetch companies.");
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
export async function fetchUsers() {
  try {
    const url = new URL(`http://localhost:8080/v1/users`);
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
export async function fetchConfigurationItems() {
  try {
    const url = new URL(`http://localhost:8080/v1/configuration_items`);
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const configurationItems = await data.json();
      if (configurationItems) {
        return configurationItems;
      } else {
        return [];
      }
    }
  } catch (error) {
    console.log(`fetchConfigurationItems error: ${error}`);
    throw new Error("Failed to fetch configuration items data.");
  }
}
export async function fetchIncidentById(id: string) {
  console.log(`fetchIncidentByID getting following data: ${id}`);
  const url = new URL(`http://localhost:8080/v1/incident_by_id`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const incident = await data.json();
      console.log(
        `fetchIncidentByID got the following data: ${JSON.stringify(
          incident,
          null,
          2
        )}`
      );
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
const UpdateIncident = FormSchema.omit({ id: true });
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
    await new Promise((resolve) => setTimeout(resolve, 2000));
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

const CreateIncident = FormSchema.omit({ id: true });

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

export async function fetchLatestIncidents() {
  const offset = 0;
  try {
    const url = new URL(`http://localhost:8080/v1/incidents_latest`);
    const searchParams = url.searchParams;
    searchParams.set("limit", ITEMS_PER_PAGE.toString());
    searchParams.set("offset", offset.toString());

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

export async function fetchUsersByCompany(id: string) {
  const url = new URL(`http://localhost:8080/v1/users_by_company`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const users = await data.json();
      // console.log(`got the following data: ${JSON.stringify(users, null, 2)}`);
      if (users) {
        return users;
      } else {
        return "";
      }
    } else {
      console.log(`data not ok: ${JSON.stringify(data, null, 2)}`);
      return "";
    }
  } catch (error) {
    console.log(`fetchUsersByCompany error: ${error}`);
    throw new Error("Failed to fetch users data.");
  }
}
