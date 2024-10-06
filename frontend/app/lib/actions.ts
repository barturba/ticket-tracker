"use server";

import { signIn } from "@/auth";
import { AuthError } from "next-auth";
import { Incident } from "./definitions";

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

export async function fetchIncidents() {
  try {
    const data = await fetch(`http://localhost:8080/v1/incidents`, {
      method: "GET",
    });
    // Simulate slow load
    await new Promise((resolve) => setTimeout(resolve, 3000));

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
