import {
  fetchCompanies,
  fetchConfigurationItems,
  fetchIncidentById,
  fetchUsers,
} from "@/app/lib/actions";
import Breadcrumbs from "@/app/ui/incidents/breadcrumbs";
import EditForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";

export const metadata: Metadata = {
  title: "Edit Company",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [incident, companies, users, configurationItems] = await Promise.all([
    fetchIncidentById(id),
    fetchCompanies(),
    fetchUsers(),
    fetchConfigurationItems(),
  ]);
  if (!incident) {
    notFound();
  }
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Companies", href: "/dashboard/companies" },
          {
            label: "Edit Company",
            href: `/dashboard/companies/${id}/edit`,
            active: true,
          },
        ]}
      />
      <EditForm
        incident={incident}
        companies={companies}
        users={users}
        configurationItems={configurationItems}
      />
    </main>
  );
}
