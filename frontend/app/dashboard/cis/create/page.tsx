import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/incidents/create-form";
import { Metadata } from "next";
import { getCIs } from "@/app/lib/actions/cis";
import { getCompanies } from "@/app/lib/actions/companies";
import { getUsers } from "@/app/lib/actions/users";

export const metadata: Metadata = {
  title: "Create Configuration Item",
};

export default async function Page() {
  const companiesData = await getCompanies("", 1);
  const usersData = await getUsers("", 1);
  const cisData = await getCIs("", 1);
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          {
            label: "Configuration Items",
            href: "/dashboard/cis",
          },
          {
            label: "Create Configuraction Item",
            href: "/dashboard/cis/create",
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
