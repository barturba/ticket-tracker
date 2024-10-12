import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { State } from "@/app/lib/actions";
import { ITEMS_PER_PAGE } from "@/app/lib/constants";
import { IncidentData } from "@/app/lib/definitions/incidents";

// Incidents

// GET
export async function fetchIncidents(
  query: string,
  currentPage: number
): Promise<IncidentData> {
  try {
    const url = new URL(`http://localhost:8080/v1/incidents`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const incidentData: IncidentData = await data.json();
      if (incidentData) {
        return {
          incidents: incidentData.incidents,
          metadata: incidentData.metadata,
        };
      } else {
        // console.log(`fetchIncidents url: ${url.toString()}`);
        // console.log(`fetchIncidents error: !incidentData`);
        throw new Error("Failed to fetch incidents data: !incidentData");
      }
    } else {
      console.log(`fetchIncidents url: ${url.toString()}`);
      console.log(
        `fetchIncidents error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch incidents data: !data.ok");
    }
  } catch (error) {
    console.log(`fetchIncidents error: ${error}`);
    throw new Error(`Failed to fetch incidents data: ${error}`);
  }
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

const FormSchemaIncident = z.object({
  id: z.string(),
  shortDescription: z.string({
    required_error: "Please enter a short description.",
  }),
  description: z.string({
    required_error: "Please enter a description.",
  }),
  incidentId: z.string({
    invalid_type_error: "Please select a incident.",
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

// POST

const CreateIncident = FormSchemaIncident.omit({ id: true });
export async function createIncident(prevState: State, formData: FormData) {
  // Validate form fields using Zod
  const validatedFields = CreateIncident.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    incidentId: formData.get("incident_id"),
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
    incidentId,
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
        incident_id: incidentId,
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

// PUT

const UpdateIncident = FormSchemaIncident.omit({ id: true });
export async function updateIncident(
  id: string,
  prevState: State,
  formData: FormData
) {
  const validatedFields = UpdateIncident.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    incidentId: formData.get("incident_id"),
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
    incidentId,
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
        incident_id: incidentId,
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

// DELETE

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
