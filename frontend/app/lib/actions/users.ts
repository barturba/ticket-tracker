"use server";
// Users

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";
import { ALL_ITEMS_LIMIT, ITEMS_PER_PAGE } from "@/app/lib/constants";
import { UserData } from "@/app/lib/definitions/users";
import { IncidentState } from "@/app/lib/actions/incidents";

export async function getUsers(
  query: string,
  currentPage: number
): Promise<UserData> {
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
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const userData: UserData = await data.json();
      if (userData) {
        return {
          users: userData.users,
          metadata: userData.metadata,
        };
      } else {
        // console.log(`getUsers url: ${url.toString()}`);
        // console.log(`getUsers error: !userData`);
        throw new Error("Failed to fetch users data: !userData");
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
): Promise<UserData> {
  try {
    const url = new URL(`http://localhost:8080/v1/users_all`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "last_name");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ALL_ITEMS_LIMIT.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const userData: UserData = await data.json();
      if (userData) {
        return {
          users: userData.users,
          metadata: userData.metadata,
        };
      } else {
        // console.log(`getUsers url: ${url.toString()}`);
        // console.log(`getUsers error: !userData`);
        throw new Error("Failed to fetch users data: !userData");
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
      message: "Database Error: Failed to Delete User.",
    };
  }
}

const FormSchemaUser = z.object({
  id: z.string(),
  name: z.string({
    required_error: "Please enter a short description.",
  }),
});

const CreateUser = FormSchemaUser.omit({ id: true });
export async function createUser(prevState: IncidentState, formData: FormData) {
  // Validate form fields using Zod
  const validatedFields = CreateUser.safeParse({
    name: formData.get("name"),
  });

  // If form validation fails, return errors early. Otherwise, continue.
  if (!validatedFields.success) {
    console.log(
      `createUser error: ${JSON.stringify(
        validatedFields.error.flatten().fieldErrors,
        null,
        2
      )}`
    );
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create User.",
    };
  }

  // Prepare data for sending to the API.
  const { name } = validatedFields.data;
  try {
    const url = new URL(`http://localhost:8080/v1/users`);
    await fetch(url.toString(), {
      method: "POST",
      body: JSON.stringify({
        name: name,
      }),
    });
  } catch (error) {
    // If a database error occurs, return a more specific error.
    console.log(`createUser error: ${error}`);
    return {
      message: "Database Error: Failed to Create User.",
    };
  }
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath("/dashboard/users");
  redirect("/dashboard/users");
}
const UpdateUser = FormSchemaUser.omit({ id: true });
export async function updateUser(
  id: string,
  prevState: IncidentState,
  formData: FormData
) {
  const validatedFields = UpdateUser.safeParse({
    name: formData.get("name"),
  });

  if (!validatedFields.success) {
    console.log(
      "actions.ts updateUser !validatedFields.success: ",
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
    const url = new URL(`http://localhost:8080/v1/users/${id}`);
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
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath("/dashboard/users");
  redirect("/dashboard/users");
}

export async function getUser(id: string) {
  const url = new URL(`http://localhost:8080/v1/user_by_id`);

  const searchParams = url.searchParams;
  searchParams.set("id", id);
  try {
    const data = await fetch(url.toString(), {
      method: "GET",
    });

    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
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
