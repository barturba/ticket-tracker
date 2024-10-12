import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import Form from "@/app/ui/companies/create-form";
import { Metadata } from "next";
import { CompanyData } from "@/app/lib/definitions/companies";
import { fetchCompanies } from "@/app/lib/actions/companies";

export const metadata: Metadata = {
  title: "Create Company",
};

export default async function Page() {
  const companydata: CompanyData = await fetchCompanies("", 1);
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
      <Form />
    </main>
  );
}
