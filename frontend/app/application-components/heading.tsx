"use client";
import { Heading } from "@/app/components/heading";
import { Button } from "@/app/components/button";
import Alert from "./alert";
import { SignIn } from "./sign-in";

export default function AppHeading({
  name,
  createLabel,
  createLink,
}: {
  name: string;
  createLabel: string;
  createLink: string;
}) {
  return (
    <>
      <Alert />
      <SignIn />
      <div className="flex flex-wrap items-end justify-between gap-4">
        <div className="max-sm:w-full sm:flex-1">
          <Heading>{name}</Heading>
        </div>
        <Button href={createLink}>{createLabel}</Button>
      </div>
    </>
  );
}
