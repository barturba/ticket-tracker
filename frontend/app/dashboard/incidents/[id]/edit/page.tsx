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
  title: "Edit Incident",
};

export default async function Page({ params }: { params: { id: string } }) {
  const id = params.id;
  const [incident, users, companies, configurationItems] = await Promise.all([
    fetchIncidentById(id),
    fetchUsers(),
    fetchCompanies(),
    fetchConfigurationItems(),
  ]);
  console.log(
    `dashboard/incidents/[id]/edit/incident: ${JSON.stringify(
      incident,
      null,
      2
    )}`
  );
  if (!incident) {
    console.log(`calling notFound`);
    notFound();
  }
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Incidents", href: "/dashboard/incidents" },
          {
            label: "Edit Incident",
            href: `/dashboard/incidents/${id}/edit`,
            active: true,
          },
        ]}
      />
      <EditForm
        incident={incident}
        companies={companies}
        initialUsers={users}
        configurationItems={configurationItems}
      />
    </main>
  );
}
