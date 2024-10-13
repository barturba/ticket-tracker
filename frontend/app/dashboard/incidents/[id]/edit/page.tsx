import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import EditForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { fetchCIs } from "@/app/lib/actions/cis";
import { fetchCompanies } from "@/app/lib/actions/companies";
import { fetchIncidentById } from "@/app/lib/actions/incidents";
import { fetchUsers } from "@/app/lib/actions/users";
import { Heading, Subheading } from "@/app/components/heading";
import { Badge } from "@/app/components/badge";
import { BanknotesIcon, ChevronLeftIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export const metadata: Metadata = {
  title: "Edit Incident",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [incident, usersData, companiesData, cisData] = await Promise.all([
    fetchIncidentById(id),
    fetchUsers("", 1),
    fetchCompanies("", 1),
    fetchCIs("", 1),
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
