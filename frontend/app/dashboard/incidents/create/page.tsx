import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/incidents/create-form";
import { Metadata } from "next";
import { fetchCIs } from "@/app/lib/actions/cis";
import { fetchCompanies } from "@/app/lib/actions/companies";
import { fetchUsers } from "@/app/lib/actions/users";

export const metadata: Metadata = {
  title: "Create Incident",
};

export default async function Page() {
  const companiesData = await fetchCompanies("", 1);
  const usersData = await fetchUsers("", 1);
  const cisData = await fetchCIs("", 1);
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
        companies={companiesData.companies}
        users={usersData.users}
        cis={cisData.cis}
      />
    </main>
  );
}
