import { auth } from "@/auth";
import { ALL_ITEMS_LIMIT, ITEMS_PER_PAGE } from "../constants/constants";
import { UsersResponse } from "./types";
import { JWT_SECRET } from "./constants";
import jwt from "jsonwebtoken";
import { GetUserParams } from "@/types/users/base";

export async function getUsers(params: GetUserParams): Promise<UsersResponse> {
  try {
    const url = new URL(`${process.env.BACKEND}/v1/users`);

    const searchParams = url.searchParams;
    if (params.query) {
      searchParams.set("query", params.query);
    }
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", (params.page ?? 1).toString());
    searchParams.set("page_size", ITEMS_PER_PAGE.toString());

    const session = await auth();

    if (!session) {
      throw new Error("Session is null");
    }
    if (!JWT_SECRET) {
      throw new Error("JWT secret not defined");
    }
    const payload = { userId: session.userId! };
    const newToken = jwt.sign(payload, JWT_SECRET, {
      algorithm: "HS256",
      audience: "api",
      expiresIn: "1h",
    });

    const data = await fetch(url.toString(), {
      method: "GET",
      headers: {
        Authorization: `Bearer ${newToken}`,
      },
    });
    if (data.ok) {
      const UsersResponse: UsersResponse = await data.json();
      if (UsersResponse) {
        return {
          users: UsersResponse.users,
          metadata: UsersResponse.metadata,
        };
      } else {
        throw new Error("Failed to fetch users data: !UsersResponse");
      }
    } else {
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
): Promise<UsersResponse> {
  try {
    const url = new URL(`${process.env.BACKEND}/v1/users_all`);

    const searchParams = url.searchParams;
    searchParams.set("query", query);
    searchParams.set("sort", "-updated_at");
    searchParams.set("page", currentPage.toString());
    searchParams.set("page_size", ALL_ITEMS_LIMIT.toString());

    const data = await fetch(url.toString(), {
      method: "GET",
    });
    if (data.ok) {
      const UsersResponse: UsersResponse = await data.json();
      if (UsersResponse) {
        return {
          users: UsersResponse.users,
          metadata: UsersResponse.metadata,
        };
      } else {
        throw new Error("Failed to fetch users data: !UsersResponse");
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
    const url = new URL(`${process.env.BACKEND}/v1/users_latest`);
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
  const url = new URL(`${process.env.BACKEND}/v1/users/${id}`);

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
