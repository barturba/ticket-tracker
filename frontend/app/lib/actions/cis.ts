"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import {
  ALL_ITEMS_LIMIT,
  ITEMS_PER_PAGE,
} from "@/app/lib/definitions/constants";
import { CIData, CIsData } from "@/app/lib/definitions/cis";

export type CIState = {
  message?: string;
  errors?: {
    name?: string[];
  };
};

const FormSchemaCI = z.object({
  id: z.string(),
  name: z
    .string({
      required_error: "Please enter a name.",
    })
    .min(1, { message: "Name must be at least 1 character." }),
});

// GET
export async function getCIs(
  query: string,
  currentPage: number
): Promise<CIsData> {
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
    if (data.ok) {
      const CIsData: CIsData = await data.json();
      if (CIsData) {
        return {
          cis: CIsData.cis,
          metadata: CIsData.metadata,
        };
      } else {
        throw new Error("Failed to fetch cis data: !CIsData");
      }
    } else {
      console.log(`getCIs url: ${url.toString()}`);
      console.log(
        `getCIs error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch cis data: !data.ok");
    }
  } catch (error) {
    console.log(`getCIs error: ${error}`);
    throw new Error(`Failed to fetch cis data: ${error}`);
  }
}

export async function getCIsAll(
  query: string,
  currentPage: number
): Promise<CIsData> {
  try {
    const url = new URL(`http://localhost:8080/v1/cis_all`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ALL_ITEMS_LIMIT.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const CIsData: CIsData = await data.json();
      if (CIsData) {
        return {
          cis: CIsData.cis,
          metadata: CIsData.metadata,
        };
      } else {
        throw new Error("Failed to fetch cis data: !CIsData");
      }
    } else {
      console.log(`getCIs url: ${url.toString()}`);
      console.log(
        `getCIs error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch cis data: !data.ok");
    }
  } catch (error) {
    console.log(`getCIs error: ${error}`);
    throw new Error(`Failed to fetch cis data: ${error}`);
  }
}

export async function fetchLatestCIs() {
  try {
    const url = new URL(`http://localhost:8080/v1/cis_latest`);
    const searchParams = url.searchParams;
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", "1");

    const data = await fetch(url.toString(), {
      method: "GET",
    });
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

export async function getCI(id: string) {
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
    console.log(`getCI error: ${error}`);
    throw new Error("Failed to fetch ci data.");
  }
}

// POST

const CreateCI = FormSchemaCI.omit({ id: true });
export async function createCI(
  prevState: CIState,
  formData: FormData
): Promise<CIState> {
  const validatedFields = UpdateCI.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create CI.",
    };
  }

  // Validate form fields using Zod
  const { name } = validatedFields.data;

  try {
    const url = new URL(`http://localhost:8080/v1/cis`);
    const data = await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        name: name,
      }),
    });
    if (data.ok) {
      const ci = await data.json();
      if (ci) {
        console.log(`createCI success`);
      } else {
        console.log(`createCI error: !ci`);
        return {
          message: "Database Error: Failed to Create CI.",
        };
      }
    } else {
      console.log(
        `createCI error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Create CI.",
      };
    }
  } catch (error) {
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
  prevState: CIState,
  formData: FormData
): Promise<CIState> {
  // Parse the form data using Zod
  const validatedFields = UpdateCI.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update CI.",
    };
  }

  // Validate form fields using Zod
  const { name } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/cis/${id}`);
    console.log(`updateCI PUT`);
    const data = await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        name: name,
      }),
    });
    if (data.ok) {
      const ci = await data.json();
      if (ci) {
        console.log(`update success`);
      } else {
        console.log(`update error: !ci`);
        return {
          message: "Database Error: Failed to Update CI.",
        };
      }
    } else {
      console.log(
        `update error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Update CI.",
      };
    }
  } catch (error) {
    console.log(`createCI error: ${error}`);
    return {
      message: "Database Error: Failed to Update CI.",
    };
  }
  // Revalidate the cache for the cis page and redirect the user.
  revalidatePath(`/dashboard/cis/${id}/edit`);
  return {
    message: "Update Successful",
  };
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
