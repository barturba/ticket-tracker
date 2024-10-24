import NextAuth, { type DefaultSession } from "next-auth";
import authConfig from "./auth.config";

import { JWT } from "next-auth/jwt";
import PostgresAdapter from "@auth/pg-adapter";
import type { Session } from "next-auth";

import { Pool } from "pg";

// Types for your session and token
interface CustomSession extends Session {
  accessToken: string;
  userRole?: string;
  userId?: string;
}

interface CustomToken extends JWT {
  accessToken?: string;
  userRole?: string;
  userId?: string;
}

declare module "next-auth" {
  /**
   * Returned by `auth`, `useSession`, `getSession` and received as a prop on the `SessionProvider` React Context
   */
  interface Session {
    user: {
      id: string;
      role: string;
      accessToken: string;
    } & DefaultSession["user"];
    accessToken: string;
    userRole?: string;
    userId?: string;
  }

  interface User {
    role?: string;
  }
}

declare module "next-auth/jwt" {
  interface User {
    role?: string;
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

export const { handlers, auth, signIn, signOut } = NextAuth({
  adapter: PostgresAdapter(pool),
  session: {
    strategy: "jwt",
    maxAge: 30 * 24 * 60 * 60, // 30 days
  },

  // providers are already defined in authConfig

  callbacks: {
    // Modify the JWT content
    async jwt({ token, user, account }): Promise<CustomToken> {
      // Initial sign in
      if (account && user) {
        // console.log(
        //   "jwt account.access_token: ",
        //   JSON.stringify(account.access_token, null, 2)
        // );
        // console.log("jwt user.role: ", JSON.stringify(user.role, null, 2));
        // console.log("jwt user.id: ", JSON.stringify(user.id, null, 2));
        return {
          ...token,
          userRole: user.role || "user", // Add custom claims
          userId: user.id,
        };
      }
      // console.log("jwt token: ", JSON.stringify(token, null, 2));
      return token;
    },
    async session({
      session,
      token,
    }: {
      session: CustomSession;
      token: CustomToken;
    }): Promise<CustomSession> {
      // console.log("session session: ", JSON.stringify(session, null, 2));
      // console.log("session token: ", JSON.stringify(token, null, 2));
      return {
        ...session,
        userRole: token.userRole,
        userId: token.userId,
      };
    },
    authorized: async ({ auth }) => {
      return !!auth;
    },
  },
  ...authConfig,
});
