import {
  fetchCompanies,
  fetchConfigurationItems,
  fetchIncidentById,
  fetchUsers,
} from "@/app/lib/actions";
import Breadcrumbs from "@/app/ui/incidents/breadcrumbs";
import Form from "@/app/ui/incidents/create-form";
import EditForm from "@/app/ui/incidents/edit-form";

export default async function Page({ params }: { params: { id: string } }) {
  const id = params.id;
  const [incident, companies, users, configurationItems] = await Promise.all([
    fetchIncidentById(id),
    fetchCompanies(),
    fetchUsers(),
    fetchConfigurationItems(),
  ]);
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Incidents", href: "/dashboard/incidents" },
          {
            label: "Edit Invoice",
            href: `/dashboard/incidents/${id}/edit`,
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
