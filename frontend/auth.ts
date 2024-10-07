import NextAuth from "next-auth";
import { authConfig } from "./auth.config";
import z from "zod";
import type { User } from "@/app/lib/definitions";
import Credentials from "next-auth/providers/credentials";

type UserResponse = {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
  email: string;
  api_key: string;
  token: string;
};

async function getUser(
  email: string,
  password: string
): Promise<User | undefined> {
  try {
    const data = await fetch(`http://localhost:8080/v1/login-test`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    });

    if (data.ok) {
      const user = await data.json();
      if (user) {
        return user;
      } else {
        return undefined;
      }
    } else {
      return undefined;
    }
  } catch (error) {
    throw new Error("Failed to login");
  }
}

export const { auth, signIn, signOut } = NextAuth({
  ...authConfig,
  providers: [
    Credentials({
      async authorize(credentials) {
        const parsedCredentials = z
          .object({
            email: z.string().email(),
            password: z.string().min(6),
          })
          .safeParse(credentials);

        if (parsedCredentials.success) {
          const { email, password } = parsedCredentials.data;
          const user = await getUser(email, password);
          if (!user) {
            return null;
          }
          return user;
        }
        return null;
      },
    }),
  ],
});
