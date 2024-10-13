import Breadcrumbs from "@/app/ui/utils/breadcrumbs";
import EditForm from "@/app/ui/incidents/edit-form";
import { Metadata } from "next";
import { notFound } from "next/navigation";
import { fetchCIs } from "@/app/lib/actions/cis";
import { fetchCompanies } from "@/app/lib/actions/companies";
import { fetchIncidentById } from "@/app/lib/actions/incidents";
import { fetchUsers } from "@/app/lib/actions/users";
import { Heading, Subheading } from "@/app/components/heading";
import { Badge } from "@/app/components/badge";
import { BanknotesIcon, ChevronLeftIcon } from "@heroicons/react/24/outline";
import Link from "next/link";

export const metadata: Metadata = {
  title: "Edit Incident",
};

export default async function Page(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [incident, usersData, companiesData, cisData] = await Promise.all([
    fetchIncidentById(id),
    fetchUsers("", 1),
    fetchCompanies("", 1),
    fetchCIs("", 1),
  ]);
  if (!incident) {
    notFound();
  }
  console.log(`incident`, incident);
  return (
    <main>
      <div className="max-lg:hidden">
        <Link
          href="/incidents"
          className="inline-flex items-center gap-2 text-sm/6 text-zinc-500 dark:text-zinc-400"
        >
          <ChevronLeftIcon className="size-4 fill-zinc-400 dark:fill-zinc-500" />
          Incidents
        </Link>
      </div>
      <div className="mt-4 lg:mt-8">
        <div className="flex items-center gap-4">
          <Heading>Incident #{incident.id}</Heading>
          <Badge color="lime" state={incident.state}>
            {incident.state}
          </Badge>
        </div>
        <div className="isolate mt-2.5 flex flex-wrap justify-between gap-x-6 gap-y-4">
          <div className="flex flex-wrap gap-x-10 gap-y-4 py-1.5">
            <span className="flex items-center gap-3 text-base/6 text-zinc-950 sm:text-sm/6 dark:text-white">
              <BanknotesIcon className="size-4 shrink-0 fill-zinc-400 dark:fill-zinc-500" />
            </span>
          </div>
        </div>
      </div>
      <div className="mt-12">
        <EditForm
          incident={incident}
          companies={companiesData.companies}
          initialUsers={usersData.users}
          cis={cisData.cis}
        />
      </div>
    </main>
  );
}
