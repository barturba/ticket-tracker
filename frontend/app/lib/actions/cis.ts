import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { State } from "@/app/lib/actions";
import { ITEMS_PER_PAGE } from "@/app/lib/constants";
import { CIData } from "@/app/lib/definitions/cis";

// CIs

// GET
export async function fetchCIs(
  query: string,
  currentPage: number
): Promise<CIData> {
  try {
    const url = new URL(`http://localhost:8080/v1/cis`);

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
      const ciData: CIData = await data.json();
      if (ciData) {
        return {
          cis: ciData.cis,
          metadata: ciData.metadata,
        };
      } else {
        // console.log(`fetchCIs url: ${url.toString()}`);
        // console.log(`fetchCIs error: !ciData`);
        throw new Error("Failed to fetch cis data: !ciData");
      }
    } else {
      console.log(`fetchCIs url: ${url.toString()}`);
      console.log(
        `fetchCIs error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch cis data: !data.ok");
    }
  } catch (error) {
    console.log(`fetchCIs error: ${error}`);
    throw new Error(`Failed to fetch cis data: ${error}`);
  }
}

export async function fetchLatestCIs() {
  try {
    console.log(`calling fetchLatestCIs()`);
    const url = new URL(`http://localhost:8080/v1/cis_latest`);
    const searchParams = url.searchParams;
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", "1");

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const cis = await data.json();
      if (cis) {
        return cis;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchFilteredCIs error: ${error}`);
    throw new Error("Failed to fetch cis.");
  }
}

export async function fetchCIById(id: string) {
  const url = new URL(`http://localhost:8080/v1/cis/${id}`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    if (data.ok) {
      const ci = await data.json();
      if (ci) {
        return ci;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchCIById error: ${error}`);
    throw new Error("Failed to fetch ci data.");
  }
}

const FormSchemaCI = z.object({
  id: z.string(),
  shortDescription: z.string({
    required_error: "Please enter a short description.",
  }),
  description: z.string({
    required_error: "Please enter a description.",
  }),
  ciId: z.string({
    invalid_type_error: "Please select a ci.",
  }),
  assignedToId: z.string({
    invalid_type_error: "Please select a user to assign to.",
  }),
  configurationItemId: z.string({
    invalid_type_error: "Please select a configuration item to assign to.",
  }),
  state: z.enum(["New", "Assigned", "In Progress", "On Hold", "Resolved"], {
    invalid_type_error: "Please select an ci state.",
  }),
});

// POST

const CreateCI = FormSchemaCI.omit({ id: true });
export async function createCI(prevState: State, formData: FormData) {
  // Validate form fields using Zod
  const validatedFields = CreateCI.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    ciId: formData.get("ci_id"),
    assignedToId: formData.get("assigned_to_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  // If form validation fails, return errors early. Otherwise, continue.
  if (!validatedFields.success) {
    console.log(
      `createCI error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create CI.",
    };
  }

  // Prepare data for sending to the API.
  const {
    shortDescription,
    description,
    ciId,
    assignedToId,
    configurationItemId,
    state,
  } = validatedFields.data;
  try {
    const url = new URL(`http://localhost:8080/v1/cis`);
    await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        short_description: shortDescription,
        description: description,
        ci_id: ciId,
        assigned_to_id: assignedToId,
        configuration_item_id: configurationItemId,
        state: state,
      }),
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    // if (data.ok) {
    //   console.log("got ok message");
    //   const ci = await data.json();
    //   if (ci) {
    //     return ci;
    //   }
    // }
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`createCI error: ${error}`);
    return {
      message: "Database Error: Failed to Create CI.",
    };
  }
  // Revalidate the cache for the cis page and redirect the user.
  revalidatePath("/dashboard/cis");
  redirect("/dashboard/cis");
}

// PUT

const UpdateCI = FormSchemaCI.omit({ id: true });
export async function updateCI(
  id: string,
  prevState: State,
  formData: FormData
) {
  const validatedFields = UpdateCI.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    ciId: formData.get("ci_id"),
    assignedToId: formData.get("assigned_to_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  if (!validatedFields.success) {
    console.log(
      `updateCI error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update CI.",
    };
  }

  // Validate form fields using Zod
  const {
    shortDescription,
    description,
    ciId,
    assignedToId,
    configurationItemId,
    state,
  } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/cis/${id}`);
    await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        id: id,
        short_description: shortDescription,
        description: description,
        ci_id: ciId,
        assigned_to_id: assignedToId,
        configuration_item_id: configurationItemId,
        state: state,
      }),
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 2000));
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`updateCI error: ${error}`);
    return {
      message: "Database Error: Failed to Update CI.",
    };
  }
  // Revalidate the cache for the cis page and redirect the user.
  revalidatePath("/dashboard/cis");
  redirect("/dashboard/cis");
}

// DELETE

export async function deleteCI(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/cis/${id}`);
    await fetch(url.toString(), {
      method: "DELETE",
    });
    // Revalidate the cache for the ci page
    revalidatePath("/dashboard/cis");
    return { message: "Deleted CI." };
  } catch (error) {
    console.log(`deleteCI error: ${error}`);
    // If a database error occurs, return a more specific error.
    return {
      message: "Database Error: Failed to Update CI.",
    };
  }
}
