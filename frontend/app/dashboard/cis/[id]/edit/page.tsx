import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import EditIncidentForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCompanies } from "@/app/lib/actions/companies";
import { getIncident } from "@/app/lib/actions/incidents";
import { getUsers } from "@/app/lib/actions/users";
import { getCI, getCIs } from "@/app/lib/actions/cis";

export const metadata: Metadata = {
  title: "Edit Configuration Item",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [incident, usersData, companiesData, cisData] = await Promise.all([
    getCI(id),
    getUsers("", 1),
    getCompanies("", 1),
    getCIs("", 1),
  ]);
  if (!incident) {
    notFound();
  }
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Configuration Items", href: "/dashboard/companies" },
          {
            label: "Edit Configuration Item",
            href: `/dashboard/cis/${id}/edit`,
            active: true,
          },
        ]}
      />
      <EditIncidentForm
        incident={incident}
        users={usersData.users}
        companies={companiesData.companies}
        cis={cisData.cis}
      />
    </main>
  );
}
