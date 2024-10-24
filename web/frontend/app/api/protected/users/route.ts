// Example: Using auth in API Routes
// app/api/protected/route.ts
import { getToken } from "next-auth/jwt";
import { NextResponse } from "next/server";
import jwt from "jsonwebtoken";

const JWT_SECRET = process.env.AUTH_SECRET;
// const JWT_EXPIRES_IN = process.env.JWT_EXPIRES_IN || "1h";

export async function GET(request: Request) {
  const token = await getToken({
    req: request,
    secret: process.env.AUTH_SECRET,
    // raw: true,
  });

  if (!token) {
    return new NextResponse(JSON.stringify({ error: "Not authenticated" }), {
      status: 401,
    });
  }

  console.log(`/protected/users token: ${JSON.stringify(token, null, 2)}`);
  console.log(
    `/protected/users JWT_SECRET: ${JSON.stringify(JWT_SECRET, null, 2)}`
  );
  if (!JWT_SECRET) {
    return new NextResponse(
      JSON.stringify({ error: "JWT secret not defined" }),
      {
        status: 500,
      }
    );
  }

  const newToken = jwt.sign(token, JWT_SECRET, {
    algorithm: "HS256",
    audience: "api",
    expiresIn: "1h",
  });
  console.log(
    `/protected/users newToken: ${JSON.stringify(newToken, null, 2)}`
  );
  // const session = await auth();

  // const secretKey = process.env.AUTH_SECRET;
  // if (!secretKey) {
  //   throw new Error("AUTH_SECRET is not defined in environment variables");
  // }
  // const token = await jwt.sign(
  //   { sessionToken: session?.user?.sessionToken },
  //   secretKey,
  //   {
  //     expiresIn: "1h",
  //   }
  // );

  console.log(
    `/protected/users newToken: ${JSON.stringify(verifyJWT(newToken), null, 2)}`
  );

  const url = new URL(`${process.env.BACKEND}/v1/users`);

  // Make request to Go backend
  const response = await fetch(url.toString(), {
    headers: {
      Authorization: `Bearer ${newToken}`,
    },
  });

  return Response.json(await response.json());
}

export const verifyJWT = (token: string) => {
  try {
    if (!JWT_SECRET) {
      throw new Error("JWT secret not defined");
    }
    const decoded = jwt.verify(token, JWT_SECRET);
    return { decoded };
  } catch (error) {
    return { error: `Invalid token ${error}` };
  }
};
