"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import {
  ALL_ITEMS_LIMIT,
  ITEMS_PER_PAGE,
} from "@/app/lib/definitions/constants";
import { UserData, UsersData } from "@/app/lib/definitions/users";

export type UserState = {
  message?: string;
  errors?: {
    shortDescription?: string[];
    description?: string[];
    userId?: string[];
    companyId?: string[];
    assignedToId?: string[];
    configurationItemId?: string[];
    state?: string[];
  };
};

const FormSchemaUser = z.object({
  id: z.string(),
  shortDescription: z
    .string({
      required_error: "Please enter a short description.",
    })
    .min(1, { message: "Short description must be at least 1 character." }),
  description: z.string({
    required_error: "Please enter a description.",
  }),
  assignedToId: z.string({
    invalid_type_error: "Please select a user to assign to.",
  }),
  companyId: z.string({
    invalid_type_error: "Please select a company.",
  }),
  configurationItemId: z.string({
    invalid_type_error: "Please select a configuration item to assign to.",
  }),
  state: z.enum(["New", "Assigned", "In Progress", "On Hold", "Resolved"], {
    invalid_type_error: "Please select an user state.",
  }),
});

// GET
export async function getUsers(
  query: string,
  currentPage: number
): Promise<UsersData> {
  try {
    const url = new URL(`http://localhost:8080/v1/users`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const UsersData: UsersData = await data.json();
      if (UsersData) {
        return {
          users: UsersData.users,
          metadata: UsersData.metadata,
        };
      } else {
        throw new Error("Failed to fetch users data: !UsersData");
      }
    } else {
      console.log(`getUsers url: ${url.toString()}`);
      console.log(
        `getUsers error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch users data: !data.ok");
    }
  } catch (error) {
    console.log(`getUsers error: ${error}`);
    throw new Error(`Failed to fetch users data: ${error}`);
  }
}

export async function getUsersAll(
  query: string,
  currentPage: number
): Promise<UsersData> {
  try {
    const url = new URL(`http://localhost:8080/v1/users_all`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ALL_ITEMS_LIMIT.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const UsersData: UsersData = await data.json();
      if (UsersData) {
        return {
          users: UsersData.users,
          metadata: UsersData.metadata,
        };
      } else {
        throw new Error("Failed to fetch users data: !UsersData");
      }
    } else {
      console.log(`getUsers url: ${url.toString()}`);
      console.log(
        `getUsers error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      throw new Error("Failed to fetch users data: !data.ok");
    }
  } catch (error) {
    console.log(`getUsers error: ${error}`);
    throw new Error(`Failed to fetch users data: ${error}`);
  }
}

export async function fetchLatestUsers() {
  try {
    const url = new URL(`http://localhost:8080/v1/users_latest`);
    const searchParams = url.searchParams;
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());
    searchParams.set("page", "1");

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const users = await data.json();
      if (users) {
        return users;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`fetchFilteredUsers error: ${error}`);
    throw new Error("Failed to fetch users.");
  }
}

export async function getUser(id: string) {
  const url = new URL(`http://localhost:8080/v1/users/${id}`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    if (data.ok) {
      const user = await data.json();
      if (user) {
        return user;
      } else {
        return "";
      }
    } else {
      return "";
    }
  } catch (error) {
    console.log(`getUser error: ${error}`);
    throw new Error("Failed to fetch user data.");
  }
}

// POST

const CreateUser = FormSchemaUser.omit({ id: true });
export async function createUser(
  prevState: UserState,
  formData: FormData
): Promise<UserState> {
  const validatedFields = UpdateUser.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    assignedToId: formData.get("assigned_to_id"),
    companyId: formData.get("company_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create User.",
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

  try {
    const url = new URL(`http://localhost:8080/v1/users`);
    const data = await fetch(url.toString(), {
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
    if (data.ok) {
      const user = await data.json();
      if (user) {
        console.log(`createUser success`);
      } else {
        console.log(`createUser error: !user`);
        return {
          message: "Database Error: Failed to Create User.",
        };
      }
    } else {
      console.log(
        `createUser error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Create User.",
      };
    }
  } catch (error) {
    console.log(`createUser error: ${error}`);
    return {
      message: "Database Error: Failed to Create User.",
    };
  }
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath("/dashboard/users");
  redirect("/dashboard/users");
}

// PUT

const UpdateUser = FormSchemaUser.omit({ id: true });
export async function updateUser(
  id: string,
  prevState: UserState,
  formData: FormData
): Promise<UserState> {
  // Parse the form data using Zod
  const validatedFields = UpdateUser.safeParse({
    shortDescription: formData.get("short_description"),
    description: formData.get("description"),
    assignedToId: formData.get("assigned_to_id"),
    companyId: formData.get("company_id"),
    configurationItemId: formData.get("configuration_item_id"),
    state: formData.get("state"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update User.",
    };
  }

  // Validate form fields using Zod
  const {
    shortDescription,
    description,
    assignedToId,
    companyId,
    configurationItemId,
    state,
  } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/users/${id}`);
    console.log(`updateUser PUT`);
    const data = await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        short_description: shortDescription,
        description: description,
        company_id: companyId,
        assigned_to_id: assignedToId,
        configuration_item_id: configurationItemId,
        state: state,
      }),
    });
    if (data.ok) {
      const user = await data.json();
      if (user) {
        console.log(`update success`);
      } else {
        console.log(`update error: !user`);
        return {
          message: "Database Error: Failed to Update User.",
        };
      }
    } else {
      console.log(
        `update error: !data.ok ${data.status} ${JSON.stringify(
          data.statusText
        )}`
      );
      return {
        message: "Database Error: Failed to Update User.",
      };
    }
  } catch (error) {
    console.log(`createUser error: ${error}`);
    return {
      message: "Database Error: Failed to Update User.",
    };
  }
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath(`/dashboard/users/${id}/edit`);
  return {
    message: "Update Successful",
  };
}

// DELETE

export async function deleteUser(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://localhost:8080/v1/users/${id}`);
    await fetch(url.toString(), {
      method: "DELETE",
    });
    // Revalidate the cache for the user page
    revalidatePath("/dashboard/users");
    return { message: "Deleted User." };
  } catch (error) {
    console.log(`deleteUser error: ${error}`);
    // If a database error occurs, return a more specific error.
    return {
      message: "Database Error: Failed to Update User.",
    };
  }
}
