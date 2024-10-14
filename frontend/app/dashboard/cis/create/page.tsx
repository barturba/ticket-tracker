import { Metadata } from "next";
import { getCIsAll } from "@/app/api/cis/cis";
import { getCompaniesAll } from "@/app/lib/actions/companies";
import { getUsersAll } from "@/app/lib/actions/users";
import CreateCIForm from "@/app/ui/sections/cis/create-form";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export const metadata: Metadata = {
  title: "Create CI",
};

export default async function CreateCI() {
  const [usersData, companiesData, cisData] = await Promise.all([
    getUsersAll("", 1),
    getCompaniesAll("", 1),
    getCIsAll("", 1),
  ]);

  return (
    <>
      <HeadingEdit name="CIs" backLink="/dashboard/cis" />
      <HeadingSubEdit name={`Create CI`} badgeState={"New"} badgeText={"New"} />
      <CreateCIForm
        companies={companiesData.companies}
        users={usersData.users}
        cis={cisData.cis}
      />
    </>
  );
}
