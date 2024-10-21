import { Metadata } from "next";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";
import CreateCIForm from "@/app/ui/sections/cis/create-form";

export const metadata: Metadata = {
  title: "Create CI",
};

export default async function CreateCI() {
  return (
    <>
      <HeadingEdit name="CIs" backLink="/dashboard/cis" />
      <HeadingSubEdit
        name={`Create CI`}
        badgeState={undefined}
        badgeText={undefined}
      />
      <CreateCIForm />
    </>
  );
}
