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
  title: "Edit Configuration Item",
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
          { label: "Configuration Items", href: "/dashboard/companies" },
          {
            label: "Edit Configuration Item",
            href: `/dashboard/configuration-items/${id}/edit`,
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
