import { fetchCompanies, fetchCIs, fetchUsers } from "@/app/lib/actions";
import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/companies/create-form";
import { Metadata } from "next";
import { CompanyData } from "@/app/lib/definitions";

export const metadata: Metadata = {
  title: "Create Company",
};

export default async function Page() {
  const companydata: CompanyData = await fetchCompanies("", 1);
  const companies = companydata.companies;
  const users = await fetchUsers();
  const configurationItems = await fetchCIs();
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Companies", href: "/dashboard/companies" },
          {
            label: "Create Company",
            href: "/dashboard/companies/create",
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
