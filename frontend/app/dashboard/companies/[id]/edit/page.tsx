import { fetchCompanyById } from "@/app/lib/actions/companies";
import EditForm from "@/app/ui/companies/edit-form";
import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import { Metadata } from "next";
import { notFound } from "next/navigation";

export const metadata: Metadata = {
  title: "Edit Company",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [company] = await Promise.all([fetchCompanyById(id)]);
  if (!company) {
    notFound();
  }
  return (
    <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: "Companies", href: "/dashboard/companies" },
          {
            label: "Edit Company",
            href: `/dashboard/companies/${id}/edit`,
            active: true,
          },
        ]}
      />
      <EditForm company={company} />
    </main>
  );
}
