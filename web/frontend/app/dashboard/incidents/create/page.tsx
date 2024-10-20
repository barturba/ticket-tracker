import { Metadata } from "next";
import { getCIsAll } from "@/app/api/cis/cis";
import { getCompaniesAll } from "@/app/api/companies/companies";
import { getUsersAll } from "@/app/api/users/users";
import CreateIncidentForm from "@/app/ui/sections/incidents/create-form";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export const metadata: Metadata = {
  title: "Create Incident",
};

export default async function CreateIncident() {
  const [usersData, companiesData, cisData] = await Promise.all([
    getUsersAll("", 1),
    getCompaniesAll("", 1),
    getCIsAll("", 1),
  ]);

  return (
    <>
      <HeadingEdit name="Incidents" backLink="/dashboard/incidents" />
      <HeadingSubEdit
        name={`Create Incident`}
        badgeState={"New"}
        badgeText={"New"}
      />
      <CreateIncidentForm
        companies={companiesData.companies}
        users={usersData.users}
        cis={cisData.cis}
      />
    </>
  );
}
