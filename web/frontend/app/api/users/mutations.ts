"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import setAlert from "@/app/lib/setAlert";
import { UserState } from "./types";
import { CreateUser, FormSchemaUser } from "./constants";

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
    const url = new URL(`${process.env.BACKEND}/v1/users`);
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
  await setAlert({ type: "success", value: "User created successfully!" });
  // Revalidate the cache for the users page and redirect the user.
  revalidatePath("/dashboard/users");
  redirect("/dashboard/users");
}

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
