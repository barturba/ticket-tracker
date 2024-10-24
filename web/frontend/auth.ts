import NextAuth, { type DefaultSession } from "next-auth";

import GitHub from "next-auth/providers/github";
import Google from "next-auth/providers/google";
import PostgresAdapter from "@auth/pg-adapter";

import { Pool } from "pg";
import { getUser } from "./app/api/users/users";
declare module "next-auth" {
  /**
   * Returned by `auth`, `useSession`, `getSession` and received as a prop on the `SessionProvider` React Context
   */
  interface Session {
    user: {
      id: string;
      role: string;
      sessionToken: string;
    } & DefaultSession["user"];
  }
}

const pool = new Pool({
  host: process.env.DATABASE_HOST,
  user: process.env.DATABASE_USER,
  password: process.env.DATABASE_PASSWORD,
  database: process.env.DATABASE_NAME,
  max: 20,
  idleTimeoutMillis: 30000,
  connectionTimeoutMillis: 2000,
  ssl: {
    rejectUnauthorized: false,
  },
});

export const { handlers, signIn, signOut, auth } = NextAuth({
  adapter: PostgresAdapter(pool),
  providers: [GitHub, Google],
  callbacks: {
    async session({ session, user }) {
      const [userData] = await Promise.all([getUser(user.id)]);
      // console.log(`user: ${JSON.stringify(user, null, 2)}`);
      // console.log(`session: ${JSON.stringify(session, null, 2)}`);
      // console.log(`userData: ${JSON.stringify(userData, null, 2)}`);

      // session.user.id = user.id;
      return {
        ...session,
        user: {
          ...session.user,
          role: userData.role,
          sessionToken: session.sessionToken,
        },
      };
    },
  },
});
