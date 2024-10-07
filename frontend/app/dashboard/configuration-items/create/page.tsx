import {
  fetchCompanies,
  fetchConfigurationItems,
  fetchUsers,
} from "@/app/lib/actions";
import Breadcrumbs from "@/app/ui/incidents/breadcrumbs";
import Form from "@/app/ui/incidents/create-form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Create Configuration Item",
};

export default async function Page() {
  const companies = await fetchCompanies();
  const users = await fetchUsers();
  const configurationItems = await fetchConfigurationItems();
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          {
            label: "Configuration Items",
            href: "/dashboard/configuration-items",
          },
          {
            label: "Create Configuraction Item",
            href: "/dashboard/configuration-items/create",
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
