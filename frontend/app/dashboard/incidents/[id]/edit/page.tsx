import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import EditForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCIs, getCIsAll } from "@/app/lib/actions/cis";
import { getCompanies, getCompaniesAll } from "@/app/lib/actions/companies";
import { getIncident } from "@/app/lib/actions/incidents";
import { getUsers, getUsersAll } from "@/app/lib/actions/users";
import { Heading, Subheading } from "@/app/components/heading";
import { Badge } from "@/app/components/badge";
import { BanknotesIcon, ChevronLeftIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export async function generateMetadata({
  params,
}: {
  params: { id: string };
}): Promise<Metadata> {
  let incident = await getIncident(params.id);

  return {
    title: incident && `Incident #${incident.id}`,
  };
}

export default async function Page(props: { params: Promise<{ id: string }> }) {
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
      <div className="mt-12">
        <EditForm
          incident={incident}
          companies={companiesData.companies}
          initialUsers={usersData.users}
          cis={cisData.cis}
        />
      </div>
    </>
  );
}
