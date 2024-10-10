import {
  fetchCompanies,
  fetchConfigurationItems,
  fetchUsers,
} from "@/app/lib/actions";
import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/incidents/create-form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Create Incident",
};

export default async function Page() {
  const companies = await fetchCompanies();
  const users = await fetchUsers();
  const configurationItems = await fetchConfigurationItems();
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Incidents", href: "/dashboard/incidents" },
          {
            label: "Create Incident",
            href: "/dashboard/incidents/create",
            active: true,
          },
        ]}
      />
      <Form
        companies={companies}
        users={users}
        configurationItems={configurationItems}
      />
    </main>
  );
}
