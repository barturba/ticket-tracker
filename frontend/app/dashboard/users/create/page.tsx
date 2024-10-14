import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/users/create-form";
import { Metadata } from "next";
import { getCompanies } from "@/app/lib/actions/companies";
import { getUsers } from "@/app/lib/actions/users";
import { getCIs } from "@/app/lib/actions/cis";

export const metadata: Metadata = {
  title: "Create User",
};

export default async function Page() {
  const companiesData = await getCompanies("", 1);
  const usersData = await getUsers("", 1);
  const cisData = await getCIs("", 1);
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Users", href: "/dashboard/users" },
          {
            label: "Create User",
            href: "/dashboard/users/create",
            active: true,
          },
        ]}
      />
      <Form
      // companies={companiesData.companies}
      // users={usersData.users}
      // cis={cisData.cis}
      />
    </main>
  );
}
