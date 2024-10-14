import EditCompanyForm from "@/app/ui/companies/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCIsAll } from "@/app/lib/actions/cis";
import { getCompaniesAll } from "@/app/lib/actions/companies";
import { getCompany } from "@/app/lib/actions/companies";
import { getUsersAll } from "@/app/lib/actions/users";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export async function generateMetadata(props: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const params = await props.params;
  let company = await getCompany(params.id);

  return {
    title: company && `Company ID: ${company.id}`,
  };
}
export default async function Company(props: {
  params: Promise<{ id: string }>;
}) {
  const params = await props.params;
  const id = params.id;
  const [company] = await Promise.all([getCompany(id)]);

  if (!company) {
    notFound();
  }

  return (
    <>
      <HeadingEdit name="Companies" backLink="/dashboard/companies" />
      <HeadingSubEdit
        name={`Company #${company.id}`}
        badgeState={company.state}
        badgeText={company.state}
      />
      <EditCompanyForm company={company} />
    </>
  );
}
