"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import setAlert from "@/app/lib/setAlert";
import { CreateUser, UpdateUser } from "./constants";
import {
  UserCreateInput,
  UserFormState,
  UserResponse,
  UserUpdateInput,
} from "@/types/users/base";
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
    await fetchApi<UserResponse>(url.toString(), {
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

export async function updateUser(
  id: string,
  prevState: UserFormState,
  formData: FormData
): Promise<UserFormState> {
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

  try {
    const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);
    await fetchApi<UserResponse>(url.toString(), {
      method: "PUT",
      body: JSON.stringify(validatedFields.data),
    });

    await setAlert({ type: "success", value: "User updated successfully!" });
    revalidatePath("/dashboard/users");
    redirect("/dashboard/users");

    return { message: "Update Successful" };
  } catch (error) {
    console.error(`update user error:`, error);

    if (error instanceof ApiError) {
      return {
        message: error.message || "Failed to update user.",
      };
    }

    return {
      message: "An unexpected error occurred. Failed to update user.",
    };
  }
}

export async function deleteUser(id: string): Promise<{ message: string }> {
  try {
    const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);
    await fetchApi<UserResponse>(url.toString(), {
      method: "DELETE",
    });

    await setAlert({ type: "success", value: "User deleted successfully!" });
    revalidatePath("/dashboard/users");

    return { message: "Deleted User." };
  } catch (error) {
    console.log(`Delete user error: ${error}`);

    if (error instanceof ApiError) {
      throw error;
    }

    throw new ApiError(500, "Failed to delete user.");
  }
}
