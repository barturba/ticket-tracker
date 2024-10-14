import { Metadata } from "next";
import CreateUserForm from "@/app/ui/sections/users/create-form";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";
import { Underdog } from "next/font/google";

export const metadata: Metadata = {
  title: "Create User",
};

export default async function CreateUser() {
  return (
    <>
      <HeadingEdit name="Users" backLink="/dashboard/users" />
      <HeadingSubEdit
        name={`Create User`}
        badgeState={undefined}
        badgeText={undefined}
      />
      <CreateUserForm />
    </>
  );
}
