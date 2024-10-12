import { lusitana } from "@/app/ui/fonts";
import { CreateIncident } from "@/app/ui/incidents/buttons";
import Pagination from "@/app/ui/utils/pagination";
import Table from "@/app/ui/incidents/table";
import Search from "@/app/ui/search";
import { Metadata } from "next";
import { Suspense } from "react";
import { fetchIncidents } from "@/app/lib/actions/incidents";
import { IncidentData } from "@/app/lib/definitions/incidents";
import { IncidentsTableSkeleton } from "@/app/ui/skeletons/incidents";

export const metadata: Metadata = {
  title: "Incidents",
};

export default async function Page(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const incidentData: IncidentData = await fetchIncidents(query, currentPage);

  const totalPages = incidentData.metadata.last_page;
  const incidents = incidentData.incidents;

  return (
    <div className="w-full">
      <div className="flex w-full items-center justify-between">
        <h1 className={`${lusitana.className} text-2xl`}>Incidents</h1>
      </div>

      <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
        <Search placeholder="Search incidents..." />
        <CreateIncident />
      </div>
      <Suspense key={query + currentPage} fallback={<IncidentsTableSkeleton />}>
        <Table incidents={incidents} />
      </Suspense>
      <div className="mt-5 flex w-full justify-center">
        <Pagination totalPages={totalPages} />
      </div>
    </div>
  );
}
