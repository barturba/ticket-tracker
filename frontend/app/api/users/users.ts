"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { ALL_ITEMS_LIMIT, ITEMS_PER_PAGE } from "@/app/api/constants/constants";
import { UsersData } from "@/app/api/users/users.d";

export type UserState = {
  message?: string;
  errors?: {
    first_name?: string[];
    last_name?: string[];
    email?: string[];
  };
};

const FormSchemaUser = z.object({
  id: z.string(),
  first_name: z
    .string({
      required_error: "Please enter the first name.",
    })
    .min(1, { message: "First name must be at least 1 character." }),
  last_name: z
    .string({
      required_error: "Please enter the last name.",
    })
    .min(1, { message: "last name must be at least 1 character." }),
  email: z
    .string({
      required_error: "Please enter the email.",
    })
    .min(1, { message: "email must be at least 1 character." }),
});

// GET
export async function getUsers(
  query: string,
  currentPage: number
): Promise<UsersData> {
  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users`);

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
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users_all`);

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
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users_latest`);
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
  const url = new URL(`http://${process.env.BACKEND}:8080/v1/users/${id}`);

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
  const validatedFields = CreateUser.safeParse({
    first_name: formData.get("first_name"),
    last_name: formData.get("last_name"),
    email: formData.get("email"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create User.",
    };
  }

  // Validate form fields using Zod
  const { first_name, last_name, email } = validatedFields.data;

  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users`);
    const data = await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        first_name: first_name,
        last_name: last_name,
        email: email,
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
    first_name: formData.get("first_name"),
    last_name: formData.get("last_name"),
    email: formData.get("email"),
  });

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update User.",
    };
  }
  console.log(`updateUser PUT`);
  console.log(`id: ${id} ${JSON.stringify(validatedFields)}`);
  // Validate form fields using Zod
  const { first_name, last_name, email } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users/${id}`);
    console.log(`updateUser PUT`);
    const data = await fetch(url.toString(), {
      method: "PUT",
      body: JSON.stringify({
        first_name: first_name,
        last_name: last_name,
        email: email,
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
    const url = new URL(`http://${process.env.BACKEND}:8080/v1/users/${id}`);
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
