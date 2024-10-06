"use server";

import { signIn } from "@/auth";
import { AuthError } from "next-auth";
import { Incident } from "./definitions";
import { off } from "process";

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
    // await new Promise((resolve) => setTimeout(resolve, 1000));

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

const ITEMS_PER_PAGE = 6;
export async function fetchFilteredIncidents(
  query: string,
  currentPage: number
) {
  const offset = (currentPage - 1) * ITEMS_PER_PAGE;
  try {
    const url = new URL(`http://localhost:8080/v1/filtered_incidents`);
    const searchParams = url.searchParams;
    searchParams.set("query", query); // Set a new search param
    searchParams.set("limit", ITEMS_PER_PAGE.toString()); // Set a new search param
    searchParams.set("offset", offset.toString()); // Set a new search param

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const incidents = await data.json();
      if (incidents) {
        return incidents;
      }
    }
  } catch (error) {
    console.error("Database Error:", error);
    throw new Error("Failed to fetch incidents.");
  }
}
export async function deleteIncident(id: string) {
  // Simulate slow load
  // await new Promise((resolve) => setTimeout(resolve, 1000));
  console.log(`deleted incident; ${id}`);
}
