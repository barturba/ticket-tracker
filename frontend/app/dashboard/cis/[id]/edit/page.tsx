import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import EditForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { fetchCompanies } from "@/app/lib/actions/companies";
import { fetchIncidentById } from "@/app/lib/actions/incidents";
import { fetchUsers } from "@/app/lib/actions/users";
import { fetchCIById, fetchCIs } from "@/app/lib/actions/cis";

export const metadata: Metadata = {
  title: "Edit Configuration Item",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [incident, usersData, companiesData, cisData] = await Promise.all([
    fetchCIById(id),
    fetchUsers("", 1),
    fetchCompanies("", 1),
    fetchCIs("", 1),
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
      <EditForm
        incident={incident}
        initialUsers={usersData.users}
        companies={companiesData.companies}
        cis={cisData.cis}
      />
    </main>
  );
}
