import EditCIForm from "@/app/ui/sections/cis/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getCIsAll } from "@/app/api/cis/cis";
import { getCompaniesAll } from "@/app/api/companies/companies";
import { getCI } from "@/app/api/cis/cis";
import { getUsersAll } from "@/app/api/users/users";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";

export async function generateMetadata(props: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const params = await props.params;
  let ci = await getCI(params.id);

  return {
    title: ci && `CI #${ci.id}`,
  };
}
export default async function CI(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [ci, usersData, companiesData, cisData] = await Promise.all([
    getCI(id),
  ]);

  if (!ci) {
    notFound();
  }

  return (
    <>
      <HeadingEdit name="CIs" backLink="/dashboard/cis" />
      <HeadingSubEdit
        name={`CI #${ci.id}`}
        badgeState={ci.state}
        badgeText={ci.state}
      />
      <EditCIForm ci={ci} />
    </>
  );
}
