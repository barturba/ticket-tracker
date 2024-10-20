import EditIncidentForm from "@/app/ui/sections/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCIsAll } from "@/app/api/cis/cis";
import { getCompaniesAll } from "@/app/api/companies/companies";
import { getIncident } from "@/app/api/incidents/incidents";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";
import { getUsersAll } from "@/app/api/users/users";

export async function generateMetadata(props: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const params = await props.params;
  const incident = await getIncident(params.id);

  return {
    title: incident && `Incident #${incident.id}`,
  };
}
export default async function Incident(props: {
  params: Promise<{ id: string }>;
}) {
  const params = await props.params;
  const id = params.id;
  const [incident, usersData, companiesData, cisData] = await Promise.all([
    getIncident(id),
    getUsersAll("", 1),
    getCompaniesAll("", 1),
    getCIsAll("", 1),
  ]);

  if (!incident) {
    notFound();
  }

  return (
    <>
      <HeadingEdit name="Incidents" backLink="/dashboard/incidents" />
      <HeadingSubEdit
        name={`Incident #${incident.id}`}
        badgeState={incident.state}
        badgeText={incident.state}
      />
      <EditIncidentForm
        incident={incident}
        companies={companiesData.companies}
        users={usersData.users}
        cis={cisData.cis}
      />
    </>
  );
}
