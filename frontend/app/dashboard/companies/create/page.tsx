import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/sections/companies/create-form";
import { Metadata } from "next";
import { getCIs, getCIsAll } from "@/app/api/cis/cis";
import { getCompanies, getCompaniesAll } from "@/app/api/companies/companies";
import { getUsers, getUsersAll } from "@/app/api/users/users";
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
