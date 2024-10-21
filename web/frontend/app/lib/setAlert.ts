"use server";
import { setCookie } from "cookies-next";
import { cookies } from "next/headers";

export default async function setAlert({
  type,
  value,
}: {
  type: string;
  value: string;
}) {
  await setCookie(type, value, { cookies });
}
