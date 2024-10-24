"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import setAlert from "@/app/lib/setAlert";
import { UserCreateInput, UserResponse, UserUpdateInput } from "./types";
import { CreateUser, FormSchemaUser } from "./constants";
import { UserFormState } from "@/types/users/base";
import { ApiError } from "next/dist/server/api-utils";
import { fetchApi } from "@/app/lib/api";

export async function createUser(
  prevState: UserFormState,
  formData: FormData
): Promise<UserFormState> {
  const input: UserCreateInput = {
    first_name: formData.get("first_name") as string,
    last_name: formData.get("last_name") as string,
    email: formData.get("email") as string,
  };

  const validatedFields = CreateUser.safeParse(input);

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Create User.",
    };
  }

  try {
    const url = new URL(`${process.env.BACKEND}/v1/users`);
    const response = await fetchApi<UserResponse>(url.toString(), {
      method: "POST",
      body: JSON.stringify(validatedFields.data),
    });

    await setAlert({ type: "success", value: "User created successfully!" });
    revalidatePath("/dashboard/users");
    redirect("/dashboard/users");
    return { message: "User created successfully!" };
  } catch (error) {
    console.error(`create user error:`, error);
    if (error instanceof ApiError) {
      return {
        message: error.message || "Failed to create user.",
      };
    }
    return {
      message: "An unexpected error occurred. Failed to create user.",
    };
  }
}

const UpdateUser = FormSchemaUser.omit({ id: true });
export async function updateUser(
  id: string,
  prevState: UserFormState,
  formData: FormData
): Promise<UserFormState> {
  // Parse the form data using Zod
  const input: UserUpdateInput = {
    first_name: formData.get("first_name") as string,
    last_name: formData.get("last_name") as string,
    email: formData.get("email") as string,
  };

  const validatedFields = UpdateUser.safeParse(input);

  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Missing Fields. Failed to Update User.",
    };
  }

  // Validate form fields using Zod
  const { first_name, last_name, email } = validatedFields.data;

  // Prepare data for sending to the API.
  try {
    const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);
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
  await setAlert({ type: "success", value: "User updated successfully!" });
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath("/dashboard/users");
  redirect("/dashboard/users");
  return {
    message: "Update Successful",
  };
}

// DELETE

export async function deleteUser(id: string) {
  // Prepare data for sending to the API.
  try {
    const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);
    await fetch(url.toString(), {
      method: "DELETE",
    });
    await setAlert({ type: "success", value: "User deleted successfully!" });
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
