import EditIncidentForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCIsAll } from "@/app/lib/actions/cis";
import { getCompaniesAll } from "@/app/lib/actions/companies";
import { getIncident } from "@/app/lib/actions/incidents";
import { getUsersAll } from "@/app/lib/actions/users";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export async function generateMetadata(props: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const params = await props.params;
  let incident = await getIncident(params.id);

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
