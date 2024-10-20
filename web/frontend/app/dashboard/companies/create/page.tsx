import { Metadata } from "next";
import CreateCompanyForm from "@/app/ui/sections/companies/create-form";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export const metadata: Metadata = {
  title: "Create Company",
};

export default async function CreateCompany() {
  return (
    <>
      <HeadingEdit name="Companies" backLink="/dashboard/companies" />
      <HeadingSubEdit
        name={`Create Company`}
        badgeState={undefined}
        badgeText={undefined}
      />
      <CreateCompanyForm />
    </>
  );
}
